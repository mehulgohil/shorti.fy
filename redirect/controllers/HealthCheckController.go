package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/mehulgohil/shorti.fy/redirect/interfaces"
)

type HealthCheckController struct {
	interfaces.IHealthCheckService
}

//	@Summary		Check HealthCheckStatus
//	@Description	Check Server Health
//	@Tags			healthcheck
//	@Produce		json
//	@Success		200	{object}	models.HealthCheckResponse
//	@Router			/healthcheck [get]
func (controller *HealthCheckController) CheckServerHealthCheck(ctx iris.Context) {
	_ = ctx.JSON(controller.CheckHealthCheck())
}
