package monitor

import (
	"cdnboss-middle-pro/modules/public"

	"github.com/labstack/echo"
)

// Stm 可用性趋势图
func Stm(rg *echo.Group) {

	// 获取基础线路/融合线路/中间源当前可用性 type=""&start_time=""&end_time=""
	rg.GET("/fstats/line/info", public.ReqRelay(host))

	// 趋势图（获取线路一段时间内可用性） type=""&lineid=""&start_time=""&end_time=""
	rg.GET("/fstats/line/tend", public.ReqRelay(host))

	// 地区运营商的平均可用性 type=""&lineid=""&start_time=""&end_time=""
	rg.GET("/fstats/line/viewsavg", public.ReqRelay(host))

	// 融合线路监控列表
	rg.GET("/fusionlines/monlist", public.ReqRelay(host))

	// 设置线路为监控状态
	rg.PUT("/fusionlines/:lineid/stats/enable", public.ReqRelay(host))

	// 设置线路为非监控状态
	rg.PUT("/fusionlines/:lineid/stats/disable", public.ReqRelay(host))
}
