package InfrastructureConcrete

import (
	"gorm.io/gorm"
)

type UnitOfWork struct {
	transaction *gorm.DB
	db          *gorm.DB
}

func (u *UnitOfWork) Transaction() *gorm.DB {
	return u.transaction
}

func (u *UnitOfWork) CommitTransaction() error {

	if u.transaction != nil {
		u.transaction.Commit()
	}

	return nil

}

func (u *UnitOfWork) RollbackTransaction() error {

	if u.transaction != nil {
		u.transaction.Rollback()
	}

	return nil

}

func NewUnitOfWork(db *gorm.DB) *UnitOfWork {

	unitOfWork := &UnitOfWork{db: db}

	unitOfWork.transaction = db.Begin()

	return unitOfWork

}
