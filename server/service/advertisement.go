package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateAdvertisement
//@description: 创建Advertisement记录
//@param: advertisement model.Advertisement
//@return: err error

func CreateAdvertisement(advertisement model.Advertisement) (err error) {
	err = global.GVA_DB.Create(&advertisement).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAdvertisement
//@description: 删除Advertisement记录
//@param: advertisement model.Advertisement
//@return: err error

func DeleteAdvertisement(advertisement model.Advertisement) (err error) {
	err = global.GVA_DB.Delete(&advertisement).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAdvertisementByIds
//@description: 批量删除Advertisement记录
//@param: ids request.IdsReq
//@return: err error

func DeleteAdvertisementByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Advertisement{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateAdvertisement
//@description: 更新Advertisement记录
//@param: advertisement *model.Advertisement
//@return: err error

func UpdateAdvertisement(advertisement model.Advertisement) (err error) {
	err = global.GVA_DB.Save(&advertisement).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAdvertisement
//@description: 根据id获取Advertisement记录
//@param: id uint
//@return: err error, advertisement model.Advertisement

func GetAdvertisement(id uint) (err error, advertisement model.Advertisement) {
	err = global.GVA_DB.Where("id = ?", id).First(&advertisement).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAdvertisementInfoList
//@description: 分页获取Advertisement记录
//@param: info request.AdvertisementSearch
//@return: err error, list interface{}, total int64

func GetAdvertisementInfoList(info request.AdvertisementSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Advertisement{})
    var advertisements []model.Advertisement
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&advertisements).Error
	return err, advertisements, total
}