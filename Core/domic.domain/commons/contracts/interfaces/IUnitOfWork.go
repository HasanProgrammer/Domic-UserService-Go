package interfaces

import (
	"context"
	"domic.domain/commons/dtos"
	RoleContracts "domic.domain/role/contracts/contracts"
	RoleUserContracts "domic.domain/role_user/contracts/contracts"
	UserContarcts "domic.domain/user/contracts/contracts"
)

type IUnitOfWork interface {
	StartTransaction(ctx context.Context) *dtos.Result[bool]
	Commit(ctx context.Context) *dtos.Result[bool]
	RollBack(ctx context.Context) *dtos.Result[bool]

	RoleUserRepository() RoleUserContracts.IRoleUserRepository
	UserRepository() UserContarcts.IUserRepository
	RoleRepository() RoleContracts.IRoleRepository
	EventRepository() IEventRepository
}
