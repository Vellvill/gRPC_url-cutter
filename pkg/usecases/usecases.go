package usecases

import (
	"context"
	"gRPC_cutter/pkg/model"
)

type Repository interface {
	AddModel(ctx context.Context, url string) (*model.Model, error)
	GetModel(ctx context.Context, shortURL string) (string, error)
}
