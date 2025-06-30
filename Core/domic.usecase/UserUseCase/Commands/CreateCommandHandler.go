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
	unitOfWork  CommonInterface.IUnitOfWork
	idGenerator CommonInterface.IIdentityGenerator
}

func (handler *CreateCommandHandler) Handle(command *CreateCommand) *DTOs.Result[bool] {

	//validation

	validateResult := commandValidation(command, handler.unitOfWork.UserRepository())

	if !validateResult.Result {
		return validateResult
	}

	//endValidation

	newUser := Entities.NewUser(handler.idGenerator,
		command.FirstName, command.LastName, command.Username, command.Password,
		command.EMail, command.CreatedBy, command.CreatedRole,
	)

	txResult := handler.unitOfWork.StartTransaction()

	if !txResult.Result {
		return txResult
	}

	handler.unitOfWork.EventRepository().AddRange(newUser.GetEvents())
	handler.unitOfWork.UserRepository().Add(newUser)

	commitResult := handler.unitOfWork.Commit()

	if !commitResult.Result {
		return handler.unitOfWork.RollBack()
	}

	return commitResult
}

func NewCreateCommandHandler(
	unitOfWork CommonInterface.IUnitOfWork,
	idGenerator CommonInterface.IIdentityGenerator,
) *CreateCommandHandler {
	return &CreateCommandHandler{
		unitOfWork:  unitOfWork,
		idGenerator: idGenerator,
	}
}

/*-------------------------------------------------------------------*/

func commandValidation(command *CreateCommand, repository UserInterface.IUserRepository) *DTOs.Result[bool] {

	targetUser := repository.IsExistByUsername(command.Username)

	if !targetUser.Result {
		return &DTOs.Result[bool]{
			Errors: []error{errors.New("نام کاربری قبلا انتخاب شده است")},
			Result: false,
		}
	}

	targetUser = repository.IsExistByPhoneNumber(command.PhoneNumber)

	if !targetUser.Result {
		return &DTOs.Result[bool]{
			Errors: []error{errors.New("شماره تماس قبلا انتخاب شده است")},
			Result: false,
		}
	}

	targetUser = repository.IsExistByEmail(command.EMail)

	if !targetUser.Result {
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
