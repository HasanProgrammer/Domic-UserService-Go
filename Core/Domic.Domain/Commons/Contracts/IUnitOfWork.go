package DomainCommonContract

import DomainCommonDTO "Domic.Domain/Commons/DTOs"

type IUnitOfWork interface {
	CommitTransaction(result chan DomainCommonDTO.Result[bool])
	RollbackTransaction(result chan DomainCommonDTO.Result[bool])
}
