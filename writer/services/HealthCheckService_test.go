package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHealthCheckService_CheckHealthCheck(t *testing.T) {
	t.Parallel()

	var mockHealthCheckService HealthCheckService

	t.Run("positive test", func(t *testing.T) {
		t.Parallel()
		resp := mockHealthCheckService.CheckHealthCheck()

		assert.Equal(t, "ok", resp.Status)
	})
}
