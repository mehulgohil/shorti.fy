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
	InjectShortifyController(dbClient interfaces.IDataAccessLayer) controllers.ShortifyController
}

type serviceContainer struct{}

func (sc *serviceContainer) InjectHealthCheckController() controllers.HealthCheckController {
	// injecting service layer in controller
	return controllers.HealthCheckController{IHealthCheckService: &services.HealthCheckService{}}
}

func (sc *serviceContainer) InjectShortifyController(dbClient interfaces.IDataAccessLayer) controllers.ShortifyController {
	// injecting service layer in controller
	return controllers.ShortifyController{IShortifyService: &services.ShortifyService{
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
