package main

import (
	"cdnboss-middle-pro/modules/public"
	"os"

	"github.com/labstack/gommon/log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	test, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer test.Close()
	e.Logger.SetOutput(test)
	e.Logger.SetLevel(log.DEBUG)

	// 日志输出格式
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
		Output: test,
	}))

	// 透传处理
	// e.GET("/*", passthrough)

	// e.HTTPErrorHandler = func(err error, ctx echo.Context) {
	// 	code := http.StatusInternalServerError
	// 	if he, ok := err.(*echo.HTTPError); ok {
	// 		code = he.Code
	// 	}
	// 	if code == 404 {
	// 		url := ctx.Request().RequestURI
	// 		var host string
	// 		isExist := strings.Contains(url, "/api/v1/admin")

	// 		if isExist {
	// 			host = "http://cs.zjmanageplatform.qiniu.io:8090"
	// 		}
	// 		isExist = strings.Contains(url, "/v2")
	// 		if isExist {
	// 			host = "http://124.160.26.132:51030"
	// 		}
	// 		isExist = strings.Contains(url, "/alarm")
	// 		if isExist {
	// 			host = "http://115.236.16.115:8081"
	// 		}

	// 		err := public.ReqRelay2(host, ctx)
	// 		if err != nil {
	// 			ctx.Logger().Errorf("error: %v", err)
	// 		}
	// 		fmt.Println("host:", host)
	// 		fmt.Println("enter code:", code)
	// 	}
	// }

	disR := e.Group("/api/v1/admin")
	disR.GET("/*", public.ReqRelay2("http://cs.zjmanageplatform.qiniu.io:8090", ""))
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
	monitorR := e.Group("/v2")
	monitorR.GET("/*", public.ReqRelay2("http://124.160.26.132:51030", ""))
	// monitor.Stm(monitorR)

	// // 告警系统
	alarmR := e.Group("/alarm")
	alarmR.GET("/*", public.ReqRelay2("http://115.236.16.115:8081", "/alarm"))
	// alarm.Config(alarmR)

	// tR := e.Group("/")
	// url1, err := url.Parse("http://cs.zjmanageplatform.qiniu.io:8090")
	// if err != nil {
	// 	e.Logger.Fatal(err)
	// }
	// tR.Use(middleware.Proxy(&middleware.RoundRobinBalancer{
	// 	Targets: []*middleware.ProxyTarget{
	// 		&middleware.ProxyTarget{
	// 			URL: url1,
	// 		},
	// 	},
	// }))

	e.Logger.Fatal(e.Start(":1323"))
}

// 错误处理的透传时使用
// func passthrough(ctx echo.Context) error {
// 	url := ctx.Request().RequestURI
// 	var host string
// 	isExist := strings.Contains(url, "/api/v1/admin")

// 	if isExist {
// 		host = "http://cs.zjmanageplatform.qiniu.io:8090"
// 	}
// 	isExist = strings.Contains(url, "/v2")
// 	if isExist {
// 		host = "http://124.160.26.132:51030"
// 	}
// 	isExist = strings.Contains(url, "/alarm")
// 	if isExist {
// 		host = "http://115.236.16.115:8081"
// 	}

// 	err := public.ReqRelay2(host, ctx)
// 	if err != nil {
// 		ctx.Logger().Errorf("error: %v", err)
// 	}
// 	fmt.Println("host:", host)
// 	return nil
// }
