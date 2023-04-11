package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/mehulgohil/shorti.fy/interfaces"
)

type HealthCheckController struct {
	interfaces.IHealthCheckService
}

func (controller *HealthCheckController) CheckServerHealthCheck(ctx iris.Context) {
	_ = ctx.JSON(controller.CheckHealthCheck())
}
