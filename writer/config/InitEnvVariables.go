package config

import (
	"github.com/caarlos0/env/v8"
)

type EnvConfig struct {
	DynamoDBURL          string `env:"DYNAMO_DB_URL"`
	AppPort              string `env:"APP_PORT"`
	AWSAccessKeyID       string `env:"AWS_ACCESS_KEY_ID"`
	AWSSecretAccessToken string `env:"AWS_SECRET_ACCESS_TOKEN"`
	AWSRegion            string `env:"AWS_REGION"`
	APPDomain            string `env:"APP_DOMAIN"`
}

var EnvVariables EnvConfig

func LoadEnvVariables() {
	if err := env.Parse(&EnvVariables); err != nil {
		panic("unable to load environment variables")
	}
}
