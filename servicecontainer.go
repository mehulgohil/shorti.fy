package main

import (
	"github.com/mehulgohil/shorti.fy/controllers"
	"github.com/mehulgohil/shorti.fy/services"
	"sync"
)

var (
	serviceContainerObj *serviceContainer
	containerOnce       sync.Once
)

type IServiceContainer interface {
	InjectHealthCheckController() controllers.HealthCheckController
	InjectShortifyController() controllers.ShortifyController
}

type serviceContainer struct{}

func (sc *serviceContainer) InjectHealthCheckController() controllers.HealthCheckController {
	// injecting service layer in controller
	return controllers.HealthCheckController{IHealthCheckService: &services.HealthCheckService{}}
}

func (sc *serviceContainer) InjectShortifyController() controllers.ShortifyController {
	// injecting service layer in controller
	return controllers.ShortifyController{IShortifyService: &services.ShortifyService{}}
}

func ServiceContainer() IServiceContainer {
	if serviceContainerObj == nil {
		containerOnce.Do(func() {
			serviceContainerObj = &serviceContainer{}
		})
	}
	return serviceContainerObj
}
