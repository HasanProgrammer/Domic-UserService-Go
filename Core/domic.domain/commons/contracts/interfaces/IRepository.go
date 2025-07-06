package contracts

import (
	"domic.domain/commons/DTOs"
)

type IRepository[TIdentity any, TEntity any] interface {
	Add(entity TEntity) *DTOs.Result[bool]
	AddRange(entities []TEntity) *DTOs.Result[bool]
	Change(entity TEntity) *DTOs.Result[bool]
	ChangeRange(entities []TEntity) *DTOs.Result[bool]
	Remove(entity TEntity) *DTOs.Result[bool]
	RemoveRange(entities []TEntity) *DTOs.Result[bool]
	FindById(id TIdentity) *DTOs.Result[TEntity]
	FindAll(paginationRequest *DTOs.PaginationRequest) *DTOs.Result[*DTOs.PaginationResponse[TEntity]]
}
