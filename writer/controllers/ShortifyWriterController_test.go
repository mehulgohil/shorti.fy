package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/kataras/iris/v12/httptest"
	"github.com/mehulgohil/shorti.fy/writer/interfaces/mocks"
	"github.com/mehulgohil/shorti.fy/writer/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

var mockWriterController ShortifyWriterController

func TestShortifyWriterController_WriterController(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ZapLogger, _ := zap.NewDevelopment()
	mockWriterService := mocks.NewMockIShortifyWriterService(ctrl)
	mockWriterController.IShortifyWriterService = mockWriterService
	mockWriterController.Logger = ZapLogger

	mockWriteReqBody := models.WriterRequest{
		LongURL:   "http://mocklongurl",
		UserEmail: "mock@mock.com",
	}
	requestByte, _ := json.Marshal(mockWriteReqBody)

	mockShortURL := "http://shorturl"

	t.Run("negative: invalid body error", func(t *testing.T) {
		t.Parallel()

		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/v1/shorten", nil)
		httptest.Do(w, r, mockWriterController.WriterController)

		assert.Equal(t, 400, w.Code)

		var resp models.ErrorResponse
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		if err != nil {
			t.Error("error unmarshalling actual body")
		}

		assert.Equal(t, "EOF", resp.Error)
	})

	t.Run("negative: error in writer service", func(t *testing.T) {
		t.Parallel()
		mockWriterService.EXPECT().Writer(gomock.Any(), gomock.Any()).Return(mockShortURL, errors.New("mock error")).Times(1)

		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/v1/shorten", bytes.NewReader(requestByte))
		httptest.Do(w, r, mockWriterController.WriterController)

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
		mockWriterService.EXPECT().Writer(gomock.Any(), gomock.Any()).Return("http://shorturl", nil).Times(1)

		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/v1/shorten", bytes.NewReader(requestByte))
		httptest.Do(w, r, mockWriterController.WriterController)

		assert.Equal(t, 200, w.Code)

		var resp models.WriterResponse
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		if err != nil {
			t.Error("error unmarshalling actual body")
		}

		assert.Equal(t, mockWriteReqBody.LongURL, resp.LongURL)
		assert.Equal(t, mockShortURL, resp.ShortURL)
	})
}
