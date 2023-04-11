package services

type HealthCheckService struct{}

func (hc *HealthCheckService) CheckHealthCheck() string {
	return "ok"
}
