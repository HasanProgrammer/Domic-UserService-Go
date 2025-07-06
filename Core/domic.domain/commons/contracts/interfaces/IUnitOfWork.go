package contracts

import (
	"domic.domain/commons/DTOs"
	"domic.domain/user/contracts/contracts"
)

type IUnitOfWork interface {
	StartTransaction() *DTOs.Result[bool]
	Commit() *DTOs.Result[bool]
	RollBack() *DTOs.Result[bool]

	UserRepository() contracts.IUserRepository
	EventRepository() IEventRepository
}
