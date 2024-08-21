package DomainCommonContract

import "Dotris.Domain/Commons/DTOs"

type IRepository[TIdentity any, TEntity interface{}] interface {
	Add(entity TEntity) error
	AddAsync(entity TEntity, result chan error)
	Change(entity TEntity) error
	ChangeAsync(entity TEntity, result chan error)
	Remove(entity TEntity) error
	RemoveAsync(entity TEntity, result chan error)
	FindById(id TIdentity) (TEntity, error)
	FindByIdAsync(id TIdentity, result chan DomainCommonDTO.Result[TEntity])
	FindAll(pageSize int64, pageIndex int64) ([]TEntity, error)
	FindAllAsync(pageSize int64, pageIndex int64, result chan DomainCommonDTO.Result[[]TEntity])
	Count(conditions ...interface{}) (int64, error)
}
