package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/accesslog"
	"io"
	"os"
	"sync"
)

var (
	irisRouter *router
	routerOnce sync.Once
)

type IRouter interface {
	InitRouter() *iris.Application
}

type router struct{}

func (router *router) InitRouter() *iris.Application {
	app := iris.New()
	ac := makeAccessLog()
	app.UseRouter(ac.Handler)

	healthCheckController := ServiceContainer().InjectHealthCheckController()

	app.Get("/healthcheck", healthCheckController.CheckServerHealthCheck)

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
