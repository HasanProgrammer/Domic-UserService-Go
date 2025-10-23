package interfaces

import (
	"context"
	"domic.domain/commons/dtos"
	RoleContracts "domic.domain/role/contracts/interfaces"
	RoleUserContracts "domic.domain/role_user/contracts/interfaces"
	UserContracts "domic.domain/user/contracts/interfaces"
)

type IUnitOfWork interface {
	StartTransaction(ctx context.Context) *dtos.Result[bool]
	CommitTransaction(ctx context.Context) *dtos.Result[bool]
	RollBackTransaction(ctx context.Context) *dtos.Result[bool]

	RoleUserRepository() RoleUserContracts.IRoleUserRepository
	RoleRepository() RoleContracts.IRoleRepository
	UserRepository() UserContracts.IUserRepository
	EventRepository() IEventRepository
}
