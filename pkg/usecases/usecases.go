package usecases

import (
	"context"
	"sync"
)

type Repository interface {
	AddModel(ctx context.Context, wg *sync.WaitGroup, url string) error
	GetModel(ctx context.Context, wg *sync.WaitGroup, shortURL string) (string, error)
}
