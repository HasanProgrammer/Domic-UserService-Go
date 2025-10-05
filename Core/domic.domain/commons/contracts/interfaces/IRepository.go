package interfaces

import (
	"context"
	"domic.domain/commons/dtos"
)

type IRepository[TIdentity any, TEntity any] interface {
	Add(entity TEntity, context context.Context) *dtos.Result[bool]
	AddRange(entities []TEntity, context context.Context) *dtos.Result[bool]
	Change(entity TEntity, context context.Context) *dtos.Result[bool]
	ChangeRange(entities []TEntity, context context.Context) *dtos.Result[bool]
	Remove(entity TEntity, context context.Context) *dtos.Result[bool]
	RemoveRange(entities []TEntity, context context.Context) *dtos.Result[bool]
	FindById(id TIdentity, context context.Context) *dtos.Result[TEntity]
	FindAll(paginationRequest *dtos.PaginationRequest, context context.Context) *dtos.Result[*dtos.PaginationResponse[TEntity]]
}
