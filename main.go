package main

import (
	_ "github.com/mehulgohil/shorti.fy/docs"
)

// @title			shorti.fy
// @version		1.0
// @description	This is a backend api application for shorti.fy.
// @host			localhost:8080
// @BasePath		/
func main() {

	//initialize DB
	dbClient := DynamoDB().InitDBConnection()

	//initialize api routes
	app := Router().InitRouter(dbClient)

	//initialize swagger routes
	SwaggerRouter().InitSwaggerRouter(app)

	err := app.Listen(":8080")
	if err != nil {
		panic("unable to start server")
	}
}
