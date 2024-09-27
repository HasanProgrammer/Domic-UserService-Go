package UseCaseUserCommand

import (
	"Domic.Domain/Commons/Contracts"
	"Domic.Domain/Commons/DTOs"
	"Domic.Domain/Commons/Entities"
	"Domic.Domain/User/Contracts"
	"Domic.Domain/User/Entities"
)

type CreateCommandHandler struct {
	serializer      DomainCommonContract.ISerializer
	idGenerator     DomainCommonContract.IGlobalIdentityGenerator
	unitOfWork      DomainCommonContract.IUnitOfWork
	userRepository  DomainUserContract.IUserRepository
	eventRepository DomainCommonContract.IRepository[string, *DomainCommonEntity.Event]
}

func (commandHandler *CreateCommandHandler) Handle(command *CreateCommand) DomainCommonDTO.Results[bool] {

	var errors []error

	user, err := DomainUserEntity.NewUser(
		commandHandler.idGenerator,
		commandHandler.serializer,
		command.FirstName,
		command.LastName,
		command.Username,
		command.Password,
		command.Email,
		"",
		"Admin",
	)

	if err != nil {
		return DomainCommonDTO.Results[bool]{
			Errors: append(errors, err),
			Result: false,
		}
	}

	queryChannel := make(chan DomainCommonDTO.Result[bool])

	//user creation

	go commandHandler.userRepository.Add(user, queryChannel)

	addUserQueryResult := <-queryChannel

	if addUserQueryResult.Error != nil {
		errors = append(errors, addUserQueryResult.Error)
	}

	//event creation

	go commandHandler.eventRepository.AddRange(user.GetEvents(), queryChannel)

	addEventQueryResult := <-queryChannel

	if addEventQueryResult.Error != nil {
		errors = append(errors, addEventQueryResult.Error)
	}

	//transaction section

	if len(errors) > 0 {

		go commandHandler.unitOfWork.Rollback(queryChannel)

		transactionResult := <-queryChannel

		if transactionResult.Error != nil {
			return DomainCommonDTO.Results[bool]{
				Errors: append(errors, transactionResult.Error),
				Result: false,
			}
		}

		return DomainCommonDTO.Results[bool]{
			Errors: errors,
			Result: false,
		}

	}

	go commandHandler.unitOfWork.Commit(queryChannel)

	transactionResult := <-queryChannel

	if transactionResult.Error != nil {
		errors = append(errors, transactionResult.Error)
	}

	if len(errors) > 0 {
		return DomainCommonDTO.Results[bool]{
			Errors: errors,
			Result: false,
		}
	}

	return DomainCommonDTO.Results[bool]{
		Errors: nil,
		Result: true,
	}

}

func NewCreateCommandHandler(
	IdGenerator DomainCommonContract.IGlobalIdentityGenerator,
	Serializer DomainCommonContract.ISerializer,
	UnitOfWork DomainCommonContract.IUnitOfWork,
	UserRepository DomainUserContract.IUserRepository,
	EventRepository DomainCommonContract.IRepository[string, *DomainCommonEntity.Event],
) *CreateCommandHandler {

	return &CreateCommandHandler{
		idGenerator:     IdGenerator,
		serializer:      Serializer,
		unitOfWork:      UnitOfWork,
		userRepository:  UserRepository,
		eventRepository: EventRepository,
	}

}
