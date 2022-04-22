package repository_test

import (
	"context"
	"gRPC_cutter/pkg/model"
	"gRPC_cutter/pkg/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepositoryPostgres_AddModel(t *testing.T) {
	s, down := repository.TestStore(t, DSN, MigrationsPath)
	defer down("url")

	testModel := model.TestModel(t)

	err := testModel.BeforeCreate()
	if err != nil {
		t.Fatal(err)
	}

	short, err := s.AddModel(context.TODO(), testModel.Longurl)
	if err != nil {
		t.Fatal(err)
	}

	assert.NoError(t, err)
	assert.NotNil(t, testModel, short)
}

func TestRepositoryPostgres_GetModel(t *testing.T) {
	s, down := repository.TestStore(t, DSN, MigrationsPath)
	defer down("url")

	testModel := model.TestModel(t)

	err := testModel.BeforeCreate()
	if err != nil {
		t.Fatal(err)
	}

	short, err := s.AddModel(context.TODO(), testModel.Longurl)
	if err != nil {
		t.Fatal(err)
	}

	long, err := s.GetModel(context.TODO(), short.Shorturl)
	if err != nil {
		t.Fatal(err)
	}

	code, err := short.Check()
	if err != nil {
		t.Log(err)
	}

	assert.NoError(t, err)
	assert.NotNil(t, long, code)
	t.Logf("http status of %v is %d\n", long, code)
}
