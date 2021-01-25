// 自动生成模板Advertisement
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Advertisement struct {
      global.GVA_MODEL
      Text  string `json:"text" form:"text" gorm:"column:text;comment:广告内容;type:text(1024);size:1024;"`
      Key  string `json:"key" form:"key" gorm:"column:key;comment:关键词;type:varchar(255);size:255;"`
}


func (Advertisement) TableName() string {
  return "advertisement"
}


// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type AdvertisementWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	Advertisement   `json:"business"`
// }

// func (Advertisement) TableName() string {
// 	return "advertisement"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["advertisement"] = func() model.GVA_Workflow {
//   return new(model.AdvertisementWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["advertisement"] = func() interface{} {
// 	return new(model.Advertisement)
// }
