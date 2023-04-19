package interfaces

import "github.com/mehulgohil/shorti.fy/writer/models"

type IHealthCheckService interface {
	CheckHealthCheck() models.HealthCheckResponse
}
