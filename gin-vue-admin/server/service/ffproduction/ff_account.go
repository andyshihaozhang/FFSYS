package ffproduction

import (
	"gin-vue-admin/server/global"
	"gin-vue-admin/server/model/ffproduction"
)

type AccountService struct{}

var AccountServiceApp = new(AccountService)

// CreateAccount 创建客户
func (a *AccountService) CreateAccount(account ffproduction.FfAccount) error {
	return global.GVA_DB.Create(&account).Error
}

// DeleteAccount 删除客户
func (a *AccountService) DeleteAccount(account ffproduction.FfAccount) error {
	return global.GVA_DB.Delete(&account).Error
}

// UpdateAccount 更新客户
func (a *AccountService) UpdateAccount(account *ffproduction.FfAccount) error {
	return global.GVA_DB.Save(account).Error
}

// GetAccount 获取单个客户
func (a *AccountService) GetAccount(id uint) (ffproduction.FfAccount, error) {
	var account ffproduction.FfAccount
	err := global.GVA_DB.Where("id = ?", id).First(&account).Error
	return account, err
}

// GetAccountList 获取客户列表
func (a *AccountService) GetAccountList(info ffproduction.FfAccountSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&ffproduction.FfAccount{})

	if info.AccountName != "" {
		db = db.Where("account_name LIKE ?", "%"+info.AccountName+"%")
	}
	if info.Phone != "" {
		db = db.Where("phone LIKE ?", "%"+info.Phone+"%")
	}
	if info.Address != "" {
		db = db.Where("address LIKE ?", "%"+info.Address+"%")
	}

	var accountList []ffproduction.FfAccount
	err = db.Count(&total).Error
	if err != nil {
		return accountList, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&accountList).Error
	return accountList, total, err
}

// GetAllAccounts 获取所有客户
func (a *AccountService) GetAllAccounts() ([]ffproduction.FfAccount, error) {
	var accountList []ffproduction.FfAccount
	err := global.GVA_DB.Find(&accountList).Error
	return accountList, err
}
