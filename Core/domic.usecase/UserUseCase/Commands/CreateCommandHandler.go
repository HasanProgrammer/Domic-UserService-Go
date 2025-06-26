package UserUseCase

import (
	CommonInterface "domic.domain/Commons/Contracts/Interfaces"
	"domic.domain/Commons/DTOs"
	UserInterface "domic.domain/User/Contracts/Interfaces"
	"domic.domain/User/Entities"
	"errors"
)

type CreateCommand struct {
	FirstName   string
	LastName    string
	Username    string
	Password    string
	PhoneNumber string
	EMail       string
	Description string
	Roles       []string
	Permissions []string

	//audit
	CreatedBy   string
	CreatedRole string
}

type CreateCommandHandler struct {
	unitOfWork      CommonInterface.IUnitOfWork
	userRepository  UserInterface.IUserRepository[string]
	eventRepository CommonInterface.IEventRepository[string]
	idGenerator     CommonInterface.IIdentityGenerator
}

func (createCommandHandler *CreateCommandHandler) Handle(command *CreateCommand) *DTOs.Result[bool] {

	//validation

	validateResult := commandValidation(command, createCommandHandler.userRepository)

	if !validateResult.Result {
		return validateResult
	}

	//endValidation

	newUser := Entities.NewUser(createCommandHandler.idGenerator,
		command.FirstName, command.LastName, command.Username, command.Password,
		command.EMail, command.CreatedBy, command.CreatedRole,
	)

	createCommandHandler.eventRepository.AddRange(newUser.GetEvents())

	return createCommandHandler.userRepository.Add(newUser)

}

func NewCreateCommandHandler(
	userRepository UserInterface.IUserRepository[string],
	eventRepository CommonInterface.IEventRepository[string],
	idGenerator CommonInterface.IIdentityGenerator,
) *CreateCommandHandler {
	return &CreateCommandHandler{
		userRepository:  userRepository,
		eventRepository: eventRepository,
		idGenerator:     idGenerator,
	}
}

/*-------------------------------------------------------------------*/

func commandValidation(command *CreateCommand, repository UserInterface.IUserRepository[string]) *DTOs.Result[bool] {

	targetUser := repository.IsExistByUsername(command.Username)

	if len(targetUser.Errors) > 0 {
		return &DTOs.Result[bool]{
			Errors: targetUser.Errors,
			Result: false,
		}
	} else if targetUser.Result {
		return &DTOs.Result[bool]{
			Errors: []error{errors.New("نام کاربری قبلا انتخاب شده است")},
			Result: false,
		}
	}

	targetUser = repository.IsExistByPhoneNumber(command.PhoneNumber)

	if len(targetUser.Errors) > 0 {
		return &DTOs.Result[bool]{
			Errors: targetUser.Errors,
			Result: false,
		}
	} else if targetUser.Result {
		return &DTOs.Result[bool]{
			Errors: []error{errors.New("شماره تماس قبلا انتخاب شده است")},
			Result: false,
		}
	}

	targetUser = repository.IsExistByEmail(command.EMail)

	if len(targetUser.Errors) > 0 {
		return &DTOs.Result[bool]{
			Errors: targetUser.Errors,
			Result: false,
		}
	} else if targetUser.Result {
		return &DTOs.Result[bool]{
			Errors: []error{errors.New("پست الکترونیکی قبلا انتخاب شده است")},
			Result: false,
		}
	}

	return &DTOs.Result[bool]{
		Errors: []error{},
		Result: true,
	}
}
