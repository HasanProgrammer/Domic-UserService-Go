package interfaces

import (
	"domic.domain/commons/dtos"
	"domic.domain/user/contracts/contracts"
)

type IUnitOfWork interface {
	StartTransaction() *dtos.Result[bool]
	Commit() *dtos.Result[bool]
	RollBack() *dtos.Result[bool]

	UserRepository() contracts.IUserRepository
	EventRepository() IEventRepository
}
