package interfaces

import "github.com/mehulgohil/shorti.fy/redirect/models"

type IHealthCheckService interface {
	CheckHealthCheck() models.HealthCheckResponse
}
