package request

import (
	"gin-vue-admin/server/model/common/request"
)

type FfOperationSearch struct {
	ProductId     uint   `json:"productId" form:"productId"`
	OperationName string `json:"operationName" form:"operationName"`
	request.PageInfo
}
