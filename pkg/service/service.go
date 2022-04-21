package service

import "gRPC_cutter/pkg/usecases"

type Implementation struct {
	repo usecases.Repository
}

func New(repo usecases.Repository) (Implementation, error) {
	return Implementation{repo: repo}, nil
}
