package ffproduction

import (
	"gin-vue-admin/server/global"
	"gin-vue-admin/server/model/ffproduction"
)

type OperationService struct{}

var OperationServiceApp = new(OperationService)

// CreateOperation 创建工序
func (o *OperationService) CreateOperation(operation ffproduction.FfOperation) error {
	return global.GVA_DB.Create(&operation).Error
}

// DeleteOperation 删除工序
func (o *OperationService) DeleteOperation(operation ffproduction.FfOperation) error {
	return global.GVA_DB.Delete(&operation).Error
}

// UpdateOperation 更新工序
func (o *OperationService) UpdateOperation(operation *ffproduction.FfOperation) error {
	return global.GVA_DB.Save(operation).Error
}

// GetOperationList 获取工序列表
func (o *OperationService) GetOperationList(info ffproduction.FfOperationSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&ffproduction.FfOperation{})

	if info.ProductId != 0 {
		db = db.Where("product_id = ?", info.ProductId)
	}
	if info.OperationName != "" {
		db = db.Where("operation_name LIKE ?", "%"+info.OperationName+"%")
	}

	var operationList []ffproduction.FfOperation
	err = db.Count(&total).Error
	if err != nil {
		return operationList, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&operationList).Error
	return operationList, total, err
}
