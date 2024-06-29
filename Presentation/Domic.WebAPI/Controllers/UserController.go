package Controllers

import (
	"domic.usecase/UserUseCase/Commands/Create"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RegisterUserControllerActions(e *echo.Echo) {

	go e.GET("/", func(c echo.Context) error {

		command := Create.CreateCommand{FirstName: "", LastName: ""}

		result := Create.CreateCommandHandler(&command)

		if result == true {
			return c.String(http.StatusOK, "Hello, World!")
		}

		return c.String(http.StatusBadRequest, "Hello, World!")

	})

}
