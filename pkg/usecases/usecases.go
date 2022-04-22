package usecases

import (
	"context"
	"gRPC_cutter/pkg/model"
	"sync"
)

type Repository interface {
	AddModel(ctx context.Context, wg *sync.WaitGroup, url string) (*model.Model, error)
	GetModel(ctx context.Context, wg *sync.WaitGroup, shortURL string) (string, error)
}
