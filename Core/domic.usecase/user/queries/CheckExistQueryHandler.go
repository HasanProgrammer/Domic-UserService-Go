package queries

import (
	"context"
	"domic.domain/commons/contracts/interfaces"
	"domic.domain/commons/dtos"
)

type CheckExistQuery struct {
	Id string
}

type CheckExistQueryHandler struct {
	UnitOfWork interfaces.IUnitOfWork
}

func (handler *CheckExistQueryHandler) Handle(query *CheckExistQuery, context context.Context) *dtos.Result[bool] {

	return handler.UnitOfWork.UserRepository().IsExistById(query.Id, context)

}
