package ffproduction

import (
	"gin-vue-admin/server/global"
	"gin-vue-admin/server/model/common/response"
	"gin-vue-admin/server/model/ffproduction"
	"gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type OperationAssignmentApi struct{}

// CreateOperationAssignment
// @Tags      FfOperationAssignment
// @Summary   创建工序分配
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      ffproduction.FfOperationAssignment            true  "工序分配信息"
// @Success   200   {object}  response.Response{msg=string}  "创建工序分配"
// @Router    /ffProduct/createOperationAssignment [post]
func (e *OperationAssignmentApi) CreateOperationAssignment(c *gin.Context) {
	var assignment ffproduction.FfOperationAssignment
	err := c.ShouldBindJSON(&assignment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = operationAssignmentService.CreateOperationAssignment(assignment)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteOperationAssignment
// @Tags      FfOperationAssignment
// @Summary   删除工序分配
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      ffproduction.FfOperationAssignment            true  "工序分配ID"
// @Success   200   {object}  response.Response{msg=string}  "删除工序分配"
// @Router    /ffProduct/deleteOperationAssignment [delete]
func (e *OperationAssignmentApi) DeleteOperationAssignment(c *gin.Context) {
	var assignment ffproduction.FfOperationAssignment
	err := c.ShouldBindJSON(&assignment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(assignment.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = operationAssignmentService.DeleteOperationAssignment(assignment)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateOperationAssignment
// @Tags      FfOperationAssignment
// @Summary   更新工序分配信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      ffproduction.FfOperationAssignment            true  "工序分配ID, 工序分配信息"
// @Success   200   {object}  response.Response{msg=string}  "更新工序分配信息"
// @Router    /ffProduct/updateOperationAssignment [put]
func (e *OperationAssignmentApi) UpdateOperationAssignment(c *gin.Context) {
	var assignment ffproduction.FfOperationAssignment
	err := c.ShouldBindJSON(&assignment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(assignment.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = operationAssignmentService.UpdateOperationAssignment(&assignment)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetOperationAssignmentList
// @Tags      FfOperationAssignment
// @Summary   获取工序分配列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     ffproduction.FfOperationAssignmentSearch                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "获取工序分配列表"
// @Router    /ffProduct/getOperationAssignmentList [get]
func (e *OperationAssignmentApi) GetOperationAssignmentList(c *gin.Context) {
	var pageInfo ffproduction.FfOperationAssignmentSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	assignmentList, total, err := operationAssignmentService.GetOperationAssignmentList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     assignmentList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}
