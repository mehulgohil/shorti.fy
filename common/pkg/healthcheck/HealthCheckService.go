package healthcheck

import "github.com/mehulgohil/shorti.fy/common/models"

type HealthCheckService struct{}

func (hc *HealthCheckService) CheckHealthCheck() models.HealthCheckResponse {
	return models.HealthCheckResponse{
		Status: "ok",
	}
}
