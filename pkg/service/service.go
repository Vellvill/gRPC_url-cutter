package service

import "gRPC_cutter/pkg/usecases"

type Implementation struct {
	repo usecases.Repository
}

func NewImplementation(repo usecases.Repository) *Implementation {
	return &Implementation{repo: repo}
}
