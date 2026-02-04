package ffproduction

import (
	"gin-vue-admin/server/global"
	"gin-vue-admin/server/model/common/response"
	"gin-vue-admin/server/model/ffproduction"
	"gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MaterialApi struct{}

// CreateMaterial
// @Tags      FfMaterial
// @Summary   创建物料
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      ffproduction.FfMaterial            true  "物料信息"
// @Success   200   {object}  response.Response{msg=string}  "创建物料"
// @Router    /ffProduct/createMaterial [post]
func (e *MaterialApi) CreateMaterial(c *gin.Context) {
	var material ffproduction.FfMaterial
	err := c.ShouldBindJSON(&material)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = materialService.CreateMaterial(material)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteMaterial
// @Tags      FfMaterial
// @Summary   删除物料
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      ffproduction.FfMaterial            true  "物料ID"
// @Success   200   {object}  response.Response{msg=string}  "删除物料"
// @Router    /ffProduct/deleteMaterial [delete]
func (e *MaterialApi) DeleteMaterial(c *gin.Context) {
	var material ffproduction.FfMaterial
	err := c.ShouldBindJSON(&material)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(material.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = materialService.DeleteMaterial(material)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateMaterial
// @Tags      FfMaterial
// @Summary   更新物料信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      ffproduction.FfMaterial            true  "物料ID, 物料信息"
// @Success   200   {object}  response.Response{msg=string}  "更新物料信息"
// @Router    /ffProduct/updateMaterial [put]
func (e *MaterialApi) UpdateMaterial(c *gin.Context) {
	var material ffproduction.FfMaterial
	err := c.ShouldBindJSON(&material)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(material.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = materialService.UpdateMaterial(&material)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetMaterialList
// @Tags      FfMaterial
// @Summary   获取物料列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     ffproduction.FfMaterialSearch                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "获取物料列表"
// @Router    /ffProduct/getMaterialList [get]
func (e *MaterialApi) GetMaterialList(c *gin.Context) {
	var pageInfo ffproduction.FfMaterialSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	materialList, total, err := materialService.GetMaterialList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     materialList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}
