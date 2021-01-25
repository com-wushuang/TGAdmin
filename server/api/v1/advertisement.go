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

// @Tags Advertisement
// @Summary 创建Advertisement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Advertisement true "创建Advertisement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /advertisement/createAdvertisement [post]
func CreateAdvertisement(c *gin.Context) {
	var advertisement model.Advertisement
	_ = c.ShouldBindJSON(&advertisement)
	if err := service.CreateAdvertisement(advertisement); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Advertisement
// @Summary 删除Advertisement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Advertisement true "删除Advertisement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /advertisement/deleteAdvertisement [delete]
func DeleteAdvertisement(c *gin.Context) {
	var advertisement model.Advertisement
	_ = c.ShouldBindJSON(&advertisement)
	if err := service.DeleteAdvertisement(advertisement); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Advertisement
// @Summary 批量删除Advertisement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Advertisement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /advertisement/deleteAdvertisementByIds [delete]
func DeleteAdvertisementByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteAdvertisementByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Advertisement
// @Summary 更新Advertisement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Advertisement true "更新Advertisement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /advertisement/updateAdvertisement [put]
func UpdateAdvertisement(c *gin.Context) {
	var advertisement model.Advertisement
	_ = c.ShouldBindJSON(&advertisement)
	if err := service.UpdateAdvertisement(advertisement); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Advertisement
// @Summary 用id查询Advertisement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Advertisement true "用id查询Advertisement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /advertisement/findAdvertisement [get]
func FindAdvertisement(c *gin.Context) {
	var advertisement model.Advertisement
	_ = c.ShouldBindQuery(&advertisement)
	if err, readvertisement := service.GetAdvertisement(advertisement.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"readvertisement": readvertisement}, c)
	}
}

// @Tags Advertisement
// @Summary 分页获取Advertisement列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AdvertisementSearch true "分页获取Advertisement列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /advertisement/getAdvertisementList [get]
func GetAdvertisementList(c *gin.Context) {
	var pageInfo request.AdvertisementSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetAdvertisementInfoList(pageInfo); err != nil {
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
