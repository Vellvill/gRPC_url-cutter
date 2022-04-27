package main

import (
	"flag"
	"gRPC_cutter/internal/server"
)

var (
	cache      *bool
	migrations *bool
)

// init инциализация флагов
func init() {
	cache = flag.Bool("cache", false, "use cache instead of database")
	migrations = flag.Bool("migrations", false, "creating scheme")
}

func main() {
	flag.Parse()

	server.ApplicationStart(cache, migrations)
}
