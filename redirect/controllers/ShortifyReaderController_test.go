package controllers

import (
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/kataras/iris/v12/httptest"
	"github.com/kataras/iris/v12/x/errors"
	"github.com/mehulgohil/shorti.fy/redirect/interfaces/mocks"
	"github.com/mehulgohil/shorti.fy/redirect/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

var mockReaderController ShortifyReaderController

func TestShortifyReaderController_ReaderController(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ZapLogger, _ := zap.NewDevelopment()
	mockReaderService := mocks.NewMockIShortifyReaderService(ctrl)
	mockReaderController.IShortifyReaderService = mockReaderService
	mockReaderController.Logger = ZapLogger

	t.Run("negative: error in reader service", func(t *testing.T) {
		t.Parallel()
		mockReaderService.EXPECT().Reader(gomock.Any()).Return("", errors.New("mock error")).Times(1)

		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/v1/r123", nil)
		httptest.Do(w, r, mockReaderController.ReaderController)

		assert.Equal(t, 500, w.Code)

		var resp models.ErrorResponse
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		if err != nil {
			t.Error("error unmarshalling actual body")
		}

		assert.Equal(t, "mock error", resp.Error)
	})

	t.Run("positive", func(t *testing.T) {
		t.Parallel()
		mockReaderService.EXPECT().Reader(gomock.Any()).Return("http://mocklongurl", nil).Times(1)

		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/v1/r123", nil)
		httptest.Do(w, r, mockReaderController.ReaderController)

		assert.Equal(t, 301, w.Code)
	})
}
