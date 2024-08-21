package UseCaseUserQuery

import (
	"Dotris.Domain/User/Contracts"
	"Dotris.Domain/User/Entities"
	"Dotris.UseCase/Commons/DTOs"
)

type FetchAllQueryHandler struct {
	query          *FetchAllQuery
	userRepository DomainUserContract.IUserRepository
}

func (handler *FetchAllQueryHandler) Handle() (*UseCaseCommonDTO.PaginationResponse[*DomainUserEntity.User], error) {

	count, countError := handler.userRepository.Count()

	if countError == nil {
		return nil, countError
	}

	result, findAllError := handler.userRepository.FindAll(handler.query.PageSize, handler.query.PageIndex)

	if findAllError == nil {
		return nil, findAllError
	}

	return &UseCaseCommonDTO.PaginationResponse[*DomainUserEntity.User]{
		PageIndex: handler.query.PageIndex,
		PageSize:  handler.query.PageSize,
		TotalItem: count,
		Items:     result,
		HasNext:   false,
		HasPrev:   false,
	}, nil
}

func NewFetchAllQueryHandler(Query *FetchAllQuery, UserRepository DomainUserContract.IUserRepository) *FetchAllQueryHandler {
	return &FetchAllQueryHandler{query: Query, userRepository: UserRepository}
}
