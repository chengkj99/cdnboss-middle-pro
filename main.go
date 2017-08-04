package main

import (
	"cdnboss-middle-pro/modules/common"
	"cdnboss-middle-pro/modules/dispatch"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	rg := e.Group("/api/v1/admin")

	dispatch.Line(rg)
	dispatch.View(rg)
	dispatch.Record(rg)

	common.Platform(rg)
	common.Group(rg)
	common.Vip(rg)
	common.DispatchMode(rg)

	e.Logger.Fatal(e.Start(":1323"))
}
