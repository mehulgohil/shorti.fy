package services

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/kataras/iris/v12/x/errors"
	"github.com/mehulgohil/shorti.fy/writer/interfaces/mocks"
	"github.com/mehulgohil/shorti.fy/writer/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	mockError = "mock error"
)

func TestShortifyWriterService_Writer(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("error in get item", func(t *testing.T) {
		t.Parallel()
		mockEncodingAlgo := mocks.NewMockIEncodingAlgorithm(ctrl)
		mockHashingAlgo := mocks.NewMockIHashingAlgorithm(ctrl)
		mockDataAccess := mocks.NewMockIDataAccessLayer(ctrl)

		var mockWriterService = ShortifyWriterService{
			mockEncodingAlgo,
			mockHashingAlgo,
			mockDataAccess,
		}

		mockHashingAlgo.EXPECT().Hash(gomock.Any()).Return("mockhash").Times(1)
		mockEncodingAlgo.EXPECT().Encode(gomock.Any()).Return("mockencode").Times(1)
		mockDataAccess.EXPECT().GetItem(gomock.Any()).Return(models.URLTable{}, errors.New(mockError)).Times(1)

		_, err := mockWriterService.Writer("mock long", "mock email")
		assert.Equal(t, mockError, err.Error())
	})

	t.Run("error in get item, 2nd iteration", func(t *testing.T) {
		t.Parallel()
		mockEncodingAlgo := mocks.NewMockIEncodingAlgorithm(ctrl)
		mockHashingAlgo := mocks.NewMockIHashingAlgorithm(ctrl)
		mockDataAccess := mocks.NewMockIDataAccessLayer(ctrl)

		var mockWriterService = ShortifyWriterService{
			mockEncodingAlgo,
			mockHashingAlgo,
			mockDataAccess,
		}

		mockHashingAlgo.EXPECT().Hash(gomock.Any()).Return("mockhash").Times(2)
		mockEncodingAlgo.EXPECT().Encode(gomock.Any()).Return("mockencode").Times(2)
		mockDataAccess.EXPECT().GetItem(gomock.Any()).Return(models.URLTable{HashKey: "mockEnc"}, nil).Times(1)
		mockDataAccess.EXPECT().GetItem(gomock.Any()).Return(models.URLTable{HashKey: "mockEnc"}, errors.New(mockError)).Times(1)

		_, err := mockWriterService.Writer("mock long", "mock email")
		assert.Equal(t, mockError, err.Error())
	})

	t.Run("error in save item", func(t *testing.T) {
		t.Parallel()
		mockEncodingAlgo := mocks.NewMockIEncodingAlgorithm(ctrl)
		mockHashingAlgo := mocks.NewMockIHashingAlgorithm(ctrl)
		mockDataAccess := mocks.NewMockIDataAccessLayer(ctrl)

		var mockWriterService = ShortifyWriterService{
			mockEncodingAlgo,
			mockHashingAlgo,
			mockDataAccess,
		}

		mockHashingAlgo.EXPECT().Hash(gomock.Any()).Return("mockHash").Times(1)
		mockEncodingAlgo.EXPECT().Encode(gomock.Any()).Return("mockEncode").Times(1)
		mockDataAccess.EXPECT().GetItem(gomock.Any()).Return(models.URLTable{}, nil).Times(1)
		mockDataAccess.EXPECT().SaveItem(gomock.Any()).Return(errors.New(mockError)).Times(1)

		_, err := mockWriterService.Writer("mock long", "mock email")
		assert.Equal(t, mockError, err.Error())
	})

	t.Run("positive test", func(t *testing.T) {
		t.Parallel()
		mockEncodingAlgo := mocks.NewMockIEncodingAlgorithm(ctrl)
		mockHashingAlgo := mocks.NewMockIHashingAlgorithm(ctrl)
		mockDataAccess := mocks.NewMockIDataAccessLayer(ctrl)

		var mockWriterService = ShortifyWriterService{
			mockEncodingAlgo,
			mockHashingAlgo,
			mockDataAccess,
		}

		mockHashingAlgo.EXPECT().Hash(gomock.Any()).Return("mockHash").Times(1)
		mockEncodingAlgo.EXPECT().Encode(gomock.Any()).Return("mockEncode").Times(1)
		mockDataAccess.EXPECT().GetItem(gomock.Any()).Return(models.URLTable{}, nil).Times(1)
		mockDataAccess.EXPECT().SaveItem(gomock.Any()).Return(nil).Times(1)

		res, err := mockWriterService.Writer("mock long", "mock email")
		assert.Equal(t, nil, err)
		assert.Equal(t, fmt.Sprintf("http://localhost:80/v1/%s", "mockEnc"), res)
	})
}
