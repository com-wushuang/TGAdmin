package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitAdvertisementRouter(Router *gin.RouterGroup) {
	AdvertisementRouter := Router.Group("advertisement").Use(middleware.OperationRecord())
	{
		AdvertisementRouter.POST("createAdvertisement", v1.CreateAdvertisement)   // 新建Advertisement
		AdvertisementRouter.DELETE("deleteAdvertisement", v1.DeleteAdvertisement) // 删除Advertisement
		AdvertisementRouter.DELETE("deleteAdvertisementByIds", v1.DeleteAdvertisementByIds) // 批量删除Advertisement
		AdvertisementRouter.PUT("updateAdvertisement", v1.UpdateAdvertisement)    // 更新Advertisement
		AdvertisementRouter.GET("findAdvertisement", v1.FindAdvertisement)        // 根据ID获取Advertisement
		AdvertisementRouter.GET("getAdvertisementList", v1.GetAdvertisementList)  // 获取Advertisement列表
	}
}
