package ffproduction

import (
	"gin-vue-admin/server/middleware"

	"github.com/gin-gonic/gin"
)

type ProductRouter struct{}

func (e *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	productRouter := Router.Group("ffProduct").Use(middleware.OperationRecord())
	productRouterWithoutRecord := Router.Group("ffProduct")
	{
		productRouter.POST("createProduct", productApi.CreateProduct)
		productRouter.PUT("updateProduct", productApi.UpdateProduct)
		productRouter.DELETE("deleteProduct", productApi.DeleteProduct)
		productRouter.DELETE("deleteProductByIds", productApi.DeleteProductByIds)
	}
	{
		productRouterWithoutRecord.GET("getProduct", productApi.GetProduct)
		productRouterWithoutRecord.GET("getProductList", productApi.GetProductList)
	}
}
