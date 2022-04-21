package repository

import (
	"context"
	"gRPC_cutter/pkg/model"
)

type database struct {
}

func (d *database) AddModel(ctx context.Context, model *model.Model) error {
	return nil
}
func (d *database) GetModel(ctx context.Context, shortURL string) (string, error) {
	return "", nil
}
