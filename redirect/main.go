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

	//initialize DB
	config.DynamoDB().InitLocalDBConnection()
	config.DynamoDB().InitTables()

	//initialize redis
	config.Redis().InitRedisConnection()

	//initialize api routes
	app := Router().InitRouter(config.DynamoDB().(*config.DBClientHandler).DBClient, config.Redis().(*config.RedisHandler).RedisClient)

	//initialize swagger routes
	config.SwaggerRouter().InitSwaggerRouter(app)

	err := app.Listen(":80")
	if err != nil {
		panic("unable to start server")
	}
}
