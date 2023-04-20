package config

import "go.uber.org/zap"

var ZapLogger *zap.Logger

func InitializeLogger() {
	// TODO: to change the level to NewProduction when deploying to production
	ZapLogger, _ = zap.NewDevelopment()
}
