package UseCaseCommonContract

import DomainCommonDTO "Domic.Domain/Commons/DTOs"

type IUseCaseHandler[TRequest any, TResult any] interface {
	Handle(request TRequest, result chan DomainCommonDTO.Result[TResult])
}
