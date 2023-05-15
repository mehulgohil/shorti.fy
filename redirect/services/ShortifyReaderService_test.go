package services

import (
	"github.com/golang/mock/gomock"
	"github.com/kataras/iris/v12/x/errors"
	"github.com/mehulgohil/shorti.fy/redirect/interfaces/mocks"
	"github.com/mehulgohil/shorti.fy/redirect/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"sync"
	"testing"
)

const (
	mockError = "mock error"
)

func TestShortifyReaderService_Reader(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ZapLogger, _ := zap.NewDevelopment()

	mockRedisLongURL := "mockRedisLongURL"
	mockDBLongURL := "mockDBLongURL"

	t.Run("positive test - got values from redis", func(t *testing.T) {
		t.Parallel()
		mockDataAccess := mocks.NewMockIDataAccessLayer(ctrl)
		mockRedis := mocks.NewMockIRedisLayer(ctrl)

		var mockReaderService = ShortifyReaderService{
			mockDataAccess,
			mockRedis,
			ZapLogger,
			sync.Mutex{},
		}

		mockRedis.EXPECT().GetKeyValue(gomock.Any()).Return(mockRedisLongURL, nil).Times(1)

		//defer statement mock
		mockDataAccess.EXPECT().GetItem(gomock.Any()).Return(models.URLTable{}, errors.New(mockError)).AnyTimes()

		reader, err := mockReaderService.Reader("shortUrl")
		assert.Equal(t, nil, err)
		assert.Equal(t, reader, mockRedisLongURL)
	})

	// error getting value from redis
	// getting from db
	t.Run("error in get value from DB", func(t *testing.T) {
		t.Parallel()
		mockDataAccess := mocks.NewMockIDataAccessLayer(ctrl)
		mockRedis := mocks.NewMockIRedisLayer(ctrl)

		var mockReaderService = ShortifyReaderService{
			mockDataAccess,
			mockRedis,
			ZapLogger,
			sync.Mutex{},
		}

		mockRedis.EXPECT().GetKeyValue(gomock.Any()).Return(mockRedisLongURL, errors.New(mockError)).Times(1)
		mockDataAccess.EXPECT().GetItem(gomock.Any()).Return(models.URLTable{}, errors.New(mockError)).Times(1)

		//defer statement mock
		mockDataAccess.EXPECT().GetItem(gomock.Any()).Return(models.URLTable{}, errors.New(mockError)).AnyTimes()

		_, err := mockReaderService.Reader("shortUrl")
		assert.Equal(t, mockError, err.Error())
	})

	t.Run("empty value from DB error", func(t *testing.T) {
		t.Parallel()
		mockDataAccess := mocks.NewMockIDataAccessLayer(ctrl)
		mockRedis := mocks.NewMockIRedisLayer(ctrl)

		var mockReaderService = ShortifyReaderService{
			mockDataAccess,
			mockRedis,
			ZapLogger,
			sync.Mutex{},
		}

		mockRedis.EXPECT().GetKeyValue(gomock.Any()).Return(mockRedisLongURL, errors.New(mockError)).Times(1)
		mockDataAccess.EXPECT().GetItem(gomock.Any()).Return(models.URLTable{}, nil).Times(1)

		//defer statement mock
		mockDataAccess.EXPECT().GetItem(gomock.Any()).Return(models.URLTable{}, errors.New(mockError)).AnyTimes()

		_, err := mockReaderService.Reader("shortUrl")
		assert.Equal(t, longUrlNotFoundErrorMsg, err.Error())
	})

	t.Run("positive test", func(t *testing.T) {
		t.Parallel()
		mockDataAccess := mocks.NewMockIDataAccessLayer(ctrl)
		mockRedis := mocks.NewMockIRedisLayer(ctrl)

		var mockReaderService = ShortifyReaderService{
			mockDataAccess,
			mockRedis,
			ZapLogger,
			sync.Mutex{},
		}

		mockRedis.EXPECT().GetKeyValue(gomock.Any()).Return(mockRedisLongURL, errors.New(mockError)).Times(1)
		mockDataAccess.EXPECT().GetItem(gomock.Any()).Return(models.URLTable{LongURL: mockDBLongURL}, nil).Times(1)

		//defer statement mock
		mockDataAccess.EXPECT().GetItem(gomock.Any()).Return(models.URLTable{}, errors.New(mockError)).AnyTimes()
		// goroutine cache data mock
		mockRedis.EXPECT().SetKeyValue(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

		longUrl, err := mockReaderService.Reader("shortUrl")
		assert.Equal(t, nil, err)
		assert.Equal(t, mockDBLongURL, longUrl)
	})
}

func TestShortifyReaderService_incrementHitCount(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ZapLogger, _ := zap.NewDevelopment()

	t.Run("error in get item", func(t *testing.T) {
		t.Parallel()
		mockDataAccess := mocks.NewMockIDataAccessLayer(ctrl)
		mockRedis := mocks.NewMockIRedisLayer(ctrl)

		var mockReaderService = ShortifyReaderService{
			mockDataAccess,
			mockRedis,
			ZapLogger,
			sync.Mutex{},
		}

		mockDataAccess.EXPECT().GetItem(gomock.Any()).Return(models.URLTable{}, errors.New(mockError)).Times(1)
		err := mockReaderService.incrementHitCount("shortURL")

		assert.Equal(t, mockError, err.Error())
	})

	t.Run("error in save item", func(t *testing.T) {
		t.Parallel()
		mockDataAccess := mocks.NewMockIDataAccessLayer(ctrl)
		mockRedis := mocks.NewMockIRedisLayer(ctrl)

		var mockReaderService = ShortifyReaderService{
			mockDataAccess,
			mockRedis,
			ZapLogger,
			sync.Mutex{},
		}

		mockDataAccess.EXPECT().GetItem(gomock.Any()).Return(models.URLTable{}, nil).Times(1)
		mockDataAccess.EXPECT().SaveItem(gomock.Any()).Return(errors.New(mockError)).Times(1)
		err := mockReaderService.incrementHitCount("shortURL")

		assert.Equal(t, mockError, err.Error())
	})

	t.Run("positive test", func(t *testing.T) {
		t.Parallel()
		mockDataAccess := mocks.NewMockIDataAccessLayer(ctrl)
		mockRedis := mocks.NewMockIRedisLayer(ctrl)

		var mockReaderService = ShortifyReaderService{
			mockDataAccess,
			mockRedis,
			ZapLogger,
			sync.Mutex{},
		}

		mockDataAccess.EXPECT().GetItem(gomock.Any()).Return(models.URLTable{}, nil).Times(1)
		mockDataAccess.EXPECT().SaveItem(gomock.Any()).Return(nil).Times(1)
		err := mockReaderService.incrementHitCount("shortURL")

		assert.Equal(t, nil, err)
	})
}
