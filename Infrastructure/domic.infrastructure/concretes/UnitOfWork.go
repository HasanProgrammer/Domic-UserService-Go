package concretes

import (
	"context"
	CommonContract "domic.domain/commons/contracts/interfaces"
	"domic.domain/commons/dtos"
	RoleContract "domic.domain/role/contracts/interfaces"
	RoleUserContract "domic.domain/role_user/contracts/interfaces"
	UserContract "domic.domain/user/contracts/interfaces"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type UnitOfWork struct {
	db *gorm.DB
	tx *gorm.DB
}

func (unitOfWork *UnitOfWork) StartTransaction(ctx context.Context) *dtos.Result[bool] {

	unitOfWork.tx = unitOfWork.db.Begin().WithContext(ctx)

	if unitOfWork.tx.Error != nil {
		return &dtos.Result[bool]{
			Result: false,
			Errors: []error{unitOfWork.tx.Error},
		}
	}

	return &dtos.Result[bool]{Result: true}
}

func (unitOfWork *UnitOfWork) CommitTransaction(ctx context.Context) *dtos.Result[bool] {

	commitResult := unitOfWork.tx.Commit().WithContext(ctx)

	if commitResult.Error != nil {
		return &dtos.Result[bool]{
			Result: false,
			Errors: []error{commitResult.Error},
		}
	}

	return &dtos.Result[bool]{Result: true}

}

func (unitOfWork *UnitOfWork) RollBackTransaction(ctx context.Context) *dtos.Result[bool] {

	rollBackResult := unitOfWork.tx.Rollback().WithContext(ctx)

	if rollBackResult.Error != nil {
		return &dtos.Result[bool]{
			Result: false,
			Errors: []error{rollBackResult.Error},
		}
	}

	return &dtos.Result[bool]{Result: true}

}

func (unitOfWork *UnitOfWork) RoleRepository() RoleContract.IRoleRepository {

	if unitOfWork.tx == nil {
		return NewRoleRepository(unitOfWork.db)
	} else {
		return NewRoleRepository(unitOfWork.tx)
	}

}

func (unitOfWork *UnitOfWork) UserRepository() UserContract.IUserRepository {

	if unitOfWork.tx == nil {
		return NewUserRepository(unitOfWork.db)
	} else {
		return NewUserRepository(unitOfWork.tx)
	}

}

func (unitOfWork *UnitOfWork) RoleUserRepository() RoleUserContract.IRoleUserRepository {

	if unitOfWork.tx == nil {
		return NewRoleUserRepository(unitOfWork.db)
	} else {
		return NewRoleUserRepository(unitOfWork.tx)
	}

}

func (unitOfWork *UnitOfWork) EventRepository() CommonContract.IEventRepository {

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
		return &UnitOfWork{db: db}, nil
	}

	return nil, err
}
