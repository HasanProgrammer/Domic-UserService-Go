package Interfaces

import (
	"domic.domain/Commons/DTOs"
)

type IRepository[TIdentity any, TEntity any] interface {
	Add(entity TEntity) *DTOs.Results[bool]
	AddRange(entities []TEntity) *DTOs.Results[bool]
	Change(entity TEntity) *DTOs.Results[bool]
	ChangeRange(entities []TEntity) *DTOs.Results[bool]
	Remove(entity TEntity) *DTOs.Results[bool]
	RemoveRange(entities []TEntity) *DTOs.Results[bool]
	FindById(id TIdentity) *DTOs.Results[bool]
	FindAll(paginationRequest *DTOs.PaginationRequest) *DTOs.Results[DTOs.PaginationResponse[TEntity]]
}
