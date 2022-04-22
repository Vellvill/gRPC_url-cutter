package main

import (
	"flag"
	"gRPC_cutter/pkg/server"
	"log"
)

var cache *bool

func init() {
	cache = flag.Bool("cache", false, "use cache instead of database")
}

func main() {
	flag.Parse()
	err := server.ApplicationStart(cache)
	if err != nil {
		log.Fatal(err)
	}
}
