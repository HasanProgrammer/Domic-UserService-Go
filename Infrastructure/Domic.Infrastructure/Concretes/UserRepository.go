package InfrastructureConcrete

import (
	"Domic.Domain/Commons/DTOs"
	"Domic.Domain/User/Entities"
	"Domic.Persistence/Models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (userRepository *UserRepository) Add(entity *DomainUserEntity.User, result chan DomainCommonDTO.Result[bool]) {

	model := InfrastructureModel.UserModel{
		Id:        entity.GetId(),
		FirstName: entity.GetFirstName(),
		LastName:  entity.GetLastName(),
		Email:     entity.GetEmail(),
		Password:  entity.GetPassword(),
		Username:  entity.GetUsername(),
	}

	queryResult := userRepository.db.Create(&model)

	result <- DomainCommonDTO.Result[bool]{
		Error:  queryResult.Error,
		Result: queryResult.Error == nil,
	}

}

func (userRepository *UserRepository) AddRange(entities []*DomainUserEntity.User, result chan DomainCommonDTO.Result[bool]) {

	queryResult := userRepository.db.CreateInBatches(entities, len(entities))

	result <- DomainCommonDTO.Result[bool]{
		Error:  queryResult.Error,
		Result: queryResult.Error != nil,
	}

}

func (userRepository *UserRepository) Change(entity *DomainUserEntity.User, result chan DomainCommonDTO.Result[bool]) {

	//todo

}

func (userRepository *UserRepository) Remove(entity *DomainUserEntity.User, result chan DomainCommonDTO.Result[bool]) {

	//todo

}

func (userRepository *UserRepository) FindById(id string, result chan DomainCommonDTO.Result[*DomainUserEntity.User]) {

	var user *DomainUserEntity.User

	queryResult := userRepository.db.First(user, "id = ?", id)

	result <- DomainCommonDTO.Result[*DomainUserEntity.User]{
		Error:  queryResult.Error,
		Result: user,
	}

}

func (userRepository *UserRepository) FindAll(paginationRequest *DomainCommonDTO.PaginationRequest, result chan DomainCommonDTO.PaginationResponse[*DomainUserEntity.User]) {

	//todo

}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}
