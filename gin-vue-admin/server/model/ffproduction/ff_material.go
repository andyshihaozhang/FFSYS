package ffproduction

import (
	"gin-vue-admin/server/global"
)

// FfMaterial 物料管理结构体
type FfMaterial struct {
	global.GVA_MODEL
	ProductId    uint   `json:"productId" form:"productId" gorm:"column:product_id;comment:产品ID"`
	MaterialName string `json:"materialName" form:"materialName" gorm:"column:material_name;comment:物料名称"`
	MaterialCode string `json:"materialCode" form:"materialCode" gorm:"column:material_code;comment:物料编码"`
	Quantity     int    `json:"quantity" form:"quantity" gorm:"column:quantity;comment:数量"`
	Unit         string `json:"unit" form:"unit" gorm:"column:unit;comment:单位"`
	Remark       string `json:"remark" form:"remark" gorm:"column:remark;comment:备注"`
}

// TableName 指定表名
func (FfMaterial) TableName() string {
	return "ff_material"
}

// FfMaterialSearch 物料搜索结构体
type FfMaterialSearch struct {
	ProductId    uint   `json:"productId" form:"productId"`
	MaterialName string `json:"materialName" form:"materialName"`
	PageInfo
}
