package service

import "gRPC_cutter/internal/usecases"

//Implementation структуры с набором используемых в usecases методов для использования в рамках сервиса.
type Implementation struct {
	repo usecases.Repository
}

func NewImplementation(repo usecases.Repository) *Implementation {
	return &Implementation{repo: repo}
}
