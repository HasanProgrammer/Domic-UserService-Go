package main

import (
	"Domic.WebAPI/Controllers"
	"Domic.WebAPI/Middlewares"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	//users

	userApiGroup := e.Group("/users")

	userApiGroup.Use(WebAPIMiddleware.Auth)

	go userApiGroup.POST("users", func(c echo.Context) error {

		userController := WebAPIController.NewUserController()

		return userController.Create(c)

	})

	go userApiGroup.PATCH("users/signin", func(c echo.Context) error {

		userController := WebAPIController.NewUserController()

		return userController.SignIn(c)

	})

	e.Logger.Fatal(e.Start(":8080"))

}
