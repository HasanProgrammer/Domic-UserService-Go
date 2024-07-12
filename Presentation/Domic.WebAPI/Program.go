package main

import (
	"UserService/Routes"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	/*---------------------------------------------------------------*/

	//APIs

	Routes.UserRoutesRegister(e)

	//APIs

	/*---------------------------------------------------------------*/

	e.Logger.Fatal(e.Start(":1996"))
}
