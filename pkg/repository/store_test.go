package repository_test

import (
	"os"
	"testing"
)

var (
	DSN            string
	MigrationsPath string
)

func TestMain(m *testing.M) {
	DSN, MigrationsPath = os.Getenv("DATABASE_URL"), os.Getenv("MIGRATIONS")
	if DSN == "" || MigrationsPath == "" {
		DSN = "postgresql://postgres:12345@localhost:5432/url_test"
		MigrationsPath = "./test_migrates"
	}

	os.Exit(m.Run())
}
