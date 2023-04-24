package main

import (
	"github.com/mehulgohil/shorti.fy/redirect/config"
	_ "github.com/mehulgohil/shorti.fy/redirect/docs"
)

// @title			shorti.fy - Redirect
// @version		1.0
// @description	This is a backend microservice for shorti.fy Redirect.
// @host			localhost:80
// @BasePath		/
func main() {
	//initialize logger
	config.InitializeLogger()
	//initialize env variables
	config.LoadEnvVariables()

	//initialize DB
	config.DynamoDB().InitAWSDBConnection()
	config.DynamoDB().InitTables()

	//initialize redis
	config.Redis().InitRedisConnection()

	//initialize api routes
	app := Router().InitRouter(config.DynamoDB().(*config.DBClientHandler).DBClient, config.Redis().(*config.RedisHandler).RedisClient)

	//initialize swagger routes
	config.SwaggerRouter().InitSwaggerRouter(app)

	err := app.Listen(":" + config.EnvVariables.AppPort)
	if err != nil {
		panic("unable to start server")
	}
}
