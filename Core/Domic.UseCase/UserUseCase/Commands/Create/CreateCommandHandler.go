package Create

import (
	"domic.domain/User/Contracts/Interfaces"
	"domic.domain/User/Entities"
)

type CreateCommandHandler struct {
	userRepository Interfaces.IUserRepository
}

func (createCommandHandler *CreateCommandHandler) HandleAsync(command *CreateCommand, result chan bool) {

	newUser := Entities.NewUser(command.FirstName, command.LastName)

	channel := make(chan bool)

	go createCommandHandler.userRepository.AddAsync(newUser, channel)

	res := <-channel

	result <- res

}

func NewCreateCommandHandler(UserRepository Interfaces.IUserRepository) *CreateCommandHandler {
	return &CreateCommandHandler{
		userRepository: UserRepository,
	}
}
