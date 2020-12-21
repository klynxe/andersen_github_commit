package db

import (
	"fmt"

	"github.com/caarlos0/env"
)

const (
	ErrCreateDatabaseConfig = "Error create database config"
)

type Config struct {
	Url    string `env:"MONGO_URI,required"`
	DbName string `env:"MONGO_DB_NAME,required"`
}

func (conf *Config) GetURL() string {
	return conf.Url
}

func NewConfig() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("%v: %w", ErrCreateDatabaseConfig, err)
	}
	return &cfg, nil
}
