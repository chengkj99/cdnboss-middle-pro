package dispatch

import (
	"cdnboss-middle-pro/modules/public"

	"github.com/labstack/echo"
)

// Line 线路接口, doType: glb - dispatch
func Line(rg *echo.Group) {

	// 获取全部线路
	rg.GET("/dispatch/lines", public.ReqRelay(host))

	// 获取线路
	rg.GET("/dispatch/lines/:lineId", public.ReqRelay(host))

	// 新建线路
	rg.POST("/dispatch/lines", public.ReqRelay(host))

	// 修改线路
	rg.PUT("/dispatch/lines/:lineId", public.ReqRelay(host))

}
