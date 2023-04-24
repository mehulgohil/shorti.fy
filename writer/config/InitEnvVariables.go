package config

import (
	"github.com/caarlos0/env/v8"
)

type envConfig struct {
	DynamoDBURL          string `env:"DYNAMO_DB_URL"`
	AppPort              string `env:"APP_PORT"`
	AWSAccessKeyID       string `env:"AWS_ACCESS_KEY_ID"`
	AWSSecretAccessToken string `env:"AWS_SECRET_ACCESS_TOKEN"`
	AWSRegion            string `env:"AWS_REGION"`
}

var EnvVariables envConfig

func LoadEnvVariables() {
	if err := env.Parse(&EnvVariables); err != nil {
		panic("unable to load environment variables")
	}
}
