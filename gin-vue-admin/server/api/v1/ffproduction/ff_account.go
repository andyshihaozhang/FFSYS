package ffproduction

import (
	"gin-vue-admin/server/global"
	"gin-vue-admin/server/model/common/response"
	"gin-vue-admin/server/model/ffproduction"
	ffRes "gin-vue-admin/server/model/ffproduction/response"
	"gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AccountApi struct{}

// CreateAccount
// @Tags      FfAccount
// @Summary   创建客户
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      ffproduction.FfAccount            true  "客户信息"
// @Success   200   {object}  response.Response{msg=string}  "创建客户"
// @Router    /ffProduct/createAccount [post]
func (e *AccountApi) CreateAccount(c *gin.Context) {
	var account ffproduction.FfAccount
	err := c.ShouldBindJSON(&account)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = accountService.CreateAccount(account)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteAccount
// @Tags      FfAccount
// @Summary   删除客户
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      ffproduction.FfAccount            true  "客户ID"
// @Success   200   {object}  response.Response{msg=string}  "删除客户"
// @Router    /ffProduct/deleteAccount [delete]
func (e *AccountApi) DeleteAccount(c *gin.Context) {
	var account ffproduction.FfAccount
	err := c.ShouldBindJSON(&account)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(account.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = accountService.DeleteAccount(account)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateAccount
// @Tags      FfAccount
// @Summary   更新客户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      ffproduction.FfAccount            true  "客户ID, 客户信息"
// @Success   200   {object}  response.Response{msg=string}  "更新客户信息"
// @Router    /ffProduct/updateAccount [put]
func (e *AccountApi) UpdateAccount(c *gin.Context) {
	var account ffproduction.FfAccount
	err := c.ShouldBindJSON(&account)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(account.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = accountService.UpdateAccount(&account)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetAccount
// @Tags      FfAccount
// @Summary   获取单一客户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     ffproduction.FfAccount                                                true  "客户ID"
// @Success   200   {object}  response.Response{data=ffRes.FfAccountResponse,msg=string}  "获取单一客户信息"
// @Router    /ffProduct/getAccount [get]
func (e *AccountApi) GetAccount(c *gin.Context) {
	var account ffproduction.FfAccount
	err := c.ShouldBindQuery(&account)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(account.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data, err := accountService.GetAccount(account.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(ffRes.FfAccountResponse{Account: data}, "获取成功", c)
}

// GetAccountList
// @Tags      FfAccount
// @Summary   分页获取客户列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     ffproduction.FfAccountSearch                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取客户列表"
// @Router    /ffProduct/getAccountList [get]
func (e *AccountApi) GetAccountList(c *gin.Context) {
	var pageInfo ffproduction.FfAccountSearch
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
	accountList, total, err := accountService.GetAccountList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     accountList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetAccountOptions
// @Tags      FfAccount
// @Summary   获取客户下拉框选项
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{data=[]map[string]interface{},msg=string}  "获取客户下拉框选项"
// @Router    /ffProduct/getAccountOptions [get]
func (e *AccountApi) GetAccountOptions(c *gin.Context) {
	accountList, err := accountService.GetAllAccounts()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	options := make([]map[string]interface{}, len(accountList))
	for i, account := range accountList {
		options[i] = map[string]interface{}{
			"label": account.AccountName,
			"value": account.ID,
		}
	}
	response.OkWithDetailed(options, "获取成功", c)
}
