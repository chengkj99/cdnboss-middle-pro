package dispatch

import (
	"cdnboss-middle-pro/modules/public"

	"github.com/labstack/echo"
)

// View 获取所有view,   doType: glb
func View(rg *echo.Group) {

	rg.GET("/:doType/views", public.ReqRelay(host))
}
