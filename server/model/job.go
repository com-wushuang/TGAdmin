// 自动生成模板Job
package model

import (
	"gin-vue-admin/global"
	"time"
)

// 如果含有time.Time 请自行import time包
type Job struct {
	global.GVA_MODEL
	Interval      int       `json:"interval" form:"interval" gorm:"column:interval;comment:时间间隔（s）;type:int;size:10;"`
	Active        *bool     `json:"active" form:"active" gorm:"column:active;comment:激活;type:tinyint;"`
	Advertisement string    `json:"advertisement" form:"advertisement" gorm:"column:advertisement;comment:广告内容;type:text;"`
	Expire        time.Time `json:"expire" form:"expire" gorm:"column:expire;comment:过期时间;type:time;"`
	SuccessTime   time.Time `json:"successTime" form:"successTime" gorm:"column:success_time;comment:上次成功执行时间;type:time;"`
}

func (Job) TableName() string {
	return "job"
}

// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type JobWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	Job   `json:"business"`
// }

// func (Job) TableName() string {
// 	return "job"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["job"] = func() model.GVA_Workflow {
//   return new(model.JobWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["job"] = func() interface{} {
// 	return new(model.Job)
// }
