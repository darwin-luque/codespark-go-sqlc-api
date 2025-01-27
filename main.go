package main

import (
	"context"
	"log"

	"github.com/darwin-luque/codespark-go-sqlc-api/internal/common/config"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/common/setup"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/server"
)

func main() {
	cfg, err := config.New()

	if err != nil {
		log.Fatalf("Error loading config: %v", err)
		panic(err)
	}

	conn, err := setup.SetupDatabase(cfg.Database())

	if err != nil {
		log.Fatalf("Error setting up database: %v", err)
		panic(err)
	}
	defer conn.Close(context.Background())

	server, err := server.New(cfg, conn)

	if err != nil {
		log.Fatalf("Error setting up server: %v", err)
		panic(err)
	}

	log.Fatal(server.Run())
}
