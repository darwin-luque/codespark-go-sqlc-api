package main

import (
	"log"

	"github.com/darwin-luque/codespark-go-sqlc-api/internal/common/config"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/common/setup"
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
}
