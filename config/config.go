package config

import (
	"fmt"
	"os"
)

type Config struct {
	dbUser string
	dbPass string
	dbName string
	dbHost string
}

func Load() (*Config, error) {
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPass := os.Getenv("DB_PASS")
	dbUser := os.Getenv("DB_USER")

	return &Config{
		dbUser,
		dbPass,
		dbName,
		dbHost,
	}, nil
}

func (config *Config) DBName() string {
	return config.dbName
}

func (config *Config) ConnectionUri() string {
	return fmt.Sprintf("mongodb://%s:27017/?connect=direct", config.dbHost)
}
