package main

import (
	"cdnboss-middle-pro/modules/common"
	"cdnboss-middle-pro/modules/dispatch"
	"cdnboss-middle-pro/modules/monitor"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	disR := e.Group("/api/v1/admin")

	dispatch.Line(disR)
	dispatch.View(disR)
	dispatch.Record(disR)

	common.Platform(disR)
	common.Group(disR)
	common.Vip(disR)
	common.DispatchMode(disR)

	monitorR := e.Group("/v2")

	monitor.Stm(monitorR)

	e.Logger.Fatal(e.Start(":1323"))
}
