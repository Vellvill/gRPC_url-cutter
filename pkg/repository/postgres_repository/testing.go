package repository

import (
	"context"
	"fmt"
	"gRPC_cutter/pkg/config"
	"gRPC_cutter/pkg/postgres"
	"gRPC_cutter/pkg/usecases"
	"strings"
	"testing"
)

func TestStore(t *testing.T, databaseURL, migrationsPath string) (usecases.Repository, func(...string)) {
	t.Helper()

	var flag bool = true
	var migrations *bool

	migrations = &flag

	cfg := &config.Config{
		Db: struct {
			Dsn            string `yaml:"dsn"`
			MigrationsPath string `yaml:"migrations_path"`
		}{
			databaseURL,
			migrationsPath,
		},
	}

	client, err := postgres.NewClient(context.Background(), cfg, migrations)

	repo, err := NewDatabaseRep(client)
	if err != nil {
		t.Fatal(err)
	}

	return repo, func(tables ...string) {
		if len(tables) > 0 {
			q := fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))
			if _, err := client.Exec(context.Background(), q); err != nil {
				t.Fatal(err)
			}
		}
	}
}

func TestHash(t *testing.T) usecases.Repository {
	t.Helper()

	hash, err := NewHash()
	if err != nil {
		t.Fatal(err)
	}

	return hash
}
