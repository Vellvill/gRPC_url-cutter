package in_memory_hash_repository_test

import (
	"context"
	"gRPC_cutter/pkg/model"
	in_memory_hash_repository "gRPC_cutter/pkg/repository/in-memory-hash_repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepositoryInMemoryHash_AddModel(t *testing.T) {
	s := in_memory_hash_repository.TestHash(t)

	testModel, err := model.TestModel(t)
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

func TestRepositoryInMemoryHash_GetModel(t *testing.T) {
	s := in_memory_hash_repository.TestHash(t)

	testModel, err := model.TestModel(t)
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
