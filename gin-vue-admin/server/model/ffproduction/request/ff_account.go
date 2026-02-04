package request

import (
	"gin-vue-admin/server/model/common/request"
)

type FfAccountSearch struct {
	AccountName string `json:"accountName" form:"accountName"`
	Phone       string `json:"phone" form:"phone"`
	Address     string `json:"address" form:"address"`
	request.PageInfo
}
