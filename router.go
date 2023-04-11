package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/accesslog"
	"io"
	"os"
	"sync"
)

var (
	irisApplication *router
	routerOnce      sync.Once
)

type IRouter interface {
	InitRouter() *iris.Application
}

type router struct{}

func (router *router) InitRouter() *iris.Application {
	app := iris.New()
	ac := makeAccessLog()
	app.UseRouter(ac.Handler)

	app.Get("/", func(context iris.Context) {
		_ = context.JSON("Hi")
	})

	return app
}

func Router() IRouter {
	if irisApplication == nil {
		routerOnce.Do(func() {
			irisApplication = &router{}
		})
	}
	return irisApplication
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
