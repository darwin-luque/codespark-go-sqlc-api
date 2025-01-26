package config

import (
	"errors"
	"os"
)

type Interface interface {
	Database() DatabaseInterface
}

type Config struct {
	database DatabaseInterface
}

func New() (Interface, error) {
	database, err := NewDatabaseConfig()

	if err != nil {
		return nil, err
	}

	return &Config{
		database: database,
	}, nil
}

func (c *Config) Database() DatabaseInterface {
	return c.database
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
