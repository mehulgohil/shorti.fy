package main

import (
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
	"sync"
)

var (
	swaggerRouterObj  *swaggerRouter
	swaggerRouterOnce sync.Once
)

type ISwaggerRouter interface {
	InitSwaggerRouter(app *iris.Application)
}

type swaggerRouter struct{}

// InitSwaggerRouter initialize swagger route
func (s *swaggerRouter) InitSwaggerRouter(app *iris.Application) {
	swaggerConfig := &swagger.Config{
		URL:         "http://localhost:8081/swagger/doc.json",
		DeepLinking: true,
	}
	app.Get("/swagger/{any:path}", swagger.CustomWrapHandler(swaggerConfig, swaggerFiles.Handler))
}

func SwaggerRouter() ISwaggerRouter {
	if swaggerRouterObj == nil {
		swaggerRouterOnce.Do(func() {
			swaggerRouterObj = &swaggerRouter{}
		})
	}
	return swaggerRouterObj
}
