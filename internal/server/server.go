package server

import (
	"context"
	"fmt"
	"gRPC_cutter/internal/config"
	cut "gRPC_cutter/internal/cutter"
	"gRPC_cutter/internal/postgres"
	"gRPC_cutter/internal/repository/in-memory-hash_repository"
	"gRPC_cutter/internal/repository/postgres_repository"
	"gRPC_cutter/internal/usecases"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"sync"
)

// GRPCServer ...
type GRPCServer struct {
	config *config.Config
	repo   usecases.Repository
	cut.UnimplementedURLShortenerServer
	router *mux.Router
}

//Create сгенерированный proto файлом метод. Принимает сообщение в виде LongURL и возвращает ShortURL из репозитория.
func (s *GRPCServer) Create(ctx context.Context, req *cut.CreateURLRequest) (*cut.CreateURLResponse, error) {
	m, err := s.repo.AddModel(ctx, req.URL)
	if err != nil {
		return &cut.CreateURLResponse{ShortURL: ""}, err
	}
	return &cut.CreateURLResponse{ShortURL: m.Shorturl}, nil
}

//Get сгенерированный proto файлом метод. Принимает сообщение в виде ShortURL и возвращает подходящий LongURL из репозитория.
func (s *GRPCServer) Get(ctx context.Context, req *cut.GetURLRequest) (*cut.GetURLResponse, error) {
	res, err := s.repo.GetModel(ctx, req.ShortURL)
	if err != nil {
		return &cut.GetURLResponse{URL: ""}, err
	}
	return &cut.GetURLResponse{URL: res}, nil
}

// newServer приватная функция инициализации нового сервера
func newServer(config *config.Config, repo usecases.Repository) *GRPCServer {
	return &GRPCServer{
		config: config,
		repo:   repo,
	}
}

var wg sync.WaitGroup

//ApplicationStart функцию, задающая используемый репозиторий и регистрирующая gRPC сервис
func ApplicationStart(cache, migrations *bool) {
	wg.Add(2)
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	listner, err := net.Listen("tcp", fmt.Sprintf("%s", cfg.Server.Port))
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()

	var repo usecases.Repository

	if *cache {

		repo, err = in_memory_hash_repository.NewHash()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Cache is ready to accept new links")
	} else {

		client, err := postgres.NewClient(context.Background(), cfg, migrations)
		if err != nil {
			log.Fatal(err)
		}
		repo, err = postgres_repository.NewDatabaseRep(client)
		if err != nil {
			log.Fatal(err)
		}
	}

	server := newServer(cfg, repo)

	cut.RegisterURLShortenerServer(srv, server)

	r := mux.NewRouter()
	r.HandleFunc("/{short}", func(rw http.ResponseWriter, r *http.Request) {
		url := mux.Vars(r)
		redirectURL, err := server.Get(context.Background(), &cut.GetURLRequest{ShortURL: url["short"]})
		if err != nil {
			rw.Write([]byte("Failed to get url"))
			return
		}
		http.Redirect(rw, r, redirectURL.URL, 301)
	})

	server.router = r

	go func() {
		log.Println("Starting server...")
		if err := srv.Serve(listner); err != nil {
			log.Fatal(err)
			return
		}
		wg.Done()
	}()

	go func() {
		log.Println("Starting httpServer...")
		if err := http.ListenAndServe(":5000", server.router); err != nil {
			log.Fatal(err)
			return
		}
		wg.Done()
	}()

	wg.Wait()
}
