package contracts

import (
	"context"
	"domic.domain/commons/contracts/interfaces"
	"domic.domain/commons/dtos"
	"domic.domain/user/entities"
)

type IUserRepository interface {
	interfaces.IRepository[string, *entities.User]

	IsExistById(id string, context context.Context) *dtos.Result[bool]
	IsExistByUsername(username string, context context.Context) *dtos.Result[bool]
	IsExistByPhoneNumber(phoneNumber string, context context.Context) *dtos.Result[bool]
	IsExistByEmail(email string, context context.Context) *dtos.Result[bool]
}
