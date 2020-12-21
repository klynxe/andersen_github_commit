package worker

import (
	"github.com/caarlos0/env"
)

type Config struct {
	Url    string `env:"GITHUB_URL,required"`
	Token  string `env:"GITHUB_TOKEN,required"`
	Period int    `env:"PERIOD_SECOND,required"`
}

func (conf *Config) GetURL() string {
	return conf.Url
}

func (conf *Config) GetToken() string {
	return conf.Token
}

func (conf *Config) GetPeriod() int {
	return conf.Period
}

func NewConfig() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
