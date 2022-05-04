package server

import (
	"context"
	"fmt"
	"gRPC_cutter/internal/config"
	cut "gRPC_cutter/internal/cutter"
	"gRPC_cutter/internal/postgres"
	hash "gRPC_cutter/internal/repository/in-memory-hash_repository"
	db "gRPC_cutter/internal/repository/postgres_repository"
	"gRPC_cutter/internal/usecases"
	"google.golang.org/grpc"
	"log"
	"net"
)

// GRPCServer ...
type GRPCServer struct {
	config *config.Config
	repo   usecases.Repository
	cut.UnimplementedURLShortenerServer
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

//ApplicationStart функцию, задающая используемый репозиторий и регистрирующая gRPC сервис
func ApplicationStart(cache, migrations *bool) error {
	cfg, err := config.GetConfig()
	if err != nil {
		return err
	}

	listner, err := net.Listen("tcp", fmt.Sprintf("%s", cfg.Server.Port))
	if err != nil {
		return err
	}

	srv := grpc.NewServer()

	var repo usecases.Repository

	if *cache {

		repo, err = hash.NewHash()
		if err != nil {
			return err
		}

		log.Printf("Cache repository starting\n")

	} else {

		client, err := postgres.NewClient(context.Background(), cfg, migrations)
		if err != nil {
			return err
		}

		repo, err = db.NewDatabaseRep(client)
		if err != nil {
			return err
		}

		log.Printf("Postgres repository starting\n")
	}

	server := newServer(cfg, repo)

	cut.RegisterURLShortenerServer(srv, server)

	log.Println("Starting server...")

	return srv.Serve(listner)
}
