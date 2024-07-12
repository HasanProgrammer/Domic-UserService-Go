package Create

import (
	"domic.domain/User/Contracts"
	"domic.domain/User/Entities"
)

type CreateCommandHandler struct {
	UserRepository Contracts.IUserRepository
}

func (createCommandHandler *CreateCommandHandler) Handle(command *CreateCommand) bool {

	result := createCommandHandler.UserRepository.Add(Entities.NewUser(command.FirstName, command.LastName))

	return result

}

func NewCreateCommandHandler(UserRepository Contracts.IUserRepository) *CreateCommandHandler {
	return &CreateCommandHandler{
		UserRepository: UserRepository,
	}
}
