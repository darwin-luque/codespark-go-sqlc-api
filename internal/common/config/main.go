package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Interface interface {
	Database() DatabaseInterface
	Server() ServerInterface
}

type Config struct {
	database DatabaseInterface
	server   ServerInterface
}

func New() (Interface, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	database, err := NewDatabaseConfig()

	if err != nil {
		return nil, err
	}

	server := NewServerConfig()

	return &Config{
		database: database,
		server:   server,
	}, nil
}

func (c *Config) Database() DatabaseInterface {
	return c.database
}

func (c *Config) Server() ServerInterface {
	return c.server
}

func getEnvVar(key string, fallback string) (string, error) {
	val := os.Getenv(key)

	if val == "" {
		if fallback == "" {
			return "", errors.New("Environment variable " + key + " not set")
		}
		return fallback, nil
	}

	return val, nil
}
