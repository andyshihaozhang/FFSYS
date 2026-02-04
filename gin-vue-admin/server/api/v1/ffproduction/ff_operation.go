package ffproduction

import (
	"gin-vue-admin/server/global"
	"gin-vue-admin/server/model/common/response"
	"gin-vue-admin/server/model/ffproduction"
	"gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type OperationApi struct{}

// CreateOperation
// @Tags      FfOperation
// @Summary   创建工序
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      ffproduction.FfOperation            true  "工序信息"
// @Success   200   {object}  response.Response{msg=string}  "创建工序"
// @Router    /ffProduct/createOperation [post]
func (e *OperationApi) CreateOperation(c *gin.Context) {
	var operation ffproduction.FfOperation
	err := c.ShouldBindJSON(&operation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = operationService.CreateOperation(operation)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteOperation
// @Tags      FfOperation
// @Summary   删除工序
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      ffproduction.FfOperation            true  "工序ID"
// @Success   200   {object}  response.Response{msg=string}  "删除工序"
// @Router    /ffProduct/deleteOperation [delete]
func (e *OperationApi) DeleteOperation(c *gin.Context) {
	var operation ffproduction.FfOperation
	err := c.ShouldBindJSON(&operation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(operation.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = operationService.DeleteOperation(operation)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateOperation
// @Tags      FfOperation
// @Summary   更新工序信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      ffproduction.FfOperation            true  "工序ID, 工序信息"
// @Success   200   {object}  response.Response{msg=string}  "更新工序信息"
// @Router    /ffProduct/updateOperation [put]
func (e *OperationApi) UpdateOperation(c *gin.Context) {
	var operation ffproduction.FfOperation
	err := c.ShouldBindJSON(&operation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(operation.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = operationService.UpdateOperation(&operation)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetOperationList
// @Tags      FfOperation
// @Summary   获取工序列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     ffproduction.FfOperationSearch                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "获取工序列表"
// @Router    /ffProduct/getOperationList [get]
func (e *OperationApi) GetOperationList(c *gin.Context) {
	var pageInfo ffproduction.FfOperationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	operationList, total, err := operationService.GetOperationList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     operationList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}
