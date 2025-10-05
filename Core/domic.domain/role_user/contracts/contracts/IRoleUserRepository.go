package contracts

import (
	"domic.domain/commons/contracts/interfaces"
	"domic.domain/role/entities"
)

type IRoleUserRepository interface {
	interfaces.IRepository[string, *entities.Role]
}
