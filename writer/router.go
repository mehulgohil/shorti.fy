package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/accesslog"
	"github.com/mehulgohil/shorti.fy/writer/interfaces"
	"io"
	"os"
	"sync"
)

var (
	irisRouter *router
	routerOnce sync.Once
)

type IRouter interface {
	InitRouter(dbClient interfaces.IDataAccessLayer) *iris.Application
}

type router struct{}

func (router *router) InitRouter(dbClient interfaces.IDataAccessLayer) *iris.Application {
	app := iris.New()
	ac := makeAccessLog()
	app.UseRouter(ac.Handler)

	healthCheckController := ServiceContainer().InjectHealthCheckController()
	shortifyWriterController := ServiceContainer().InjectShortifyWriterController(dbClient)

	app.Get("/healthcheck", healthCheckController.CheckServerHealthCheck)
	app.Post("/shorten", shortifyWriterController.WriterController)

	return app
}

func Router() IRouter {
	if irisRouter == nil {
		routerOnce.Do(func() {
			irisRouter = &router{}
		})
	}
	return irisRouter
}

// This helps to log the request and its metadata
func makeAccessLog() *accesslog.AccessLog {
	ac := accesslog.New(io.MultiWriter(os.Stdout))
	ac.Delim = ' '
	ac.ResponseBody = false
	ac.RequestBody = false
	ac.BytesReceived = true
	ac.BytesSent = true

	return ac
}
