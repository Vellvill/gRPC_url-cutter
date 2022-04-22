package postgres_repository

import (
	"context"
	"fmt"
	"gRPC_cutter/internal/config"
	"gRPC_cutter/internal/postgres"
	"gRPC_cutter/internal/usecases"
	"strings"
	"testing"
)

//TestStore инциализирует подключение к тестовой базе данных. Это не best practices, однако в рамках задачи
//тест даёт понять, что подключение успешно установлено. Функция возвращает функцию, делающая удаление таблиц
//и зависимостей после завершения тестов.
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
