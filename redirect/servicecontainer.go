package main

import (
	"github.com/mehulgohil/shorti.fy/common/interfaces"
	"github.com/mehulgohil/shorti.fy/redirect/controllers"
	"github.com/mehulgohil/shorti.fy/redirect/services"
	"sync"
)

var (
	serviceContainerObj *serviceContainer
	containerOnce       sync.Once
)

type IServiceContainer interface {
	InjectHealthCheckController() controllers.HealthCheckController
	InjectShortifyReaderController(dbClient interfaces.IDataAccessLayer) controllers.ShortifyReaderController
}

type serviceContainer struct{}

func (sc *serviceContainer) InjectHealthCheckController() controllers.HealthCheckController {
	// injecting service layer in controller
	return controllers.HealthCheckController{IHealthCheckService: &services.HealthCheckService{}}
}

func (sc *serviceContainer) InjectShortifyReaderController(dbClient interfaces.IDataAccessLayer) controllers.ShortifyReaderController {
	// injecting service layer in controller
	return controllers.ShortifyReaderController{
		IShortifyReaderService: &services.ShortifyReaderService{
			IDataAccessLayer: dbClient, //injecting db client
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
