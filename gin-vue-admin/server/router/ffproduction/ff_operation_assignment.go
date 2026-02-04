package ffproduction

import (
	"gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OperationAssignmentRouter struct{}

func (e *OperationAssignmentRouter) InitOperationAssignmentRouter(Router *gin.RouterGroup) {
	assignmentRouter := Router.Group("ffProduct").Use(middleware.OperationRecord())
	assignmentRouterWithoutRecord := Router.Group("ffProduct")
	{
		assignmentRouter.POST("createOperationAssignment", operationAssignmentApi.CreateOperationAssignment)
		assignmentRouter.PUT("updateOperationAssignment", operationAssignmentApi.UpdateOperationAssignment)
		assignmentRouter.DELETE("deleteOperationAssignment", operationAssignmentApi.DeleteOperationAssignment)
	}
	{
		assignmentRouterWithoutRecord.GET("getOperationAssignmentList", operationAssignmentApi.GetOperationAssignmentList)
	}
}
