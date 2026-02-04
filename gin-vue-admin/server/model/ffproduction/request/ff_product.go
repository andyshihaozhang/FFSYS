package request

import (
	"gin-vue-admin/server/model/common/request"
)

type FfProductSearch struct {
	ProductCode  string `json:"productCode" form:"productCode"`
	ProductName  string `json:"productName" form:"productName"`
	ProductColor string `json:"productColor" form:"productColor"`
	ProductFlag  uint   `json:"productFlag" form:"productFlag"`
	AccountId    uint   `json:"accountId" form:"accountId"`
	request.PageInfo
}
