package main

import (
	"cdnboss-middle-pro/modules/public"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/gommon/log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	accCdnboss, err := os.OpenFile("acc_cdnboss.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer accCdnboss.Close()
	e.Logger.SetOutput(accCdnboss)
	e.Logger.SetLevel(log.DEBUG)

	// 日志输出格式
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
		Output: accCdnboss,
	}))

	e.HTTPErrorHandler = func(err error, ctx echo.Context) {
		code := http.StatusInternalServerError
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
			ctx.JSON(code, he)
			fmt.Println("code, he.Message..", code, he.Message)
		}
	}

	// disR := e.Group("/api/v1/admin")
	// disR.GET("/*", public.ReqRelay("http://cs.zjmanageplatform.qiniu.io:8090"))
	// // 调度系统
	// dispatch.Line(disR)
	// dispatch.View(disR)
	// dispatch.Record(disR)

	// // 公共接口
	// common.Platform(disR)
	// common.Group(disR)
	// common.Vip(disR)
	// common.DispatchMode(disR)

	// 监控
	// test url localhost:1323/fcm-api/v2/node/alarm/rules
	// monitorR := e.Group("/fcm-api")
	// monitorR.GET("/*", public.ReqRelay("http://124.160.26.132:51200", "/fcm-api"))
	// monitor.Stm(monitorR)

	// // 告警系统
	alarmR := e.Group("/alarm")
	alarmR.GET("/*", public.ReqRelay("http://115.236.16.115:8081", "/alarm"))

	e.Logger.Fatal(e.Start(":1323"))
}
