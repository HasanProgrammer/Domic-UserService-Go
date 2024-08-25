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

func (userRepository *UserRepository) Add(entity *DomainUserEntity.User[string], result chan DomainCommonDTO.Result[bool]) {

	queryChannel := make(chan DomainCommonDTO.Result[bool])

	go func() {

		model := InfrastructureModel.UserModel{
			Id:        entity.Id(),
			FirstName: entity.FirstName(),
			LastName:  entity.LastName(),
			Email:     entity.Email(),
			Password:  entity.Password(),
			Username:  entity.Username(),
		}

		queryResult := userRepository.db.Create(model)

		queryChannel <- DomainCommonDTO.Result[bool]{
			Error:  queryResult.Error,
			OutPut: queryResult.Error != nil,
		}

	}()

	result <- <-queryChannel

}

func (userRepository *UserRepository) AddRange(entities []*DomainUserEntity.User[string], result chan DomainCommonDTO.Result[bool]) {

	queryChannel := make(chan DomainCommonDTO.Result[bool])

	go func() {

		queryResult := userRepository.db.CreateInBatches(entities, len(entities))

		queryChannel <- DomainCommonDTO.Result[bool]{
			Error:  queryResult.Error,
			OutPut: queryResult.Error != nil,
		}

	}()

	result <- <-queryChannel

}

func (userRepository *UserRepository) Change(entity *DomainUserEntity.User[string], result chan DomainCommonDTO.Result[bool]) {

	//todo

}

func (userRepository *UserRepository) Remove(entity *DomainUserEntity.User[string], result chan DomainCommonDTO.Result[bool]) {

	//todo

}

func (userRepository *UserRepository) FindById(id string, result chan DomainCommonDTO.Result[*DomainUserEntity.User[string]]) {

	queryChannel := make(chan DomainCommonDTO.Result[*DomainUserEntity.User[string]])

	go func() {

		var user *DomainUserEntity.User[string]

		queryResult := userRepository.db.First(user, "id = ?", id)

		queryChannel <- DomainCommonDTO.Result[*DomainUserEntity.User[string]]{
			Error:  queryResult.Error,
			OutPut: user,
		}

	}()

	result <- <-queryChannel

}

func (userRepository *UserRepository) FindAll(paginationRequest *DomainCommonDTO.PaginationRequest, result chan DomainCommonDTO.PaginationResponse[*DomainUserEntity.User[string]]) {

}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}
