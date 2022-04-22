package server

import (
	"context"
	"fmt"
	"gRPC_cutter/pkg/config"
	cut "gRPC_cutter/pkg/cutter"
	"gRPC_cutter/pkg/postgres"
	"gRPC_cutter/pkg/repository/in-memory-hash_repository"
	"gRPC_cutter/pkg/repository/postgres_repository"
	"gRPC_cutter/pkg/usecases"
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

func (s *GRPCServer) Create(ctx context.Context, req *cut.CreateURLRequest) (*cut.CreateURLResponse, error) {
	m, err := s.repo.AddModel(ctx, req.URL)
	if err != nil {
		return &cut.CreateURLResponse{ShortURL: ""}, err
	}
	return &cut.CreateURLResponse{ShortURL: m.Shorturl}, nil
}
func (s *GRPCServer) Get(ctx context.Context, req *cut.GetURLRequest) (*cut.GetURLResponse, error) {
	res, err := s.repo.GetModel(ctx, req.ShortURL)
	if err != nil {
		return &cut.GetURLResponse{URL: ""}, err
	}
	return &cut.GetURLResponse{URL: res}, nil
}

func newServer(config *config.Config, repo usecases.Repository) *GRPCServer {
	return &GRPCServer{
		config: config,
		repo:   repo,
	}
}

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

		repo, err = in_memory_hash_repository.NewHash()
		if err != nil {
			return err
		}
		log.Println("Cache is ready to accept new links")
	} else {

		client, err := postgres.NewClient(context.Background(), cfg, migrations)
		if err != nil {
			return err
		}
		repo, err = postgres_repository.NewDatabaseRep(client)
		if err != nil {
			return err
		}
	}

	server := newServer(cfg, repo)

	cut.RegisterURLShortenerServer(srv, server)

	return srv.Serve(listner)
}
