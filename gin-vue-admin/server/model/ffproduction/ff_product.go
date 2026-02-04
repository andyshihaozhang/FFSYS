package ffproduction

import (
	"gin-vue-admin/server/global"
)

// FfProduct 产品管理结构体
type FfProduct struct {
	global.GVA_MODEL
	ProductCode     string `json:"productCode" form:"productCode" gorm:"column:product_code;comment:产品款号"`
	ProductName     string `json:"productName" form:"productName" gorm:"column:product_name;comment:产品名称"`
	ProductColor    string `json:"productColor" form:"productColor" gorm:"column:product_color;comment:产品颜色"`
	ProductImg      string `json:"productImg" form:"productImg" gorm:"column:product_img;comment:产品图片"`
	AccountId       uint   `json:"accountId" form:"accountId" gorm:"column:account_id;comment:客户ID"`
	ProductOrderNum int    `json:"productOrderNum" form:"productOrderNum" gorm:"column:product_order_num;comment:下单数量"`
	ProductFlag     uint   `json:"productFlag" form:"productFlag" gorm:"column:product_flag;comment:生产标志(1未生产2生产中3已完成)"`
}

// TableName 指定表名
func (FfProduct) TableName() string {
	return "ff_product"
}

// FfProductSearch 产品搜索结构体
type FfProductSearch struct {
	ProductCode  string `json:"productCode" form:"productCode"`
	ProductName  string `json:"productName" form:"productName"`
	ProductColor string `json:"productColor" form:"productColor"`
	ProductFlag  uint   `json:"productFlag" form:"productFlag"`
	AccountId    uint   `json:"accountId" form:"accountId"`
	PageInfo
}

// PageInfo 分页信息
type PageInfo struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}
