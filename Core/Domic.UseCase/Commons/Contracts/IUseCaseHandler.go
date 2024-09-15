package UseCaseCommonContract

import "Domic.Domain/Commons/DTOs"

type IUseCaseHandler[TRequest any, TResult any] interface {
	Handle(request TRequest) DomainCommonDTO.Results[TResult]
}
