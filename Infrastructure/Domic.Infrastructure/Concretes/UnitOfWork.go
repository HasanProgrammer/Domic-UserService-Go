package InfrastructureConcrete

import (
	"Domic.Domain/Commons/DTOs"
	"gorm.io/gorm"
)

type UnitOfWork struct {
	transaction *gorm.DB
	db          *gorm.DB
}

func (u *UnitOfWork) Transaction() *gorm.DB {
	return u.transaction
}

func (u *UnitOfWork) CommitTransaction(result chan DomainCommonDTO.Result[bool]) {

	if u.transaction != nil {

		go func() {
			queryResult := u.transaction.Commit()

			result <- DomainCommonDTO.Result[bool]{
				Error:  queryResult.Error,
				Result: true,
			}
		}()

	}

}

func (u *UnitOfWork) RollbackTransaction(result chan DomainCommonDTO.Result[bool]) {

	if u.transaction != nil {

		go func() {
			queryResult := u.transaction.Rollback()

			result <- DomainCommonDTO.Result[bool]{
				Error:  queryResult.Error,
				Result: true,
			}
		}()

	}

}

func NewUnitOfWork(db *gorm.DB) *UnitOfWork {

	unitOfWork := &UnitOfWork{db: db}

	unitOfWork.transaction = db.Begin()

	return unitOfWork

}
