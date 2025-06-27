package Interfaces

import (
	"domic.domain/Commons/DTOs"
	"domic.domain/User/Contracts/Interfaces"
)

type IUnitOfWork interface {
	StartTransaction() *DTOs.Result[bool]
	Commit() *DTOs.Result[bool]
	RollBack() *DTOs.Result[bool]

	UserRepository() Interfaces.IUserRepository
}
