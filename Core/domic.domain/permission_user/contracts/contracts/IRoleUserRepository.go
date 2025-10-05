package contracts

import (
	"domic.domain/commons/contracts/interfaces"
	"domic.domain/role_user/entities"
)

type IRoleUserRepository interface {
	interfaces.IRepository[string, *entities.RoleUser]
}
