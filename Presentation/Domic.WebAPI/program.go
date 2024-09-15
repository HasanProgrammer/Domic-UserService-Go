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

	go e.POST("users", userController.Create)

	e.Logger.Fatal(e.Start(":8080"))

}
