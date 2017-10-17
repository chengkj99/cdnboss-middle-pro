package dispatch

import (
	"cdnboss-middle-pro/modules/public"

	"github.com/labstack/echo"
)

// Record 解析记录接口, doType: glb - dispatch
func Record(rg *echo.Group) {

	// 1. 通过id获取解析记录  recordId: recordId
	// 2. 获取解析记录详情 recordId: detail
	rg.GET("/dispatch/records/:recordId", public.ReqRelay(host))

	// 新建解析记录
	rg.POST("/dispatch/records", public.ReqRelay(host))

	// 获取符合条件的解析记录
	rg.POST("/dispatch/records/search", public.ReqRelay(host))

	// 修改解析记录
	rg.PUT("/dispatch/records/:recordId", public.ReqRelay(host))

	// 删除解析记录
	rg.DELETE("/dispatch/records/:recordId", public.ReqRelay(host))

}
