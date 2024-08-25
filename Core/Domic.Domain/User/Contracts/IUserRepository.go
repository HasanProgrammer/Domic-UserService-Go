package DomainUserContract

import (
	"Domic.Domain/Commons/Contracts"
	"Domic.Domain/User/Entities"
)

type IUserRepository[TIdentity any] interface {
	DomainCommonContract.IRepository[TIdentity, *DomainUserEntity.User[TIdentity]]
}
