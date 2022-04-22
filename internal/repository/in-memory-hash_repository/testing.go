package in_memory_hash_repository

import (
	"gRPC_cutter/internal/usecases"
	"testing"
)

func TestHash(t *testing.T) usecases.Repository {
	t.Helper()

	hash, err := NewHash()
	if err != nil {
		t.Fatal(err)
	}

	return hash
}
