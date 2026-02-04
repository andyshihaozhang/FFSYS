package request

import (
	"gin-vue-admin/server/model/common/request"
)

type FfMaterialSearch struct {
	ProductId    uint   `json:"productId" form:"productId"`
	MaterialName string `json:"materialName" form:"materialName"`
	request.PageInfo
}
