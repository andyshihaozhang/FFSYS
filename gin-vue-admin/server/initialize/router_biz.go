package initialize

import (
	"gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
)

func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.RouterGroupApp
}

func initBizRouter(routers ...*gin.RouterGroup) {
	privateGroup := routers[0]
	publicGroup := routers[1]

	ffproductionRouter := router.RouterGroupApp.Ffproduction
	{
		ffproductionRouter.InitProductRouter(privateGroup)
		ffproductionRouter.InitAccountRouter(privateGroup)
		ffproductionRouter.InitMaterialRouter(privateGroup)
		ffproductionRouter.InitOperationRouter(privateGroup)
		ffproductionRouter.InitOperationAssignmentRouter(privateGroup)
		ffproductionRouter.InitEmployeeRouter(privateGroup)
	}

	holder(publicGroup, privateGroup)
}
