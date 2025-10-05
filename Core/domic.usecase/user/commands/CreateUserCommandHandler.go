package commands

import (
	"context"
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
	EMail       string
	PhoneNumber string
	ImageUrl    string
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

func (handler *CreateUserCommandHandler) Handle(command *CreateUserCommand, ctx context.Context) *dtos.Result[bool] {

	//validation

	validateResult := _commandValidation(ctx, command, handler.unitOfWork.UserRepository())

	if !validateResult.Result {
		return validateResult
	}

	//endValidation

	newUser := entities.NewUser(handler.idGenerator,
		command.FirstName, command.LastName, command.Username, command.Password,
		command.EMail, command.PhoneNumber, command.ImageUrl, command.Description, command.CreatedBy, command.CreatedRole,
	)

	txResult := handler.unitOfWork.StartTransaction(ctx)

	if !txResult.Result {
		return txResult
	}

	handler.unitOfWork.EventRepository().AddRange(newUser.GetEvents(), ctx)
	handler.unitOfWork.UserRepository().Add(newUser, ctx)

	commitResult := handler.unitOfWork.Commit(ctx)

	if !commitResult.Result {
		return handler.unitOfWork.RollBack(ctx)
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

func _commandValidation(ctx context.Context, command *CreateUserCommand, repository UserInterface.IUserRepository) *dtos.Result[bool] {

	targetUser := repository.IsExistByUsername(command.Username, ctx)

	if !targetUser.Result {
		return &dtos.Result[bool]{
			Errors: []error{errors.New("نام کاربری قبلا انتخاب شده است")},
			Result: false,
		}
	}

	targetUser = repository.IsExistByPhoneNumber(command.PhoneNumber, ctx)

	if !targetUser.Result {
		return &dtos.Result[bool]{
			Errors: []error{errors.New("شماره تماس قبلا انتخاب شده است")},
			Result: false,
		}
	}

	targetUser = repository.IsExistByEmail(command.EMail, ctx)

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
