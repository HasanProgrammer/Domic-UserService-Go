package Interfaces

import "domic.domain/Commons/DTOs"

type IUnitOfWork interface {
	BeginTransaction() *DTOs.Result[bool]
	Commit() *DTOs.Result[bool]
	RollBack() *DTOs.Result[bool]
}
