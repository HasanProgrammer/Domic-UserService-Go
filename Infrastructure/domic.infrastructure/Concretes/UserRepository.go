package InfrastructureConcrete

import (
	"domic.domain/Commons/DTOs"
	"domic.domain/User/Entities"
	"domic.persistence/Models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (repository *UserRepository) Add(entity *Entities.User) *DTOs.Result[bool] {

	userModel := Models.MapUserEntityToModel(entity)

	queryResult := repository.db.Create(userModel)

	if queryResult.Error != nil {
		return &DTOs.Result[bool]{
			Result: false,
			Errors: []error{queryResult.Error},
		}
	}

	return &DTOs.Result[bool]{
		Result: true,
	}

}

func (repository *UserRepository) AddRange(entities []*Entities.User) *DTOs.Result[bool] {

	queryResult := repository.db.CreateInBatches(entities, len(entities))

	if queryResult.Error != nil {
		return &DTOs.Result[bool]{
			Result: false,
			Errors: []error{queryResult.Error},
		}
	}

	return &DTOs.Result[bool]{
		Result: false,
	}

}

func (repository *UserRepository) Change(entity *Entities.User) *DTOs.Result[bool] {

	queryResult := repository.db.Updates(entity)

	if queryResult.Error != nil {
		return &DTOs.Result[bool]{
			Result: false,
			Errors: []error{queryResult.Error},
		}
	}

	return &DTOs.Result[bool]{
		Result: false,
	}

}

func (repository *UserRepository) ChangeRange(entities []*Entities.User) *DTOs.Result[bool] {

	var errors []error

	for user := range entities {

		queryResult := repository.db.Updates(user)

		if queryResult.Error != nil {
			errors = append(errors, queryResult.Error)
		}

	}

	if len(errors) > 0 {
		return &DTOs.Result[bool]{
			Result: false,
			Errors: errors,
		}
	}

	return &DTOs.Result[bool]{Result: true}

}

func (repository *UserRepository) Remove(entity *Entities.User) *DTOs.Result[bool] {

	queryResult := repository.db.Updates(user)

	if queryResult.Error != nil {
		errors = append(errors, queryResult.Error)
	}

}

func (repository *UserRepository) RemoveRange(entities []*Entities.User) *DTOs.Result[bool] {

	var errors []error

	for _, user := range entities {

		userModel := Models.MapUserEntityToModel(user)

		queryResult := repository.db.Delete(userModel, userModel.Id)

		if queryResult.Error != nil {
			errors = append(errors, queryResult.Error)
		}

	}

	if len(errors) > 0 {
		return &DTOs.Result[bool]{
			Result: false,
			Errors: errors,
		}
	}

	return &DTOs.Result[bool]{
		Result: true,
	}

}

func (repository *UserRepository) FindById(id string) *DTOs.Result[*Entities.User] {

	var userModel *Models.UserModel

	queryResult := repository.db.First(userModel, "id = ?", id)

	if queryResult.Error != nil {
		return &DTOs.Result[*Entities.User]{
			Result: nil,
			Errors: []error{queryResult.Error},
		}
	}

	return &DTOs.Result[*Entities.User]{
		Result: Models.MapUserModelToEntity(userModel),
	}

}

func (repository *UserRepository) FindAll(paginationRequest *DTOs.PaginationRequest) *DTOs.Result[*DTOs.PaginationResponse[*Entities.User]] {

	offset := (paginationRequest.PageIndex - 1) * paginationRequest.PageSize

	var total int64

	var userModels []Models.UserModel

	countOfItem := repository.db.Count(&total)

	if countOfItem.Error != nil {
		return &DTOs.Result[*DTOs.PaginationResponse[*Entities.User]]{
			Result: nil,
			Errors: []error{countOfItem.Error},
		}
	}

	totalPages := int(total / int64(paginationRequest.PageSize))

	queryResult := repository.db.Limit(paginationRequest.PageSize).Offset(offset).Find(&userModels)

	if queryResult.Error != nil {
		return &DTOs.Result[*DTOs.PaginationResponse[*Entities.User]]{
			Result: nil,
			Errors: []error{queryResult.Error},
		}
	}

	return &DTOs.Result[*DTOs.PaginationResponse[*Entities.User]]{
		Result: &DTOs.PaginationResponse[*Entities.User]{
			PageSize:  paginationRequest.PageSize,
			PageIndex: paginationRequest.PageIndex,
			Items:     Models.MapUserModelsToEntity(userModels),
			TotalItem: total,
			HasNext:   paginationRequest.PageIndex < totalPages,
			HasPrev:   paginationRequest.PageIndex > 1,
		},
	}

}

func (repository *UserRepository) IsExistById(username string) *DTOs.Result[bool] {

	return nil

}

func (repository *UserRepository) IsExistByUsername(username string) *DTOs.Result[bool] {

	return nil

}

func (repository *UserRepository) IsExistByPhoneNumber(username string) *DTOs.Result[bool] {

	return nil

}

func (repository *UserRepository) IsExistByEmail(username string) *DTOs.Result[bool] {

	return nil

}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}
