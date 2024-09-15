package UseCaseUserQuery

import (
	"Domic.Domain/Commons/DTOs"
	"Domic.Domain/User/Contracts"
	"Domic.Domain/User/Entities"
)

type FetchAllQueryHandler struct {
	userRepository DomainUserContract.IUserRepository
}

func (handler *FetchAllQueryHandler) Handle(query *FetchAllQuery) DomainCommonDTO.Results[DomainCommonDTO.PaginationResponse[*DomainUserEntity.User]] {

	queryChannel := make(chan DomainCommonDTO.PaginationResponse[*DomainUserEntity.User])

	go handler.userRepository.FindAll(&query.PaginationRequest, queryChannel)

	resultQuery := <-queryChannel

	return DomainCommonDTO.Results[DomainCommonDTO.PaginationResponse[*DomainUserEntity.User]]{
		Result: resultQuery,
		Errors: []error{},
	}
}

func NewFetchAllQueryHandler(UserRepository DomainUserContract.IUserRepository) *FetchAllQueryHandler {
	return &FetchAllQueryHandler{userRepository: UserRepository}
}
