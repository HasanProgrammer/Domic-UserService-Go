package UseCaseUserCommand

import (
	"Dotris.Domain/User/Contracts"
	"Dotris.Domain/User/Entities"
)

type CreateCommandHandler struct {
	command        *CreateCommand
	userRepository DomainUserContract.IUserRepository
}

func NewCreateCommandHandler(UserRepository DomainUserContract.IUserRepository, Command *CreateCommand) *CreateCommandHandler {
	return &CreateCommandHandler{
		command:        Command,
		userRepository: UserRepository,
	}
}

func (commandHandler *CreateCommandHandler) Handle() (bool, error) {

	user, e := DomainUserEntity.NewUser(
		"",
		commandHandler.command.FirstName,
		commandHandler.command.LastName,
		commandHandler.command.Username,
		commandHandler.command.Password,
		commandHandler.command.Email,
	)

	commandHandler.userRepository.Add(user)

	return true, e
}
