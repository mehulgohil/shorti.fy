package config

import (
	"github.com/caarlos0/env/v8"
)

type envConfig struct {
	DynamoDBURL string `env:"DYNAMO_DB_URL"`
	AppPort     string `env:"APP_PORT"`
}

var EnvVariables envConfig

func LoadEnvVariables() {
	if err := env.Parse(&EnvVariables); err != nil {
		panic("unable to load environment variables")
	}
}
