package public

import (
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

// ReqRelay 对前端请求不做处理直接转发的方法
func ReqRelay(host string) func(c echo.Context) error {
	return func(c echo.Context) error {
		url := host + c.Request().RequestURI
		fmt.Println("url:", url)
		req, err := http.NewRequest(c.Request().Method, url, c.Request().Body)
		if err != nil {
			fmt.Println(err)
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err)
		}
		defer res.Body.Close()
		c.Response().Header().Set("Content-Type", "application/json; charset=utf-8")
		_, err = io.Copy(c.Response().Writer, res.Body)
		if err != nil {
			fmt.Println(err)
		}
		return err
	}
}
