package interfaces

import "github.com/mehulgohil/shorti.fy/common/models"

type IHealthCheckService interface {
	CheckHealthCheck() models.HealthCheckResponse
}
