package postgres

import (
	"context"
	"fmt"
	"gRPC_cutter/internal/config"
	"gRPC_cutter/internal/utils"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"time"
)

//NewClient возвращает пул соединений к базе данных, используемый для работы с базой.
func NewClient(ctx context.Context, config *config.Config, migrations *bool) (pool *pgxpool.Pool, err error) {
	dsn := config.Db.Dsn
	err = utils.DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgxpool.Connect(ctx, dsn)
		if err != nil {
			return err
		}
		return nil
	}, 5, 5*time.Second)

	if err != nil {
		log.Fatal(err)
	}
	if *migrations {
		if err = migrateDatabase(pool, config.Db.MigrationsPath, ctx); err != nil {
			return pool, fmt.Errorf("Unable to migrate, error: %s\n", err)
		}
	}
	log.Println("Connected to postgres database")
	return pool, nil
}
