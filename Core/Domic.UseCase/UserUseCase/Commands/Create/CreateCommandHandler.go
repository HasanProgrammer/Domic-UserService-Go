package UseCaseUserCommand

import (
	"Domic.Domain/Commons/Contracts"
	"Domic.Domain/Commons/DTOs"
	"Domic.Domain/User/Contracts"
	"Domic.Domain/User/Entities"
)

type CreateCommandHandler struct {
	unitOfWork     DomainCommonContract.IUnitOfWork
	userRepository DomainUserContract.IUserRepository[string]
}

func (commandHandler *CreateCommandHandler) Handle(command *CreateCommand, result chan DomainCommonDTO.Result[bool]) {

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
		result <- DomainCommonDTO.Result[bool]{
			Error:  err,
			OutPut: false,
		}

		return
	}

	queryChannel := make(chan DomainCommonDTO.Result[bool])

	go commandHandler.userRepository.Add(user, queryChannel)

	queryResult := <-queryChannel

	//event [OutBox] processing

	//transaction section

	transactionChannel := make(chan DomainCommonDTO.Result[bool])

	if queryResult.Error != nil {

		go commandHandler.unitOfWork.RollbackTransaction(transactionChannel)

		transactionResult := <-transactionChannel

		if transactionResult.Error != nil {
			result <- DomainCommonDTO.Result[bool]{
				Error:  transactionResult.Error,
				OutPut: false,
			}

			return
		}

		result <- DomainCommonDTO.Result[bool]{
			Error:  queryResult.Error,
			OutPut: false,
		}

	}

	go commandHandler.unitOfWork.CommitTransaction(transactionChannel)

	transactionResult := <-transactionChannel

	if transactionResult.Error != nil {
		result <- DomainCommonDTO.Result[bool]{
			Error:  transactionResult.Error,
			OutPut: false,
		}

		return
	}

	result <- DomainCommonDTO.Result[bool]{
		Error:  nil,
		OutPut: true,
	}

}

func NewCreateCommandHandler(unitOfWork DomainCommonContract.IUnitOfWork, userRepository DomainUserContract.IUserRepository[string]) *CreateCommandHandler {
	return &CreateCommandHandler{userRepository: userRepository, unitOfWork: unitOfWork}
}
