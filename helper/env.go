package helper

import (
	"github.com/caarlos0/env"
	_ "github.com/joho/godotenv/autoload"
)

type config struct {
	SERVERKEY   string `env:"SERVERKEY"`
	ENVIRONMENT string `env:"ENVIRONMENT"`
}

func LoadConfig() *config {
	cfg := new(config)
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
	return cfg
}
