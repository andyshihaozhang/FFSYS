package service

import (
	"gin-vue-admin/server/service/example"
	"gin-vue-admin/server/service/ffproduction"
	"gin-vue-admin/server/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup      system.ServiceGroup
	ExampleServiceGroup     example.ServiceGroup
	FfproductionServiceGroup ffproduction.ServiceGroup
}
