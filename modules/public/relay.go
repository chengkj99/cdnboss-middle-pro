package public

import (
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/labstack/echo"
)

// ReqRelay 对前端请求不做处理直接转发
func ReqRelay(host string) func(c echo.Context) error {
	return func(c echo.Context) error {
		url := host + c.Request().RequestURI
		req, err := http.NewRequest(c.Request().Method, url, c.Request().Body)
		if err != nil {
			c.Logger().Errorf("error: %v", err)
			return nil
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			c.Logger().Errorf("error: %v", err)
			return nil
		}
		defer res.Body.Close()
		c.Response().Header().Set("Content-Type", "application/json; charset=utf-8")
		c.Response().Status = res.StatusCode
		_, err = io.Copy(c.Response().Writer, res.Body)
		if err != nil {
			c.Logger().Errorf("error: %v", err)
			return nil
		}
		return nil
	}
}

// ReqRelay2 对前端请求不做处理直接转发
func ReqRelay2(host string, rewriteURL string) func(c echo.Context) error {
	return func(c echo.Context) error {
		url := host + c.Request().RequestURI
		fmt.Println("url.....", url)
		if rewriteURL != "" {
			r, _ := regexp.Compile(rewriteURL)
			url = r.ReplaceAllString(url, "")
		}
		fmt.Println("url.....", url)
		req, err := http.NewRequest(c.Request().Method, url, c.Request().Body)
		if err != nil {
			c.Logger().Errorf("error: %v", err)
			return nil
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			c.Logger().Errorf("error: %v", err)
			return nil
		}
		defer res.Body.Close()
		c.Response().Header().Set("Content-Type", "application/json; charset=utf-8")
		c.Response().Status = res.StatusCode
		_, err = io.Copy(c.Response().Writer, res.Body)
		if err != nil {
			c.Logger().Errorf("error: %v", err)
			return nil
		}
		return nil
	}
}

// ReqRelay3 透传处理
// func ReqRelay3(host string, c echo.Context) error {
// 	url := host + c.Request().RequestURI
// 	req, err := http.NewRequest(c.Request().Method, url, c.Request().Body)
// 	if err != nil {
// 		c.Logger().Errorf("error: %v", err)
// 		return nil
// 	}

// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		c.Logger().Errorf("error: %v", err)
// 		return nil
// 	}
// 	defer res.Body.Close()
// 	c.Response().Header().Set("Content-Type", "application/json; charset=utf-8")
// 	c.Response().Status = res.StatusCode
// 	_, err = io.Copy(c.Response().Writer, res.Body)
// 	if err != nil {
// 		c.Logger().Errorf("error: %v", err)
// 		return nil
// 	}
// 	return nil
// }
