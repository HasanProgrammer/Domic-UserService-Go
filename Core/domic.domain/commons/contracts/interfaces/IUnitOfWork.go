package interfaces

import (
	"context"
	"domic.domain/commons/dtos"
	"domic.domain/user/contracts/contracts"
)

type IUnitOfWork interface {
	StartTransaction(ctx context.Context) *dtos.Result[bool]
	Commit(ctx context.Context) *dtos.Result[bool]
	RollBack(ctx context.Context) *dtos.Result[bool]

	UserRepository() contracts.IUserRepository
	EventRepository() IEventRepository
}
