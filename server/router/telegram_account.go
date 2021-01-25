package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitAccountRouter(Router *gin.RouterGroup) {
	AccountRouter := Router.Group("account").Use(middleware.OperationRecord())
	{
		AccountRouter.POST("createAccount", v1.CreateAccount)   // 新建Account
		AccountRouter.DELETE("deleteAccount", v1.DeleteAccount) // 删除Account
		AccountRouter.DELETE("deleteAccountByIds", v1.DeleteAccountByIds) // 批量删除Account
		AccountRouter.PUT("updateAccount", v1.UpdateAccount)    // 更新Account
		AccountRouter.GET("findAccount", v1.FindAccount)        // 根据ID获取Account
		AccountRouter.GET("getAccountList", v1.GetAccountList)  // 获取Account列表
	}
}
