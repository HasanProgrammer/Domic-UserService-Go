package Concrete

import (
	"domic.domain/Commons/DTOs"
	"domic.domain/User/Contracts/Interfaces"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type UnitOfWork struct {
	db *gorm.DB
	tx *gorm.DB
}

func (unitOfWork *UnitOfWork) StartTransaction() *DTOs.Result[bool] {

	unitOfWork.tx = unitOfWork.db.Begin()

	if unitOfWork.tx.Error != nil {
		return &DTOs.Result[bool]{
			Result: false,
			Errors: []error{unitOfWork.tx.Error},
		}
	}

	return &DTOs.Result[bool]{Result: true}

}

func (unitOfWork *UnitOfWork) Commit() *DTOs.Result[bool] {

	commitResult := unitOfWork.tx.Commit()

	if commitResult.Error != nil {
		return &DTOs.Result[bool]{
			Result: false,
			Errors: []error{commitResult.Error},
		}
	}

	return &DTOs.Result[bool]{Result: true}

}

func (unitOfWork *UnitOfWork) RollBack() *DTOs.Result[bool] {

	rollBackResult := unitOfWork.tx.Rollback()

	if rollBackResult.Error != nil {
		return &DTOs.Result[bool]{
			Result: false,
			Errors: []error{rollBackResult.Error},
		}
	}

	return &DTOs.Result[bool]{Result: true}

}

func (unitOfWork *UnitOfWork) UserRepository() Interfaces.IUserRepository {

	if unitOfWork.tx == nil {
		return NewUserRepository(unitOfWork.db)
	} else {
		return NewUserRepository(unitOfWork.tx)
	}

}

func NewUnitOfWork(connectionString string) (*UnitOfWork, error) {

	db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err == nil {
		return &UnitOfWork{
			db: db,
		}, nil
	}

	return nil, err
}
