package InfrastructureConcreteRepository

import (
	"Dotris.Domain/Commons/DTOs"
	"Dotris.Domain/User/Entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (userRepository *UserRepository) Add(entity *DomainUserEntity.User) error {

	result := userRepository.db.Create(entity)

	return result.Error

}

func (userRepository *UserRepository) AddAsync(entity *DomainUserEntity.User, result chan error) {

	//todo

}

func (userRepository *UserRepository) Change(entity *DomainUserEntity.User) error {

	result := userRepository.db.Updates(entity)

	return result.Error

}

func (userRepository *UserRepository) ChangeAsync(entity *DomainUserEntity.User, result chan error) {

	//todo

}

func (userRepository *UserRepository) Remove(entity *DomainUserEntity.User) error {

	result := userRepository.db.Delete(entity, entity.Id())

	return result.Error

}

func (userRepository *UserRepository) RemoveAsync(entity *DomainUserEntity.User, result chan error) {

	//todo

}

func (userRepository *UserRepository) FindById(id *string) (*DomainUserEntity.User, error) {

	var user *DomainUserEntity.User

	result := userRepository.db.First(user, "id = ?", id)

	return user, result.Error
}

func (userRepository *UserRepository) FindByIdAsync(id *string, result chan DomainCommonDTO.Result[*DomainUserEntity.User]) {

	//todo

}

func (userRepository *UserRepository) FindAll(pageSize int64, pageIndex int64) ([]*DomainUserEntity.User, error) {

	return nil, nil

}

func (userRepository *UserRepository) FindAllAsync(pageSize int64, pageIndex int64, result chan DomainCommonDTO.Result[[]*DomainUserEntity.User]) {

}

func (userRepository *UserRepository) Count(conditions ...interface{}) (int64, error) {

	return 0, nil

}
