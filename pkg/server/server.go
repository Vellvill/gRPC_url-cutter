package main

import (
	"context"
	"flag"
	"gRPC_cutter/pkg/config"
	cut "gRPC_cutter/pkg/cutter"
	"gRPC_cutter/pkg/postgres"
	"gRPC_cutter/pkg/repository"
	"gRPC_cutter/pkg/usecases"
	"log"
)

// GRPCServer ...
type GRPCServer struct {
	config *config.Config
	repo   usecases.Repository
	_      cut.UnimplementedURLShortenerServer
}

func (s *GRPCServer) Create(context.Context, *cut.CreateURLRequest) (*cut.CreateURLResponse, error) {
	return nil, nil
}
func (s *GRPCServer) Get(context.Context, *cut.GetURLRequest) (*cut.GetURLResponse, error) {
	return nil, nil
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

func Application() {
	flag.Parse()

	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
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

	newServer(cfg, repo)

	return
}
