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

// @Tags Job
// @Summary 创建Job
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Job true "创建Job"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /job/createJob [post]
func CreateJob(c *gin.Context) {
	var job model.Job
	_ = c.ShouldBindJSON(&job)
	if err := service.CreateJob(job); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Job
// @Summary 删除Job
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Job true "删除Job"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /job/deleteJob [delete]
func DeleteJob(c *gin.Context) {
	var job model.Job
	_ = c.ShouldBindJSON(&job)
	if err := service.DeleteJob(job); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Job
// @Summary 批量删除Job
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Job"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /job/deleteJobByIds [delete]
func DeleteJobByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteJobByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Job
// @Summary 更新Job
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Job true "更新Job"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /job/updateJob [put]
func UpdateJob(c *gin.Context) {
	var job model.Job
	_ = c.ShouldBindJSON(&job)
	if err := service.UpdateJob(job); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Job
// @Summary 用id查询Job
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Job true "用id查询Job"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /job/findJob [get]
func FindJob(c *gin.Context) {
	var job model.Job
	_ = c.ShouldBindQuery(&job)
	if err, rejob := service.GetJob(job.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rejob": rejob}, c)
	}
}

// @Tags Job
// @Summary 分页获取Job列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.JobSearch true "分页获取Job列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /job/getJobList [get]
func GetJobList(c *gin.Context) {
	var pageInfo request.JobSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetJobInfoList(pageInfo); err != nil {
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
