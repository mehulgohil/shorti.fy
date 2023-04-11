package interfaces

import "github.com/mehulgohil/shorti.fy/models"

type IHealthCheckService interface {
	CheckHealthCheck() models.HealthCheckResponse
}
