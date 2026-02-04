package ffproduction

import (
	"gin-vue-admin/server/global"
	"gin-vue-admin/server/model/common/response"
	"gin-vue-admin/server/model/ffproduction"
	"gin-vue-admin/server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EmployeeApi struct{}

// CreateEmployee
// @Tags      FfEmployee
// @Summary   创建员工
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      ffproduction.FfEmployee            true  "员工信息"
// @Success   200   {object}  response.Response{msg=string}  "创建员工"
// @Router    /ffProduct/createEmployee [post]
func (e *EmployeeApi) CreateEmployee(c *gin.Context) {
	var employee ffproduction.FfEmployee
	err := c.ShouldBindJSON(&employee)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = employeeService.CreateEmployee(employee)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteEmployee
// @Tags      FfEmployee
// @Summary   删除员工
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      ffproduction.FfEmployee            true  "员工ID"
// @Success   200   {object}  response.Response{msg=string}  "删除员工"
// @Router    /ffProduct/deleteEmployee [delete]
func (e *EmployeeApi) DeleteEmployee(c *gin.Context) {
	var employee ffproduction.FfEmployee
	err := c.ShouldBindJSON(&employee)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(employee.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = employeeService.DeleteEmployee(employee)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateEmployee
// @Tags      FfEmployee
// @Summary   更新员工信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      ffproduction.FfEmployee            true  "员工ID, 员工信息"
// @Success   200   {object}  response.Response{msg=string}  "更新员工信息"
// @Router    /ffProduct/updateEmployee [put]
func (e *EmployeeApi) UpdateEmployee(c *gin.Context) {
	var employee ffproduction.FfEmployee
	err := c.ShouldBindJSON(&employee)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(employee.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = employeeService.UpdateEmployee(&employee)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetEmployeeList
// @Tags      FfEmployee
// @Summary   分页获取员工列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     ffproduction.FfEmployeeSearch                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取员工列表"
// @Router    /ffProduct/getEmployeeList [get]
func (e *EmployeeApi) GetEmployeeList(c *gin.Context) {
	var pageInfo ffproduction.FfEmployeeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	employeeList, total, err := employeeService.GetEmployeeList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     employeeList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}
