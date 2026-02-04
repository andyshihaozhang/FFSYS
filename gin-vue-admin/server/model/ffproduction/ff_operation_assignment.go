package ffproduction

import (
	"gin-vue-admin/server/global"
)

// FfOperationAssignment 工序分配结构体
type FfOperationAssignment struct {
	global.GVA_MODEL
	ProductId          uint    `json:"productId" form:"productId" gorm:"column:product_id;comment:产品ID"`
	OperationId        uint    `json:"operationId" form:"operationId" gorm:"column:operation_id;comment:工序ID"`
	EmployeeId         uint    `json:"employeeId" form:"employeeId" gorm:"column:employee_id;comment:员工ID"`
	UnitPrice          float64 `json:"unitPrice" form:"unitPrice" gorm:"column:unit_price;comment:单价"`
	AssignmentQuantity int     `json:"assignmentQuantity" form:"assignmentQuantity" gorm:"column:assignment_quantity;comment:分配数量"`
}

// TableName 指定表名
func (FfOperationAssignment) TableName() string {
	return "ff_operation_assignment"
}

// FfOperationAssignmentSearch 工序分配搜索结构体
type FfOperationAssignmentSearch struct {
	ProductId   uint `json:"productId" form:"productId"`
	OperationId uint `json:"operationId" form:"operationId"`
	PageInfo
}
