package UseCaseUserCommand

import (
	"Domic.Domain/Commons/Contracts"
	"Domic.Domain/Commons/DTOs"
	"Domic.Domain/Commons/Entities"
	"Domic.Domain/User/Contracts"
	"Domic.Domain/User/Entities"
)

type CreateCommandHandler struct {
	unitOfWork      DomainCommonContract.IUnitOfWork
	userRepository  DomainUserContract.IUserRepository[string]
	eventRepository DomainCommonContract.IRepository[string, *DomainCommonEntity.Event[string]]
}

func (commandHandler *CreateCommandHandler) Handle(command *CreateCommand, result chan DomainCommonDTO.Results[bool]) {

	var errors []error

	user, err := DomainUserEntity.NewUser[string](
		"",
		command.FirstName,
		command.LastName,
		command.Username,
		command.Password,
		command.Email,
		"",
		"Admin",
	)

	if err != nil {
		result <- DomainCommonDTO.Results[bool]{
			Errors: append(errors, err),
			OutPut: false,
		}

		return
	}

	//user creation

	queryChannel := make(chan DomainCommonDTO.Result[bool])

	go commandHandler.userRepository.Add(user, queryChannel)

	addUserQueryResult := <-queryChannel

	if addUserQueryResult.Error != nil {
		errors = append(errors, addUserQueryResult.Error)
	}

	//event creation

	go commandHandler.eventRepository.AddRange(user.Events(), queryChannel)

	addEventQueryResult := <-queryChannel

	if addEventQueryResult.Error != nil {
		errors = append(errors, addEventQueryResult.Error)
	}

	//transaction section

	if len(errors) > 0 {

		go commandHandler.unitOfWork.RollbackTransaction(queryChannel)

		transactionResult := <-queryChannel

		if transactionResult.Error != nil {
			result <- DomainCommonDTO.Results[bool]{
				Errors: append(errors, transactionResult.Error),
				OutPut: false,
			}

			return
		}

		result <- DomainCommonDTO.Results[bool]{
			Errors: errors,
			OutPut: false,
		}

	}

	go commandHandler.unitOfWork.CommitTransaction(queryChannel)

	transactionResult := <-queryChannel

	if transactionResult.Error != nil {
		errors = append(errors, transactionResult.Error)
	}

	if len(errors) > 0 {
		result <- DomainCommonDTO.Results[bool]{
			Errors: errors,
			OutPut: false,
		}

		return
	}

	result <- DomainCommonDTO.Results[bool]{
		Errors: nil,
		OutPut: true,
	}

}

func NewCreateCommandHandler(
	UnitOfWork DomainCommonContract.IUnitOfWork,
	UserRepository DomainUserContract.IUserRepository[string],
	EventRepository DomainCommonContract.IRepository[string, *DomainCommonEntity.Event[string]],
) *CreateCommandHandler {

	return &CreateCommandHandler{
		unitOfWork:      UnitOfWork,
		userRepository:  UserRepository,
		eventRepository: EventRepository,
	}

}
