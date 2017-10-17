package common

import (
	"cdnboss-middle-pro/modules/public"

	"github.com/labstack/echo"
)

// Platform 平台信息接口
func Platform(rg *echo.Group) {

	// 获取信息平台
	rg.GET("/common/platforms/:platformId", public.ReqRelay(host))

	// 获取信息平台列表
	rg.GET("/common/platforms", public.ReqRelay(host))

	// 获取信息平台下缓存组列表
	rg.GET("/common/platforms/:platformId/cacheGroups", public.ReqRelay(host))

	// 新建信息平台
	rg.POST("/common/platforms", public.ReqRelay(host))

	// 修改信息平台
	rg.PUT("/common/platforms/:platformId", public.ReqRelay(host))

	// 获取指定平台下的所有线路
	rg.GET("/common/platforms/:platform_id/lines", public.ReqRelay(host))

}
