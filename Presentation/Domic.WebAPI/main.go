package main

import (
	"Domic.WebAPI/Controllers"
	_ "Domic.WebAPI/docs"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
)

// @title Swagger UserService API
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name "Authorization"
func main() {

	e := echo.New()

	//GlobalMiddlewares

	//GlobalMiddlewares

	/*---------------------------------------------------------------*/

	//UserController

	userController := WebAPIController.NewUserController()

	apiGroup := e.Group("/api/v1/" /*, WebAPIMiddleware.Auth*/)

	userApiGroup := apiGroup.Group("users")

	userApiGroup.POST("", userController.Create)
	userApiGroup.PATCH("", userController.Update)
	userApiGroup.PATCH("signin", userController.SignIn)

	//UserController

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8080"))

}
