package DomainUserContract

import (
	"Dotris.Domain/Commons/Contracts"
	"Dotris.Domain/User/Entities"
)

type IUserRepository interface {
	DomainCommonContract.IRepository[*string, *DomainUserEntity.User]
}
