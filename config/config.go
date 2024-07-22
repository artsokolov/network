package config

import (
	"fmt"
	"os"
)

type Config struct {
	dbName string
	dbHost string
}

func Load() (*Config, error) {
	return &Config{
		os.Getenv("DB_NAME"),
		os.Getenv("DB_HOST"),
	}, nil
}

func (config *Config) DBName() string {
	return config.dbName
}

func (config *Config) ConnectionUri() string {
	return fmt.Sprintf("mongodb://%s:27017/?connect=direct", config.dbHost)
}
