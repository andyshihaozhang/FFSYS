package ffproduction

import (
	"gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MaterialRouter struct{}

func (e *MaterialRouter) InitMaterialRouter(Router *gin.RouterGroup) {
	materialRouter := Router.Group("ffProduct").Use(middleware.OperationRecord())
	materialRouterWithoutRecord := Router.Group("ffProduct")
	{
		materialRouter.POST("createMaterial", materialApi.CreateMaterial)
		materialRouter.PUT("updateMaterial", materialApi.UpdateMaterial)
		materialRouter.DELETE("deleteMaterial", materialApi.DeleteMaterial)
	}
	{
		materialRouterWithoutRecord.GET("getMaterialList", materialApi.GetMaterialList)
	}
}
