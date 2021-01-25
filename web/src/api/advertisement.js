import service from '@/utils/request'

// @Tags Advertisement
// @Summary 创建Advertisement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Advertisement true "创建Advertisement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /advertisement/createAdvertisement [post]
export const createAdvertisement = (data) => {
     return service({
         url: "/advertisement/createAdvertisement",
         method: 'post',
         data
     })
 }


// @Tags Advertisement
// @Summary 删除Advertisement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Advertisement true "删除Advertisement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /advertisement/deleteAdvertisement [delete]
 export const deleteAdvertisement = (data) => {
     return service({
         url: "/advertisement/deleteAdvertisement",
         method: 'delete',
         data
     })
 }

// @Tags Advertisement
// @Summary 删除Advertisement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Advertisement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /advertisement/deleteAdvertisement [delete]
 export const deleteAdvertisementByIds = (data) => {
     return service({
         url: "/advertisement/deleteAdvertisementByIds",
         method: 'delete',
         data
     })
 }

// @Tags Advertisement
// @Summary 更新Advertisement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Advertisement true "更新Advertisement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /advertisement/updateAdvertisement [put]
 export const updateAdvertisement = (data) => {
     return service({
         url: "/advertisement/updateAdvertisement",
         method: 'put',
         data
     })
 }


// @Tags Advertisement
// @Summary 用id查询Advertisement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Advertisement true "用id查询Advertisement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /advertisement/findAdvertisement [get]
 export const findAdvertisement = (params) => {
     return service({
         url: "/advertisement/findAdvertisement",
         method: 'get',
         params
     })
 }


// @Tags Advertisement
// @Summary 分页获取Advertisement列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Advertisement列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /advertisement/getAdvertisementList [get]
 export const getAdvertisementList = (params) => {
     return service({
         url: "/advertisement/getAdvertisementList",
         method: 'get',
         params
     })
 }