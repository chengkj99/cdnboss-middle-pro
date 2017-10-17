package alarm

import (
	"github.com/labstack/echo"
)

func User(rg *echo.Group) {

	rg.GET("/user", getUser)

}

func getUser(c echo.Context) error {

	return nil
}
