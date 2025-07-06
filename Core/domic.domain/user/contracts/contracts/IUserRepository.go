package Interfaces

import (
	"domic.domain/Commons/Contracts/Interfaces"
	"domic.domain/Commons/DTOs"
	"domic.domain/user/Entities"
)

type IUserRepository interface {
	Interfaces.IRepository[string, *Entities.User]

	IsExistById(id string) *DTOs.Result[bool]
	IsExistByUsername(username string) *DTOs.Result[bool]
	IsExistByPhoneNumber(phoneNumber string) *DTOs.Result[bool]
	IsExistByEmail(email string) *DTOs.Result[bool]
}
