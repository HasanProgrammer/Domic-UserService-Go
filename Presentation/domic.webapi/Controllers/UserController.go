package Controllers

import (
	"domic.usecase/UserUseCase/Commands/Create"
	"github.com/labstack/echo/v4"
	"time"
)

type UserController struct {
	CreateCommandHandler *Create.CreateCommandHandler
}

func (userController *UserController) AddAsync(context echo.Context, result chan<- bool) {
	firstName := context.Param("FirstName")
	lastName := context.Param("LastName")

	command := Create.CreateCommand{FirstName: firstName, LastName: lastName}

	time.Sleep(5 * time.Second)

	channel := make(chan bool)

	go userController.CreateCommandHandler.HandleAsync(&command, channel)

	res := <-channel

	result <- res
}

func NewUserController(CreateCommandHandler *Create.CreateCommandHandler) *UserController {
	return &UserController{
		CreateCommandHandler: CreateCommandHandler,
	}
}
