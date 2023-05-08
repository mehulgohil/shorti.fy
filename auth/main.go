package main

import (
	"github.com/mehulgohil/shorti.fy/auth/authenticator"
	"github.com/mehulgohil/shorti.fy/auth/config"
	"log"
)

func main() {
	//initialize logger
	config.InitializeLogger()
	//initialize env variables
	config.LoadEnvVariables()

	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	//initialize api routes
	app := Router().InitRouter(auth)

	err = app.Listen(":" + config.EnvVariables.AppPort)
	if err != nil {
		panic("unable to start server")
	}
}
