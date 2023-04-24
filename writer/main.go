package main

import (
	"github.com/mehulgohil/shorti.fy/writer/config"
	_ "github.com/mehulgohil/shorti.fy/writer/docs"
)

// @title			shorti.fy
// @version		1.0
// @description	This is a backend api application for shorti.fy.
// @host			localhost:8080
// @BasePath		/
func main() {
	//initialize logger
	config.InitializeLogger()
	//initialize env variables
	config.LoadEnvVariables()

	//initialize DB
	config.DynamoDB().InitLocalDBConnection()
	config.DynamoDB().InitTables()

	//initialize api routes
	app := Router().InitRouter(config.DynamoDB().(*config.DBClientHandler).DBClient)

	//initialize swagger routes
	config.SwaggerRouter().InitSwaggerRouter(app)

	err := app.Listen(":" + config.EnvVariables.AppPort)
	if err != nil {
		panic("unable to start server")
	}
}
