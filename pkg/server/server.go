package server

import (
	"context"
	"flag"
	"fmt"
	"gRPC_cutter/pkg/config"
	cut "gRPC_cutter/pkg/cutter"
	"gRPC_cutter/pkg/postgres"
	"gRPC_cutter/pkg/repository"
	"gRPC_cutter/pkg/usecases"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

// GRPCServer ...
type GRPCServer struct {
	config *config.Config
	repo   usecases.Repository
	cut.UnimplementedURLShortenerServer
}

func (s *GRPCServer) Create(ctx context.Context, req *cut.CreateURLRequest) (*cut.CreateURLResponse, error) {
	wg := new(sync.WaitGroup)
	m, err := s.repo.AddModel(ctx, wg, req.URL)
	if err != nil {
		return &cut.CreateURLResponse{ShortURL: ""}, err
	}
	return &cut.CreateURLResponse{ShortURL: m.Shorturl}, nil
}
func (s *GRPCServer) Get(ctx context.Context, req *cut.GetURLRequest) (*cut.GetURLResponse, error) {
	wg := new(sync.WaitGroup)
	res, err := s.repo.GetModel(ctx, wg, req.ShortURL)
	if err != nil {
		return &cut.GetURLResponse{URL: ""}, err
	}
	return &cut.GetURLResponse{URL: res}, nil
}

var cache *bool

func init() {
	cache = flag.Bool("cache", false, "use cache instead of database")
}

func newServer(config *config.Config, repo usecases.Repository) *GRPCServer {
	return &GRPCServer{
		config: config,
		repo:   repo,
	}
}

func ApplicationStart() {
	flag.Parse()

	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	listner, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Server.Port))
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()

	var repo usecases.Repository

	if *cache {

		repo, err = repository.NewHash()
		if err != nil {
			log.Fatal(err)
		}

	} else {

		client, err := postgres.NewClient(context.Background(), cfg)
		if err != nil {
			log.Fatal(err)
		}
		repo, err = repository.NewDatabaseRep(client)
		if err != nil {
			log.Fatal(err)
		}
	}

	server := newServer(cfg, repo)

	cut.RegisterURLShortenerServer(srv, server)

	if err = srv.Serve(listner); err != nil {
		log.Fatal("Error starting server :8080")
	}
}
