package commands

import (
	CommonInterface "domic.domain/commons/contracts/interfaces"
	"domic.domain/commons/dtos"
	UserInterface "domic.domain/user/contracts/contracts"
	"domic.domain/user/entities"
	"errors"
)

type CreateUserCommand struct {
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

type CreateUserCommandHandler struct {
	unitOfWork  CommonInterface.IUnitOfWork
	idGenerator CommonInterface.IIdentityGenerator
}

func (handler *CreateUserCommandHandler) Handle(command *CreateUserCommand) *dtos.Result[bool] {

	//validation

	validateResult := _commandValidation(command, handler.unitOfWork.UserRepository())

	if !validateResult.Result {
		return validateResult
	}

	//endValidation

	newUser := entities.NewUser(handler.idGenerator,
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

func NewCreateUserCommandHandler(
	unitOfWork CommonInterface.IUnitOfWork,
	idGenerator CommonInterface.IIdentityGenerator,
) *CreateUserCommandHandler {
	return &CreateUserCommandHandler{
		unitOfWork:  unitOfWork,
		idGenerator: idGenerator,
	}
}

/*-------------------------------------------------------------------*/

func _commandValidation(command *CreateUserCommand, repository UserInterface.IUserRepository) *dtos.Result[bool] {

	targetUser := repository.IsExistByUsername(command.Username)

	if !targetUser.Result {
		return &dtos.Result[bool]{
			Errors: []error{errors.New("نام کاربری قبلا انتخاب شده است")},
			Result: false,
		}
	}

	targetUser = repository.IsExistByPhoneNumber(command.PhoneNumber)

	if !targetUser.Result {
		return &dtos.Result[bool]{
			Errors: []error{errors.New("شماره تماس قبلا انتخاب شده است")},
			Result: false,
		}
	}

	targetUser = repository.IsExistByEmail(command.EMail)

	if !targetUser.Result {
		return &dtos.Result[bool]{
			Errors: []error{errors.New("پست الکترونیکی قبلا انتخاب شده است")},
			Result: false,
		}
	}

	return &dtos.Result[bool]{
		Errors: []error{},
		Result: true,
	}
}
