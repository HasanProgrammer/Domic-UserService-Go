package UseCaseUserCommand

import (
	"Domic.Domain/Commons/Contracts"
	"Domic.Domain/User/Contracts"
	"Domic.Domain/User/Entities"
)

type CreateCommandHandler struct {
	command        *CreateCommand
	unitOfWork     DomainCommonContract.IUnitOfWork
	userRepository DomainUserContract.IUserRepository[string]
}

func NewCreateCommandHandler(command *CreateCommand, unitOfWork DomainCommonContract.IUnitOfWork, userRepository DomainUserContract.IUserRepository[string]) *CreateCommandHandler {
	return &CreateCommandHandler{userRepository: userRepository, unitOfWork: unitOfWork, command: command}
}

func (commandHandler *CreateCommandHandler) Handle() error {

	user, err := DomainUserEntity.NewUser[string](
		"",
		commandHandler.command.FirstName,
		commandHandler.command.LastName,
		commandHandler.command.Username,
		commandHandler.command.Password,
		commandHandler.command.Email,
		"",
		"Admin",
	)

	if err != nil {
		return err
	}

	err = commandHandler.userRepository.Add(user)

	if err != nil {
		return err
	}

	err = commandHandler.unitOfWork.CommitTransaction()

	if err != nil {
		err = commandHandler.unitOfWork.RollbackTransaction()
	}

	return err
}
