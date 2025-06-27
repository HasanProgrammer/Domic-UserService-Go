package Interfaces

import (
	"domic.domain/Commons/Contracts/Interfaces"
	"domic.domain/Commons/DTOs"
	"domic.domain/User/Entities"
)

type IUserRepository interface {
	Interfaces.IRepository[string, *Entities.User]

	IsExistById(username string) *DTOs.Result[bool]
	IsExistByUsername(username string) *DTOs.Result[bool]
	IsExistByPhoneNumber(username string) *DTOs.Result[bool]
	IsExistByEmail(username string) *DTOs.Result[bool]
}
