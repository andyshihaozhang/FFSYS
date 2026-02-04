package ffproduction

import (
	"gin-vue-admin/server/global"
	"gin-vue-admin/server/model/ffproduction"
	"gin-vue-admin/server/model/ffproduction/response"
)

type OperationAssignmentService struct{}

var OperationAssignmentServiceApp = new(OperationAssignmentService)

// CreateOperationAssignment 创建工序分配
func (o *OperationAssignmentService) CreateOperationAssignment(assignment ffproduction.FfOperationAssignment) error {
	return global.GVA_DB.Create(&assignment).Error
}

// DeleteOperationAssignment 删除工序分配
func (o *OperationAssignmentService) DeleteOperationAssignment(assignment ffproduction.FfOperationAssignment) error {
	return global.GVA_DB.Delete(&assignment).Error
}

// UpdateOperationAssignment 更新工序分配
func (o *OperationAssignmentService) UpdateOperationAssignment(assignment *ffproduction.FfOperationAssignment) error {
	return global.GVA_DB.Save(assignment).Error
}

// GetOperationAssignmentList 获取工序分配列表
func (o *OperationAssignmentService) GetOperationAssignmentList(info ffproduction.FfOperationAssignmentSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&ffproduction.FfOperationAssignment{})

	if info.ProductId != 0 {
		db = db.Where("product_id = ?", info.ProductId)
	}
	if info.OperationId != 0 {
		db = db.Where("operation_id = ?", info.OperationId)
	}

	var assignmentList []ffproduction.FfOperationAssignment
	err = db.Count(&total).Error
	if err != nil {
		return assignmentList, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&assignmentList).Error
	if err != nil {
		return assignmentList, total, err
	}

	// 包装响应数据，关联员工名称
	var responseList []response.FfOperationAssignmentResponse
	for _, assignment := range assignmentList {
		var employee ffproduction.FfEmployee
		global.GVA_DB.Model(&ffproduction.FfEmployee{}).Where("id = ?", assignment.EmployeeId).First(&employee)

		responseList = append(responseList, response.FfOperationAssignmentResponse{
			ID:                 assignment.ID,
			ProductId:          assignment.ProductId,
			OperationId:        assignment.OperationId,
			EmployeeId:         assignment.EmployeeId,
			EmployeeName:       employee.EmployeeName,
			UnitPrice:          assignment.UnitPrice,
			AssignmentQuantity: assignment.AssignmentQuantity,
			CreatedAt:          assignment.CreatedAt.Unix(),
			UpdatedAt:          assignment.UpdatedAt.Unix(),
		})
	}

	return responseList, total, err
}
