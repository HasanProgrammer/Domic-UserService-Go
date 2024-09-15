package InfrastructureConcrete

import (
	"Domic.Domain/Commons/DTOs"
	"gorm.io/gorm"
)

type UnitOfWork struct {
	db          *gorm.DB
	transaction *gorm.DB
}

func (u *UnitOfWork) GetTransaction() *gorm.DB {
	return u.transaction
}

func (u *UnitOfWork) Commit(result chan DomainCommonDTO.Result[bool]) {

	if u.transaction != nil {

		queryResult := u.transaction.Commit()

		result <- DomainCommonDTO.Result[bool]{
			Error:  queryResult.Error,
			Result: true,
		}

	}

}

func (u *UnitOfWork) Rollback(result chan DomainCommonDTO.Result[bool]) {

	if u.transaction != nil {

		queryResult := u.transaction.Rollback()

		result <- DomainCommonDTO.Result[bool]{
			Error:  queryResult.Error,
			Result: true,
		}

	}

}

func NewUnitOfWork(db *gorm.DB) *UnitOfWork {

	unitOfWork := &UnitOfWork{db: db}

	unitOfWork.transaction = db.Begin()

	return unitOfWork

}
