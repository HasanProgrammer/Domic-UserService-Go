package Routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func UserRoutesRegister(e *echo.Echo) {

	e.GET("/", func(c echo.Context) error {

		return c.JSON(http.StatusBadRequest, "Hello, World!")

	})

}
