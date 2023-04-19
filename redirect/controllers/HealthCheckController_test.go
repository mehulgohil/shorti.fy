package controllers

import (
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/kataras/iris/v12/httptest"
	"github.com/mehulgohil/shorti.fy/redirect/interfaces/mocks"
	"github.com/mehulgohil/shorti.fy/redirect/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

var mockController HealthCheckController

func TestHealthCheckController_CheckServerHealthCheck(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockIHealthCheckService(ctrl)
	mockController.IHealthCheckService = mockService

	expectedServiceResponse := models.HealthCheckResponse{Status: "ok"}

	t.Run("positive test", func(t *testing.T) {
		t.Parallel()
		mockService.EXPECT().CheckHealthCheck().Return(expectedServiceResponse)
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/healthcheck", nil)

		httptest.Do(w, r, mockController.CheckServerHealthCheck)

		assert.Equal(t, 200, w.Code)

		var resp models.HealthCheckResponse
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		if err != nil {
			t.Error("error unmarshalling actual body")
		}

		assert.Equal(t, expectedServiceResponse, resp)
	})
}
