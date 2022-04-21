package repository

import (
	"context"
	"gRPC_cutter/pkg/model"
)

type hash struct {
	hashmap map[string]*model.Model
}

func (h *hash) AddModel(ctx context.Context, model *model.Model) error {
	return nil
}
func (h *hash) GetModel(ctx context.Context, shortURL string) (string, error) {
	return "", nil
}
