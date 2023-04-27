package main

import (
	"github.com/iris-contrib/middleware/cors"
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

	config.ZapLogger.Info("Initializing Shortify Writer Service...")

	//initialize env variables
	config.LoadEnvVariables()

	//initialize DB
	config.DynamoDB().InitDBConnection()
	config.DynamoDB().InitTables()

	//initialize api routes
	app := Router().InitRouter(config.DynamoDB().(*config.DBClientHandler).DBClient)

	//allowed cors origin
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS"},
	})
	app.UseRouter(crs)

	//initialize swagger routes
	config.SwaggerRouter().InitSwaggerRouter(app)

	err := app.Listen(":" + config.EnvVariables.AppPort)
	if err != nil {
		panic("unable to start server")
	}
}
