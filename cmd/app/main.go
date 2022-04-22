package main

import (
	"flag"
	"gRPC_cutter/pkg/server"
	"log"
)

var (
	cache      *bool
	migrations *bool
)

func init() {
	cache = flag.Bool("cache", false, "use cache instead of database")
	migrations = flag.Bool("migrations", false, "creating scheme")
}

func main() {
	flag.Parse()

	err := server.ApplicationStart(cache, migrations)
	if err != nil {
		log.Fatal(err)
	}
}
