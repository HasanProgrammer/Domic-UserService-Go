package DomainCommonContract

import "Domic.Domain/Commons/DTOs"

type IUnitOfWork interface {
	Commit(result chan DomainCommonDTO.Result[bool])
	Rollback(result chan DomainCommonDTO.Result[bool])
}
