package Interfaces

import (
	"domic.domain/Commons/Contracts/Interfaces"
	"domic.domain/Commons/DTOs"
	"domic.domain/User/Entities"
)

type IUserRepository[TIdentity any] interface {
	Interfaces.IRepository[TIdentity, *Entities.User[TIdentity]]

	IsExistById(username string) DTOs.Results[bool]
	IsExistByUsername(username string) DTOs.Results[bool]
	IsExistByPhoneNumber(username string) DTOs.Results[bool]
	IsExistByEmail(username string) DTOs.Results[bool]
}
