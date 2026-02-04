package ffproduction

import (
	"gin-vue-admin/server/global"
	"gin-vue-admin/server/model/ffproduction"
)

type MaterialService struct{}

var MaterialServiceApp = new(MaterialService)

// CreateMaterial 创建物料
func (m *MaterialService) CreateMaterial(material ffproduction.FfMaterial) error {
	return global.GVA_DB.Create(&material).Error
}

// DeleteMaterial 删除物料
func (m *MaterialService) DeleteMaterial(material ffproduction.FfMaterial) error {
	return global.GVA_DB.Delete(&material).Error
}

// UpdateMaterial 更新物料
func (m *MaterialService) UpdateMaterial(material *ffproduction.FfMaterial) error {
	return global.GVA_DB.Save(material).Error
}

// GetMaterialList 获取物料列表
func (m *MaterialService) GetMaterialList(info ffproduction.FfMaterialSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&ffproduction.FfMaterial{})

	if info.ProductId != 0 {
		db = db.Where("product_id = ?", info.ProductId)
	}
	if info.MaterialName != "" {
		db = db.Where("material_name LIKE ?", "%"+info.MaterialName+"%")
	}

	var materialList []ffproduction.FfMaterial
	err = db.Count(&total).Error
	if err != nil {
		return materialList, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&materialList).Error
	return materialList, total, err
}
