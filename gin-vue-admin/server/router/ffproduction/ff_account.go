package ffproduction

import (
	"gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AccountRouter struct{}

func (e *AccountRouter) InitAccountRouter(Router *gin.RouterGroup) {
	accountRouter := Router.Group("ffProduct").Use(middleware.OperationRecord())
	accountRouterWithoutRecord := Router.Group("ffProduct")
	{
		accountRouter.POST("createAccount", accountApi.CreateAccount)
		accountRouter.PUT("updateAccount", accountApi.UpdateAccount)
		accountRouter.DELETE("deleteAccount", accountApi.DeleteAccount)
	}
	{
		accountRouterWithoutRecord.GET("getAccount", accountApi.GetAccount)
		accountRouterWithoutRecord.GET("getAccountList", accountApi.GetAccountList)
		accountRouterWithoutRecord.GET("getAccountOptions", accountApi.GetAccountOptions)
	}
}
