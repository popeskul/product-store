package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DB     Mongo
	Server Server
}

type Mongo struct {
	URI      string
	Username string
	Password string
	Name     string
}

type Server struct {
	Host string
	Port int
}

func New() (*Config, error) {
	cfg := &Config{}

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	if err := envconfig.Process("db", &cfg.DB); err != nil {
		return nil, err
	}

	if err := envconfig.Process("server", &cfg.Server); err != nil {
		return nil, err
	}

	return cfg, nil
}
