package ffproduction

import (
	"gin-vue-admin/server/global"
	"gin-vue-admin/server/model/common/request"
	"gin-vue-admin/server/model/common/response"
	"gin-vue-admin/server/model/ffproduction"
	ffRes "gin-vue-admin/server/model/ffproduction/response"
	"gin-vue-admin/server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProductApi struct{}

// CreateProduct
// @Tags      FfProduct
// @Summary   创建产品
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      ffproduction.FfProduct            true  "产品信息"
// @Success   200   {object}  response.Response{msg=string}  "创建产品"
// @Router    /ffproduction/product [post]
func (e *ProductApi) CreateProduct(c *gin.Context) {
	var product ffproduction.FfProduct
	err := c.ShouldBindJSON(&product)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = productService.CreateProduct(product)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteProduct
// @Tags      FfProduct
// @Summary   删除产品
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      ffproduction.FfProduct            true  "产品ID"
// @Success   200   {object}  response.Response{msg=string}  "删除产品"
// @Router    /ffproduction/product [delete]
func (e *ProductApi) DeleteProduct(c *gin.Context) {
	var product ffproduction.FfProduct
	err := c.ShouldBindJSON(&product)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(product.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = productService.DeleteProduct(product)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteProductByIds
// @Tags      FfProduct
// @Summary   批量删除产品
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.IdsReq            true  "产品ID数组"
// @Success   200   {object}  response.Response{msg=string}  "批量删除产品"
// @Router    /ffproduction/productByIds [delete]
func (e *ProductApi) DeleteProductByIds(c *gin.Context) {
	var ids request.IdsReq
	err := c.ShouldBindJSON(&ids)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uintIds := make([]uint, len(ids.Ids))
	for i, id := range ids.Ids {
		uintIds[i] = uint(id)
	}
	err = productService.DeleteProductByIds(uintIds)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateProduct
// @Tags      FfProduct
// @Summary   更新产品信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      ffproduction.FfProduct            true  "产品ID, 产品信息"
// @Success   200   {object}  response.Response{msg=string}  "更新产品信息"
// @Router    /ffproduction/product [put]
func (e *ProductApi) UpdateProduct(c *gin.Context) {
	var product ffproduction.FfProduct
	err := c.ShouldBindJSON(&product)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(product.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = productService.UpdateProduct(&product)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetProduct
// @Tags      FfProduct
// @Summary   获取单一产品信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     ffproduction.FfProduct                                                true  "产品ID"
// @Success   200   {object}  response.Response{data=ffRes.FfProductResponse,msg=string}  "获取单一产品信息"
// @Router    /ffproduction/product [get]
func (e *ProductApi) GetProduct(c *gin.Context) {
	var product ffproduction.FfProduct
	err := c.ShouldBindQuery(&product)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(product.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data, err := productService.GetProduct(product.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(ffRes.FfProductResponse{Product: data}, "获取成功", c)
}

// GetProductList
// @Tags      FfProduct
// @Summary   分页获取产品列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     ffproduction.FfProductSearch                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取产品列表"
// @Router    /ffproduction/productList [get]
func (e *ProductApi) GetProductList(c *gin.Context) {
	var pageInfo ffproduction.FfProductSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	productList, total, err := productService.GetProductList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     productList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}
