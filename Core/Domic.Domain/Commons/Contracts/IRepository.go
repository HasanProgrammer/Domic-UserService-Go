package DomainCommonContract

import (
	"Domic.Domain/Commons/DTOs"
)

type IRepository[TIdentity any, TEntity any] interface {
	Add(entity TEntity, result chan DomainCommonDTO.Result[bool])
	AddRange(entity []TEntity, result chan DomainCommonDTO.Result[bool])
	Change(entity TEntity, result chan DomainCommonDTO.Result[bool])
	Remove(entity TEntity, result chan DomainCommonDTO.Result[bool])
	FindById(id TIdentity, result chan DomainCommonDTO.Result[TEntity])
	FindAll(paginationRequest *DomainCommonDTO.PaginationRequest, result chan DomainCommonDTO.PaginationResponse[[]TEntity])
	Count(result chan DomainCommonDTO.Result[int64], conditions ...interface{})
}
