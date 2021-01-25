package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateAccount
//@description: 创建Account记录
//@param: account model.Account
//@return: err error

func CreateAccount(account model.Account) (err error) {
	err = global.GVA_DB.Create(&account).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAccount
//@description: 删除Account记录
//@param: account model.Account
//@return: err error

func DeleteAccount(account model.Account) (err error) {
	err = global.GVA_DB.Delete(&account).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAccountByIds
//@description: 批量删除Account记录
//@param: ids request.IdsReq
//@return: err error

func DeleteAccountByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Account{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateAccount
//@description: 更新Account记录
//@param: account *model.Account
//@return: err error

func UpdateAccount(account model.Account) (err error) {
	err = global.GVA_DB.Save(&account).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAccount
//@description: 根据id获取Account记录
//@param: id uint
//@return: err error, account model.Account

func GetAccount(id uint) (err error, account model.Account) {
	err = global.GVA_DB.Where("id = ?", id).First(&account).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAccountInfoList
//@description: 分页获取Account记录
//@param: info request.AccountSearch
//@return: err error, list interface{}, total int64

func GetAccountInfoList(info request.AccountSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Account{})
    var accounts []model.Account
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&accounts).Error
	return err, accounts, total
}