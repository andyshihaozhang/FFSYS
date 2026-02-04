package ffproduction

import (
	"gin-vue-admin/server/global"
)

// FfOperation 工序管理结构体
type FfOperation struct {
	global.GVA_MODEL
	ProductId     uint   `json:"productId" form:"productId" gorm:"column:product_id;comment:产品ID"`
	OperationName string `json:"operationName" form:"operationName" gorm:"column:operation_name;comment:工序名称"`
	OperationCode string `json:"operationCode" form:"operationCode" gorm:"column:operation_code;comment:工序编码"`
	Sequence      int    `json:"sequence" form:"sequence" gorm:"column:sequence;comment:工序顺序"`
	Duration      int    `json:"duration" form:"duration" gorm:"column:duration;comment:工序耗时(分钟)"`
	Remark        string `json:"remark" form:"remark" gorm:"column:remark;comment:备注"`
}

// TableName 指定表名
func (FfOperation) TableName() string {
	return "ff_operation"
}

// FfOperationSearch 工序搜索结构体
type FfOperationSearch struct {
	ProductId     uint   `json:"productId" form:"productId"`
	OperationName string `json:"operationName" form:"operationName"`
	PageInfo
}
