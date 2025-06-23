package Interfaces

import "domic.domain/Commons/DTOs"

type IUnitOfWork interface {
	BeginTransaction() *DTOs.Results[bool]
	Commit() *DTOs.Results[bool]
	RollBack() *DTOs.Results[bool]
}
