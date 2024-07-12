package Controllers

import (
	"domic.usecase/UserUseCase/Commands/Create"
)

type UserController struct {
	CreateCommandHandler *Create.CreateCommandHandler
}

func (userController *UserController) Get() bool {

	command := Create.CreateCommand{FirstName: "حسن", LastName: "کرمی محب"}

	result := userController.CreateCommandHandler.Handle(&command)

	return result

}

func NewUserController(CreateCommandHandler *Create.CreateCommandHandler) *UserController {
	return &UserController{
		CreateCommandHandler: CreateCommandHandler,
	}
}
