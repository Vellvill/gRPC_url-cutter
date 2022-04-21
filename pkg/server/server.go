package main

import (
	"context"
	grp "gRPC_cutter/pkg/cutter"
	"gRPC_cutter/pkg/usecases"
	"log"
	"net"
)

// GRPCServer ...
type GRPCServer struct {
	grp.UnimplementedURLShortenerServer
	repo usecases.Repository
}

func (s *GRPCServer) AddURL(ctx context.Context, in *grp.AddURLRequest) (*grp.AddURLResponse, error) {
	return nil, nil
}
func (s *GRPCServer) GetURL(ctx context.Context, in *grp.GetURLRequest) (*grp.GetURLResponse, error) {
	return nil, nil
}

func main() {

	listner, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Error listening on port 8080")
	}

}
