package queries

import (
	"context"
	"domic.domain/commons/contracts/interfaces"
	"domic.domain/commons/dtos"
	UserDtos "domic.usecase/user/dtos"
)

type ReadOneQuery struct {
	Id string
}

type ReadOneQueryHandler struct {
	UnitOfWork interfaces.IUnitOfWork
}

func (handler *ReadOneQueryHandler) Handle(query *ReadOneQuery, context context.Context) *dtos.Result[*UserDtos.UserDto] {

	targetUserObject := handler.UnitOfWork.UserRepository().FindById(query.Id, context)

	if len(targetUserObject.Errors) > 0 {
		return &dtos.Result[*UserDtos.UserDto]{
			Errors: targetUserObject.Errors, Result: nil,
		}
	}

	targetUser := targetUserObject.Result

	return &dtos.Result[*UserDtos.UserDto]{
		Result: &UserDtos.UserDto{
			Id:          targetUser.GetId(),
			Username:    targetUser.GetUsername(),
			FirstName:   targetUser.GetFirstName(),
			LastName:    targetUser.GetLastName(),
			Email:       targetUser.GetEmail(),
			PhoneNumber: targetUser.GetPassword(),
		},
	}

}
