package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jackc/tern/migrate"
	"log"
)

//migrateDatabase создание новой таблицы из файлов миграции, в случае отсутствии таблицы.
//Для запуска миграций использовать флаг -migrations.
func migrateDatabase(pool *pgxpool.Pool, path string, ctx context.Context) error {

	conn, err := pool.Acquire(ctx)
	if err != nil {
		return err
	}

	migrator, err := migrate.NewMigrator(ctx, conn.Conn(), "schema_version")
	if err != nil {
		return err
	}

	err = migrator.LoadMigrations(path)
	if err != nil {
		return err
	}

	err = migrator.Migrate(ctx)
	if err != nil {
		return err
	}

	ver, err := migrator.GetCurrentVersion(ctx)
	if err != nil {
		return err
	}

	log.Printf("migrations version: %d\n", ver)

	return nil
}
