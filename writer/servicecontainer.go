package main

import (
	"github.com/mehulgohil/shorti.fy/common/interfaces"
	"github.com/mehulgohil/shorti.fy/common/pkg/algorithm/encoding"
	"github.com/mehulgohil/shorti.fy/common/pkg/algorithm/hashing"
	"github.com/mehulgohil/shorti.fy/writer/controllers"
	"github.com/mehulgohil/shorti.fy/writer/services"
	"sync"
)

var (
	serviceContainerObj *serviceContainer
	containerOnce       sync.Once
)

type IServiceContainer interface {
	InjectHealthCheckController() controllers.HealthCheckController
	InjectShortifyWriterController(dbClient interfaces.IDataAccessLayer) controllers.ShortifyWriterController
}

type serviceContainer struct{}

func (sc *serviceContainer) InjectHealthCheckController() controllers.HealthCheckController {
	// injecting service layer in controller
	return controllers.HealthCheckController{IHealthCheckService: &services.HealthCheckService{}}
}

func (sc *serviceContainer) InjectShortifyWriterController(dbClient interfaces.IDataAccessLayer) controllers.ShortifyWriterController {
	// injecting service layer in controller
	return controllers.ShortifyWriterController{
		IShortifyWriterService: &services.ShortifyWriterService{
			IEncodingAlgorithm: &encoding.Base62Algorithm{}, //injecting base62 as the encoding algorithm
			IHashingAlgorithm:  &hashing.MD5Hash{},          //injecting md5 as hashing algorithm
			IDataAccessLayer:   dbClient,                    //injecting db client
		},
		Logger: ZapLogger,
	}
}

func ServiceContainer() IServiceContainer {
	if serviceContainerObj == nil {
		containerOnce.Do(func() {
			serviceContainerObj = &serviceContainer{}
		})
	}
	return serviceContainerObj
}
