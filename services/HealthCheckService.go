package services

import "github.com/mehulgohil/shorti.fy/models"

type HealthCheckService struct{}

func (hc *HealthCheckService) CheckHealthCheck() models.HealthCheckResponse {
	return models.HealthCheckResponse{
		Status: "ok",
	}
}
