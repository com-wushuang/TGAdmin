package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateJob
//@description: 创建Job记录
//@param: job model.Job
//@return: err error

func CreateJob(job model.Job) (err error) {
	err = global.GVA_DB.Create(&job).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteJob
//@description: 删除Job记录
//@param: job model.Job
//@return: err error

func DeleteJob(job model.Job) (err error) {
	err = global.GVA_DB.Delete(&job).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteJobByIds
//@description: 批量删除Job记录
//@param: ids request.IdsReq
//@return: err error

func DeleteJobByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Job{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateJob
//@description: 更新Job记录
//@param: job *model.Job
//@return: err error

func UpdateJob(job model.Job) (err error) {
	err = global.GVA_DB.Save(&job).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetJob
//@description: 根据id获取Job记录
//@param: id uint
//@return: err error, job model.Job

func GetJob(id uint) (err error, job model.Job) {
	err = global.GVA_DB.Where("id = ?", id).First(&job).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetJobInfoList
//@description: 分页获取Job记录
//@param: info request.JobSearch
//@return: err error, list interface{}, total int64

func GetJobInfoList(info request.JobSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Job{})
	var jobs []model.Job
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&jobs).Error
	return err, jobs, total
}
