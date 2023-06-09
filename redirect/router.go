package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/accesslog"
	"github.com/mehulgohil/shorti.fy/redirect/interfaces"
	"io"
	"os"
	"sync"
)

var (
	irisRouter *router
	routerOnce sync.Once
)

type IRouter interface {
	InitRouter(dbClient interfaces.IDataAccessLayer, redisClient interfaces.IRedisLayer) *iris.Application
}

type router struct{}

func (router *router) InitRouter(dbClient interfaces.IDataAccessLayer, redisClient interfaces.IRedisLayer) *iris.Application {
	app := iris.New()
	ac := makeAccessLog()
	app.UseRouter(ac.Handler)

	healthCheckController := ServiceContainer().InjectHealthCheckController()
	shortifyReaderController := ServiceContainer().InjectShortifyReaderController(dbClient, redisClient)

	app.Get("/healthcheck", healthCheckController.CheckServerHealthCheck)
	app.Get("/v1/{hashKey}", shortifyReaderController.ReaderController)

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
