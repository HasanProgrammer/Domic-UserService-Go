package UseCaseUserQuery

import (
	"Domic.Domain/Commons/DTOs"
	"Domic.Domain/User/Contracts"
	"Domic.Domain/User/Entities"
)

type FetchAllQueryHandler struct {
	userRepository DomainUserContract.IUserRepository[string]
}

func (handler *FetchAllQueryHandler) Handle(query *FetchAllQuery, result chan DomainCommonDTO.Result[DomainCommonDTO.PaginationResponse[[]*DomainUserEntity.User[string]]]) {

	queryChannel := make(chan DomainCommonDTO.PaginationResponse[[]*DomainUserEntity.User[string]])

	go handler.userRepository.FindAll(&query.PaginationRequest, queryChannel)

	resultQuery := <-queryChannel

	result <- DomainCommonDTO.Result[DomainCommonDTO.PaginationResponse[[]*DomainUserEntity.User[string]]]{
		OutPut: resultQuery,
	}
}

func NewFetchAllQueryHandler(UserRepository DomainUserContract.IUserRepository[string]) *FetchAllQueryHandler {
	return &FetchAllQueryHandler{userRepository: UserRepository}
}
