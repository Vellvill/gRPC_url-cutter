package usecases

import (
	"context"
	"gRPC_cutter/internal/model"
)

//Repository набор сигнатур методов, реализующиеся in-memory и postgres хранилищами
type Repository interface {
	AddModel(ctx context.Context, url string) (*model.Model, error)
	GetModel(ctx context.Context, shortURL string) (string, error)
}
