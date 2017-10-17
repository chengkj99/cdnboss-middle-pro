package common

import (
	"cdnboss-middle-pro/modules/public"

	"github.com/labstack/echo"
)

//Group 缓存组接口
func Group(rg *echo.Group) {

	// 获取所有缓存组
	rg.GET("/common/cacheGroups", public.ReqRelay(host))

	// 获取缓存组-根据缓存组ID
	rg.GET("/common/cacheGroups/:cacheGroupId", public.ReqRelay(host))

	// 获取缓存组服务信息 ../serviceStatus?vip=183.2.208.21
	rg.GET("/common/cacheGroups/serviceStatus", public.ReqRelay(host))

	// 获取缓存组涉及的解析
	rg.GET("/common/cacheGroups/:cacheGroupId/records", public.ReqRelay(host))

	// 搜索缓存组
	rg.POST("/common/cacheGroups/search", public.ReqRelay(host))

	// 新建缓存组
	rg.POST("/common/cacheGroups", public.ReqRelay(host))

	// 缓存组尝试增加vip
	rg.POST("/common/cacheGroups/:cacheGroupId/tryAddVip", public.ReqRelay(host))
	// 缓存组增加vip
	rg.POST("/common/cacheGroups/:cacheGroupId/addVip", public.ReqRelay(host))

	// 修改缓存组基本信息
	rg.PUT("/common/cacheGroups/:cacheGroupId", public.ReqRelay(host))

	// 修改缓存组服务信息
	rg.PUT("/common/cacheGroups/serviceStatus", public.ReqRelay(host))

	// 缓存组尝试切换到自动模式
	rg.PUT("/common/cacheGroups/:cacheGroupId/tryAuto", public.ReqRelay(host))
	// 缓存组切换到自动模式
	rg.PUT("/common/cacheGroups/:cacheGroupId/auto", public.ReqRelay(host))

	// 尝试人工上线缓存组
	rg.PUT("/common/cacheGroups/:cacheGroupId/tryManualOnline", public.ReqRelay(host))
	// 人工上线缓存组
	rg.PUT("/common/cacheGroups/:cacheGroupId/manualOnline", public.ReqRelay(host))

	// 尝试人工下线缓存组
	rg.PUT("/common/cacheGroups/:cacheGroupId/tryManualOffline", public.ReqRelay(host))
	// 人工下线缓存组
	rg.PUT("/common/cacheGroups/:cacheGroupId/manualOffline", public.ReqRelay(host))

	// 删除缓存组
	rg.DELETE("/common/cacheGroups/:cacheGroupId", public.ReqRelay(host))

	// 缓存组尝试删除vip
	rg.DELETE("/common/cacheGroups/:cacheGroupId/tryDelVip/:ip", public.ReqRelay(host))

	// 缓存组删除vip
	rg.DELETE("/common/cacheGroups/:cacheGroupId/delVip/:ip", public.ReqRelay(host))

}
