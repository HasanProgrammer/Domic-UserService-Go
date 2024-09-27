package main

import (
	"Domic.WebAPI/Controllers"
	"Domic.WebAPI/Middlewares"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	//GlobalMiddlewares

	e.Use(WebAPIMiddleware.Auth)

	//GlobalMiddlewares

	/*---------------------------------------------------------------*/

	//UserController

	userController := WebAPIController.NewUserController()

	userApiGroup := e.Group("api/v1/users")

	userApiGroup.POST("", userController.Create)
	userApiGroup.PATCH("", userController.Update)
	userApiGroup.PATCH("signin", userController.SignIn)

	//UserController

	e.Logger.Fatal(e.Start(":8080"))

}
