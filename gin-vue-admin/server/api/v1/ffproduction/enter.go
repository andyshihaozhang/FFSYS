package ffproduction

import (
	"gin-vue-admin/server/service"
)

type ApiGroup struct {
	ProductApi
	AccountApi
	MaterialApi
	OperationApi
	OperationAssignmentApi
	EmployeeApi
}

var productService = service.ServiceGroupApp.FfproductionServiceGroup.ProductService
var accountService = service.ServiceGroupApp.FfproductionServiceGroup.AccountService
var materialService = service.ServiceGroupApp.FfproductionServiceGroup.MaterialService
var operationService = service.ServiceGroupApp.FfproductionServiceGroup.OperationService
var operationAssignmentService = service.ServiceGroupApp.FfproductionServiceGroup.OperationAssignmentService
var employeeService = service.ServiceGroupApp.FfproductionServiceGroup.EmployeeService


