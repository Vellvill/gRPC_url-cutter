package main

import (
	"context"
	"flag"
	"fmt"
	cut "gRPC_cutter/internal/cutter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var (
	add string
	get string
)

func init() {
	flag.StringVar(&add, "add", "", "use for adding")
	flag.StringVar(&get, "get", "", "use for getting")
}

func main() {
	flag.Parse()

	if add == "" && get == "" {
		log.Fatal("Chose method with flag '-get' or 'add'")
	}
	client, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	c := cut.NewURLShortenerClient(client)
	if add != "" {
		res, err := c.Create(context.Background(), &cut.CreateURLRequest{URL: add})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res.ShortURL)
	} else {
		res, err := c.Get(context.Background(), &cut.GetURLRequest{ShortURL: get})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res.URL)
	}
}
