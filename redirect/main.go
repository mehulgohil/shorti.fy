package main

import (
	_ "github.com/mehulgohil/shorti.fy/redirect/docs"
)

// @title			shorti.fy - Redirect
// @version		1.0
// @description	This is a backend microservice for shorti.fy Redirect.
// @host			localhost:8081
// @BasePath		/
func main() {
	//initialize logger
	InitializeLogger()

	//initialize DB
	DynamoDB().InitLocalDBConnection()

	//initialize api routes
	app := Router().InitRouter(DynamoDB().(*DBClientHandler).DBClient)

	//initialize swagger routes
	SwaggerRouter().InitSwaggerRouter(app)

	err := app.Listen(":8081")
	if err != nil {
		panic("unable to start server")
	}
}
