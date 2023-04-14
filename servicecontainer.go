package main

import (
	"github.com/mehulgohil/shorti.fy/controllers"
	"github.com/mehulgohil/shorti.fy/interfaces"
	"github.com/mehulgohil/shorti.fy/pkg/algorithm/encoding"
	"github.com/mehulgohil/shorti.fy/pkg/algorithm/hashing"
	"github.com/mehulgohil/shorti.fy/services"
	"sync"
)

var (
	serviceContainerObj *serviceContainer
	containerOnce       sync.Once
)

type IServiceContainer interface {
	InjectHealthCheckController() controllers.HealthCheckController
	InjectShortifyWriterController(dbClient interfaces.IDataAccessLayer) controllers.ShortifyWriterController
	InjectShortifyReaderController(dbClient interfaces.IDataAccessLayer) controllers.ShortifyReaderController
}

type serviceContainer struct{}

func (sc *serviceContainer) InjectHealthCheckController() controllers.HealthCheckController {
	// injecting service layer in controller
	return controllers.HealthCheckController{IHealthCheckService: &services.HealthCheckService{}}
}

func (sc *serviceContainer) InjectShortifyWriterController(dbClient interfaces.IDataAccessLayer) controllers.ShortifyWriterController {
	// injecting service layer in controller
	return controllers.ShortifyWriterController{IShortifyWriterService: &services.ShortifyWriterService{
		IEncodingAlgorithm: &encoding.Base62Algorithm{}, //injecting base62 as the encoding algorithm
		IHashingAlgorithm:  &hashing.MD5Hash{},          //injecting md5 as hashing algorithm
		IDataAccessLayer:   dbClient,                    //injecting db client
	}}
}

func (sc *serviceContainer) InjectShortifyReaderController(dbClient interfaces.IDataAccessLayer) controllers.ShortifyReaderController {
	// injecting service layer in controller
	return controllers.ShortifyReaderController{IShortifyReaderService: &services.ShortifyReaderService{
		IEncodingAlgorithm: &encoding.Base62Algorithm{}, //injecting base62 as the encoding algorithm
		IHashingAlgorithm:  &hashing.MD5Hash{},          //injecting md5 as hashing algorithm
		IDataAccessLayer:   dbClient,                    //injecting db client
	}}
}

func ServiceContainer() IServiceContainer {
	if serviceContainerObj == nil {
		containerOnce.Do(func() {
			serviceContainerObj = &serviceContainer{}
		})
	}
	return serviceContainerObj
}
