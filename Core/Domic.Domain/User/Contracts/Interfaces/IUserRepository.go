package Contracts

import (
	"domic.domain/Commons"
	"domic.domain/User/Entities"
)

type IUserRepository interface {
	Commons.IRepository[*Entities.User]
}
