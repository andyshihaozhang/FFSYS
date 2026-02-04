package ffproduction

import (
	api "gin-vue-admin/server/api/v1"
)

type RouterGroup struct {
	ProductRouter
	AccountRouter
	MaterialRouter
	OperationRouter
	OperationAssignmentRouter
	EmployeeRouter
}

var (
	productApi    = api.ApiGroupApp.FfproductionApiGroup.ProductApi
	accountApi    = api.ApiGroupApp.FfproductionApiGroup.AccountApi
	materialApi   = api.ApiGroupApp.FfproductionApiGroup.MaterialApi
	operationApi  = api.ApiGroupApp.FfproductionApiGroup.OperationApi
	operationAssignmentApi = api.ApiGroupApp.FfproductionApiGroup.OperationAssignmentApi
	employeeApi   = api.ApiGroupApp.FfproductionApiGroup.EmployeeApi
)
