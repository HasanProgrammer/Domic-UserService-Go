package concretes

import (
	CommonInterface "domic.domain/commons/contracts/interfaces"
	"domic.domain/commons/dtos"
	"domic.domain/user/contracts/contracts"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type UnitOfWork struct {
	db *gorm.DB
	tx *gorm.DB
}

func (unitOfWork *UnitOfWork) StartTransaction() *dtos.Result[bool] {

	unitOfWork.tx = unitOfWork.db.Begin()

	if unitOfWork.tx.Error != nil {
		return &dtos.Result[bool]{
			Result: false,
			Errors: []error{unitOfWork.tx.Error},
		}
	}

	return &dtos.Result[bool]{Result: true}
}

func (unitOfWork *UnitOfWork) Commit() *dtos.Result[bool] {

	commitResult := unitOfWork.tx.Commit()

	if commitResult.Error != nil {
		return &dtos.Result[bool]{
			Result: false,
			Errors: []error{commitResult.Error},
		}
	}

	return &dtos.Result[bool]{Result: true}

}

func (unitOfWork *UnitOfWork) RollBack() *dtos.Result[bool] {

	rollBackResult := unitOfWork.tx.Rollback()

	if rollBackResult.Error != nil {
		return &dtos.Result[bool]{
			Result: false,
			Errors: []error{rollBackResult.Error},
		}
	}

	return &dtos.Result[bool]{Result: true}

}

func (unitOfWork *UnitOfWork) UserRepository() contracts.IUserRepository {

	if unitOfWork.tx == nil {
		return NewUserRepository(unitOfWork.db)
	} else {
		return NewUserRepository(unitOfWork.tx)
	}

}

func (unitOfWork *UnitOfWork) EventRepository() CommonInterface.IEventRepository {

	if unitOfWork.tx == nil {
		return NewEventRepository(unitOfWork.db)
	} else {
		return NewEventRepository(unitOfWork.tx)
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
