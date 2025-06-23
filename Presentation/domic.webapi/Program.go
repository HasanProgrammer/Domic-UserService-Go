package main

import (
	"UserService/Controllers"
	"domic.infrastructure"
	"domic.usecase/UserUseCase/Commands/Create"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {

	e := echo.New()

	/*---------------------------------------------------------------*/

	//APIs

	go e.GET("/users/add/:FirstName/:LastName", func(c echo.Context) error {

		userCreateCommandHandler := Create.NewCreateCommandHandler(DomicInfrastructure.NewUserRepository())

		userController := Controllers.NewUserController(userCreateCommandHandler)

		channel := make(chan bool)

		go userController.AddAsync(c, channel)

		result := <-channel

		if result == true {
			return c.JSON(http.StatusOK, "Success insertion")
		}

		return c.JSON(http.StatusBadRequest, "Hello, World!")

	})

	//APIs

	/*---------------------------------------------------------------*/

	e.Logger.Fatal(e.Start(":1996"))
}
