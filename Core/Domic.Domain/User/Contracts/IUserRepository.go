package DomainUserContract

import (
	"Domic.Domain/Commons/Contracts"
	"Domic.Domain/User/Entities"
)

type IUserRepository interface {
	DomainCommonContract.IRepository[string, *DomainUserEntity.User]
}
