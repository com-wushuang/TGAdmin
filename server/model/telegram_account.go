// 自动生成模板Account
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Account struct {
      global.GVA_MODEL
      ApiId  int `json:"apiId" form:"apiId" gorm:"column:api_id;comment:appId;"`
      ApiHash  string `json:"apiHash" form:"apiHash" gorm:"column:api_hash;comment:apiHash;"`
}


func (Account) TableName() string {
  return "account"
}


// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type AccountWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	Account   `json:"business"`
// }

// func (Account) TableName() string {
// 	return "account"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["account"] = func() model.GVA_Workflow {
//   return new(model.AccountWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["account"] = func() interface{} {
// 	return new(model.Account)
// }
