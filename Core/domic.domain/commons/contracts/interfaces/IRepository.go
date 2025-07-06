package interfaces

import (
	"domic.domain/commons/dtos"
)

type IRepository[TIdentity any, TEntity any] interface {
	Add(entity TEntity) *dtos.Result[bool]
	AddRange(entities []TEntity) *dtos.Result[bool]
	Change(entity TEntity) *dtos.Result[bool]
	ChangeRange(entities []TEntity) *dtos.Result[bool]
	Remove(entity TEntity) *dtos.Result[bool]
	RemoveRange(entities []TEntity) *dtos.Result[bool]
	FindById(id TIdentity) *dtos.Result[TEntity]
	FindAll(paginationRequest *dtos.PaginationRequest) *dtos.Result[*dtos.PaginationResponse[TEntity]]
}
