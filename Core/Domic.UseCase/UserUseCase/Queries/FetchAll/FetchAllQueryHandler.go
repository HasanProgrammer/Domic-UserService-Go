package UseCaseUserQuery

import (
	"Domic.Domain/Commons/DTOs"
	"Domic.Domain/User/Contracts"
	"Domic.Domain/User/Entities"
)

type FetchAllQueryHandler struct {
	userRepository DomainUserContract.IUserRepository
}

func (handler *FetchAllQueryHandler) Handle(query *FetchAllQuery, result chan DomainCommonDTO.Result[DomainCommonDTO.PaginationResponse[*DomainUserEntity.User]]) {

	queryChannel := make(chan DomainCommonDTO.PaginationResponse[*DomainUserEntity.User])

	go handler.userRepository.FindAll(&query.PaginationRequest, queryChannel)

	resultQuery := <-queryChannel

	result <- DomainCommonDTO.Result[DomainCommonDTO.PaginationResponse[*DomainUserEntity.User]]{
		Result: resultQuery,
		Error:  nil,
	}
}

func NewFetchAllQueryHandler(UserRepository DomainUserContract.IUserRepository) *FetchAllQueryHandler {
	return &FetchAllQueryHandler{userRepository: UserRepository}
}
