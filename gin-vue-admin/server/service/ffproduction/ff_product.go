package ffproduction

import (
	"gin-vue-admin/server/global"
	"gin-vue-admin/server/model/ffproduction"
	ffRes "gin-vue-admin/server/model/ffproduction/response"
)

type ProductService struct{}

var ProductServiceApp = new(ProductService)

// CreateProduct 创建产品
func (p *ProductService) CreateProduct(product ffproduction.FfProduct) error {
	return global.GVA_DB.Create(&product).Error
}

// DeleteProduct 删除产品
func (p *ProductService) DeleteProduct(product ffproduction.FfProduct) error {
	return global.GVA_DB.Delete(&product).Error
}

// DeleteProductByIds 批量删除产品
func (p *ProductService) DeleteProductByIds(ids []uint) error {
	return global.GVA_DB.Delete(&ffproduction.FfProduct{}, "id in ?", ids).Error
}

// UpdateProduct 更新产品
func (p *ProductService) UpdateProduct(product *ffproduction.FfProduct) error {
	return global.GVA_DB.Save(product).Error
}

// GetProduct 获取单个产品
func (p *ProductService) GetProduct(id uint) (ffproduction.FfProduct, error) {
	var product ffproduction.FfProduct
	err := global.GVA_DB.Where("id = ?", id).First(&product).Error
	return product, err
}

// GetProductList 获取产品列表
func (p *ProductService) GetProductList(info ffproduction.FfProductSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Table("ff_product").
		Select("ff_product.id, ff_product.product_code, ff_product.product_name, ff_product.product_color, ff_product.product_img, ff_product.account_id, ff_account.account_name, ff_product.product_order_num, ff_product.product_flag").
		Joins("LEFT JOIN ff_account ON ff_account.id = ff_product.account_id")

	if info.ProductCode != "" {
		db = db.Where("ff_product.product_code LIKE ?", "%"+info.ProductCode+"%")
	}
	if info.ProductName != "" {
		db = db.Where("ff_product.product_name LIKE ?", "%"+info.ProductName+"%")
	}
	if info.ProductColor != "" {
		db = db.Where("ff_product.product_color LIKE ?", "%"+info.ProductColor+"%")
	}
	if info.ProductFlag != 0 {
		db = db.Where("ff_product.product_flag = ?", info.ProductFlag)
	}
	if info.AccountId != 0 {
		db = db.Where("ff_product.account_id = ?", info.AccountId)
	}

	var productList []ffRes.FfProductListResponse
	err = db.Count(&total).Error
	if err != nil {
		return productList, total, err
	}
	err = db.Limit(limit).Offset(offset).Scan(&productList).Error
	return productList, total, err
}
