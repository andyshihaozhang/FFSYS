package ffproduction

import (
	"gin-vue-admin/server/global"
	"gin-vue-admin/server/model/ffproduction"
)

type EmployeeService struct{}

var EmployeeServiceApp = new(EmployeeService)

// CreateEmployee 创建员工
func (e *EmployeeService) CreateEmployee(employee ffproduction.FfEmployee) error {
	return global.GVA_DB.Create(&employee).Error
}

// DeleteEmployee 删除员工
func (e *EmployeeService) DeleteEmployee(employee ffproduction.FfEmployee) error {
	return global.GVA_DB.Delete(&employee).Error
}

// UpdateEmployee 更新员工
func (e *EmployeeService) UpdateEmployee(employee *ffproduction.FfEmployee) error {
	return global.GVA_DB.Save(employee).Error
}

// GetEmployeeList 获取员工列表
func (e *EmployeeService) GetEmployeeList(info ffproduction.FfEmployeeSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&ffproduction.FfEmployee{})

	if info.EmployeeName != "" {
		db = db.Where("employee_name LIKE ?", "%"+info.EmployeeName+"%")
	}
	if info.Phone != "" {
		db = db.Where("phone LIKE ?", "%"+info.Phone+"%")
	}
	if info.Sex != 0 {
		db = db.Where("sex = ?", info.Sex)
	}
	if info.AuthorityIds != 0 {
		db = db.Where("authority_ids = ?", info.AuthorityIds)
	}
	if info.EmployeeStatus != 0 {
		db = db.Where("employee_status = ?", info.EmployeeStatus)
	}

	var employeeList []ffproduction.FfEmployee
	err = db.Count(&total).Error
	if err != nil {
		return employeeList, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&employeeList).Error
	return employeeList, total, err
}
