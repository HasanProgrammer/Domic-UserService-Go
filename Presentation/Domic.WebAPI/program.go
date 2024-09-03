package main

import (
	"Domic.Persistence"
	"Domic.WebAPI/Controllers"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	context := Persistence.NewSqlContext("sqlserver://sa:Domic123@127.0.0.1:1633?database=UserService").GetContext()

	//users

	userController := WebAPIController.NewUserController(context)

	go e.GET("users", userController.Create)
	go e.GET("users/{id}", userController.Create)
	go e.POST("users/create", userController.Create)
	go e.PATCH("users/update", userController.Create)
	go e.POST("users/update", userController.Create)

	e.Logger.Fatal(e.Start(":1323"))

}
