package UseCaseUserCommand

import (
	DomainCommonContract "Dotris.Domain/Commons/Contracts"
	"Dotris.Domain/User/Contracts"
	"Dotris.Domain/User/Entities"
)

type CreateCommandHandler struct {
	command        *CreateCommand
	unitOfWork     DomainCommonContract.IUnitOfWork
	userRepository DomainUserContract.IUserRepository
}

func NewCreateCommandHandler(command *CreateCommand, unitOfWork DomainCommonContract.IUnitOfWork, userRepository DomainUserContract.IUserRepository) *CreateCommandHandler {
	return &CreateCommandHandler{userRepository: userRepository, unitOfWork: unitOfWork, command: command}
}

func (commandHandler *CreateCommandHandler) Handle() error {

	user, e := DomainUserEntity.NewUser(
		"",
		commandHandler.command.FirstName,
		commandHandler.command.LastName,
		commandHandler.command.Username,
		commandHandler.command.Password,
		commandHandler.command.Email,
	)

	commandHandler.userRepository.Add(user)

	return e
}
