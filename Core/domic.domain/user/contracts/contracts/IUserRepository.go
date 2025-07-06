package contracts

import (
	"domic.domain/commons/contracts/interfaces"
	"domic.domain/commons/dtos"
	"domic.domain/user/entities"
)

type IUserRepository interface {
	interfaces.IRepository[string, *entities.User]

	IsExistById(id string) *dtos.Result[bool]
	IsExistByUsername(username string) *dtos.Result[bool]
	IsExistByPhoneNumber(phoneNumber string) *dtos.Result[bool]
	IsExistByEmail(email string) *dtos.Result[bool]
}
