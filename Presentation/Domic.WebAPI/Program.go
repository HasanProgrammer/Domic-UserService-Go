package main

import (
	"UserService/Controllers"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type Dependencies struct {
	Contract    interface{}
	DriverClass interface{}
	Token       string
}

func main() {

	e := echo.New()

	//DI

	container := dig.New()

	var dependencies = []Dependencies{}

	for _, dep := range dependencies {
		err := container.Provide(dep.DriverClass, dig.As(dep.Contract), dig.Name(dep.Token))
		e.Logger.Printf(err.Error())
	}

	//DI

	Controllers.RegisterUserControllerActions(e)

	e.Logger.Fatal(e.Start(":1323"))

}
