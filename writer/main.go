package main

import (
	_ "github.com/mehulgohil/shorti.fy/writer/docs"
)

//	@title			shorti.fy
//	@version		1.0
//	@description	This is a backend api application for shorti.fy.
//	@host			localhost:8080
//	@BasePath		/
func main() {
	//initialize logger
	InitializeLogger()

	//initialize DB
	DynamoDB().InitLocalDBConnection()
	DynamoDB().InitTables()

	//initialize api routes
	app := Router().InitRouter(DynamoDB().(*DBClientHandler).DBClient)

	//initialize swagger routes
	SwaggerRouter().InitSwaggerRouter(app)

	err := app.Listen(":8080")
	if err != nil {
		panic("unable to start server")
	}
}
