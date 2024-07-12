package Interfaces

import (
	"domic.domain/Commons/Contracts/Interfaces"
	"domic.domain/User/Entities"
)

type IUserRepository interface {
	Interfaces.IRepository[*Entities.User]
}
