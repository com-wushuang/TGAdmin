package v1

import (
	"gin-vue-admin/global"
    "gin-vue-admin/model"
    "gin-vue-admin/model/request"
    "gin-vue-admin/model/response"
    "gin-vue-admin/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

// @Tags Account
// @Summary 创建Account
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Account true "创建Account"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /account/createAccount [post]
func CreateAccount(c *gin.Context) {
	var account model.Account
	_ = c.ShouldBindJSON(&account)
	if err := service.CreateAccount(account); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Account
// @Summary 删除Account
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Account true "删除Account"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /account/deleteAccount [delete]
func DeleteAccount(c *gin.Context) {
	var account model.Account
	_ = c.ShouldBindJSON(&account)
	if err := service.DeleteAccount(account); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Account
// @Summary 批量删除Account
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Account"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /account/deleteAccountByIds [delete]
func DeleteAccountByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteAccountByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Account
// @Summary 更新Account
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Account true "更新Account"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /account/updateAccount [put]
func UpdateAccount(c *gin.Context) {
	var account model.Account
	_ = c.ShouldBindJSON(&account)
	if err := service.UpdateAccount(account); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Account
// @Summary 用id查询Account
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Account true "用id查询Account"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /account/findAccount [get]
func FindAccount(c *gin.Context) {
	var account model.Account
	_ = c.ShouldBindQuery(&account)
	if err, reaccount := service.GetAccount(account.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reaccount": reaccount}, c)
	}
}

// @Tags Account
// @Summary 分页获取Account列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AccountSearch true "分页获取Account列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /account/getAccountList [get]
func GetAccountList(c *gin.Context) {
	var pageInfo request.AccountSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetAccountInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
