package common

import (
	"cdnboss-middle-pro/modules/public"

	"github.com/labstack/echo"
)

// Vip vip接口
func Vip(rg *echo.Group) {

	// 同步ip到监
	rg.POST("/common/addMonitor", public.ReqRelay(host))

	// 监控反馈vip可用/不可用
	rg.POST("/common/vipMonitorStatus", public.ReqRelay(host))

	// 尝试启用vip
	rg.PUT("/common/cacheGroups/:cacheGroupId/vip/:vip/tryManualEnable", public.ReqRelay(host))
	// 尝试停用vip
	rg.PUT("/common/cacheGroups/:cacheGroupId/vip/:vip/tryManualDisable", public.ReqRelay(host))

	// 人工启用vip
	rg.PUT("/common/cacheGroups/:cacheGroupId/vip/:vip/manualEnable", public.ReqRelay(host))
	// 人工停用vip
	rg.PUT("/common/cacheGroups/:cacheGroupId/vip/:vip/manualDisable", public.ReqRelay(host))

}
