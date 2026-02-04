package ffproduction

import (
	"gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OperationRouter struct{}

func (e *OperationRouter) InitOperationRouter(Router *gin.RouterGroup) {
	operationRouter := Router.Group("ffProduct").Use(middleware.OperationRecord())
	operationRouterWithoutRecord := Router.Group("ffProduct")
	{
		operationRouter.POST("createOperation", operationApi.CreateOperation)
		operationRouter.PUT("updateOperation", operationApi.UpdateOperation)
		operationRouter.DELETE("deleteOperation", operationApi.DeleteOperation)
	}
	{
		operationRouterWithoutRecord.GET("getOperationList", operationApi.GetOperationList)
	}
}
