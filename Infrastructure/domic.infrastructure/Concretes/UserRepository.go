package InfrastructureConcrete

import (
	"domic.domain/Commons/DTOs"
	"domic.domain/User/Entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (userRepository *UserRepository) Add(entity *Entities.User[string]) *DTOs.Results[bool] {

	queryResult := userRepository.db.Create(entity)

	if queryResult.Error != nil {
		return &DTOs.Results[bool]{
			Result: false,
			Errors: []error{queryResult.Error},
		}
	}

	return &DTOs.Results[bool]{
		Result: false,
	}

}

func (userRepository *UserRepository) AddRange(entities []*Entities.User[string]) *DTOs.Results[bool] {

	queryResult := userRepository.db.CreateInBatches(entities, len(entities))

	if queryResult.Error != nil {
		return &DTOs.Results[bool]{
			Result: false,
			Errors: []error{queryResult.Error},
		}
	}

	return &DTOs.Results[bool]{
		Result: false,
	}

}

func (userRepository *UserRepository) Change(entity *Entities.User[string]) *DTOs.Results[bool] {

	//todo

}

func (userRepository *UserRepository) ChangeRange(entities []*Entities.User[string]) *DTOs.Results[bool] {

	//todo

}

func (userRepository *UserRepository) Remove(entity *Entities.User[string]) *DTOs.Results[bool] {

	//todo

}

func (userRepository *UserRepository) RemoveRange(entities []*Entities.User[string]) *DTOs.Results[bool] {

	//todo

}

func (userRepository *UserRepository) FindById(id string) *DTOs.Results[bool] {

	queryChannel := make(chan DTOs.Results[*Entities.User[string]])

	var user *Entities.User[string]

	queryResult := userRepository.db.First(user, "id = ?", id)

}

func (userRepository *UserRepository) FindAll(paginationRequest *DTOs.PaginationRequest) *DTOs.Results[DTOs.PaginationResponse[Entities.User[string]]] {

}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}
