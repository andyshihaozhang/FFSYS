package ffproduction

import (
	"gin-vue-admin/server/global"
)

// FfAccount 客户管理结构体
type FfAccount struct {
	global.GVA_MODEL
	AccountName string `json:"accountName" form:"accountName" gorm:"column:account_name;comment:客户名"`
	Phone       string `json:"phone" form:"phone" gorm:"column:phone;comment:联系方式"`
	Address     string `json:"address" form:"address" gorm:"column:address;comment:联系地址"`
}

// TableName 指定表名
func (FfAccount) TableName() string {
	return "ff_account"
}
