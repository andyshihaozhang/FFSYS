package ffproduction

import (
	"gin-vue-admin/server/global"
)

// FfEmployee 员工管理结构体
type FfEmployee struct {
	global.GVA_MODEL
	EmployeeName   string `json:"employeeName" form:"employeeName" gorm:"column:employee_name;comment:员工名"`
	Phone          string `json:"phone" form:"phone" gorm:"column:phone;comment:手机号"`
	Age            int    `json:"age" form:"age" gorm:"column:age;comment:年龄"`
	Sex            int    `json:"sex" form:"sex" gorm:"column:sex;comment:性别(1男2女)"`
	AuthorityIds   int    `json:"authorityIds" form:"authorityIds" gorm:"column:authority_ids;comment:角色(1组长2员工)"`
	EmployeeStatus int    `json:"employeeStatus" form:"employeeStatus" gorm:"column:employee_status;comment:状态(1在职2离职3休假)"`
}

// TableName 指定表名
func (FfEmployee) TableName() string {
	return "ff_employee"
}
