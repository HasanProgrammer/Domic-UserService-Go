package InfrastructureConcretesRepository

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

	result := userRepository.db.Delete(entity, entity.GetId())

	return result.Error

}

func (userRepository *UserRepository) RemoveAsync(entity *DomainUserEntity.User, result chan error) {

	//todo

}

func (userRepository *UserRepository) FindById(id *string) (*DomainUserEntity.User, error) {

	userRepository.db.Where()

	return nil, nil
}

func (userRepository *UserRepository) FindByIdAsync(id *string, result chan DomainCommonDTO.Result[*DomainUserEntity.User]) {

}

func (userRepository *UserRepository) FindAll() ([]*DomainUserEntity.User, error) {
	return nil, nil
}

func (userRepository *UserRepository) FindAllAsync(result chan DomainCommonDTO.Result[[]*DomainUserEntity.User]) {

}
