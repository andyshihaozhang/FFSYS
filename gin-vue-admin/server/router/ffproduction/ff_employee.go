package ffproduction

import (
	"gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EmployeeRouter struct{}

func (e *EmployeeRouter) InitEmployeeRouter(Router *gin.RouterGroup) {
	employeeRouter := Router.Group("ffProduct").Use(middleware.OperationRecord())
	employeeRouterWithoutRecord := Router.Group("ffProduct")
	{
		employeeRouter.POST("createEmployee", employeeApi.CreateEmployee)
		employeeRouter.PUT("updateEmployee", employeeApi.UpdateEmployee)
		employeeRouter.DELETE("deleteEmployee", employeeApi.DeleteEmployee)
	}
	{
		employeeRouterWithoutRecord.GET("getEmployeeList", employeeApi.GetEmployeeList)
	}
}
