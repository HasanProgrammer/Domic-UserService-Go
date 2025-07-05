package Concrete

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

	model := Models.ConvertUserEntityToModel(entity)

	queryResult := repository.db.Model(&Models.UserModel{}).Create(model)

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

	models := Models.ConvertUserEntitiesToModels(entities)

	queryResult := repository.db.Model(&Models.UserModel{}).CreateInBatches(models, len(entities))

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

	model := Models.ConvertUserEntityToModel(entity)

	queryResult := repository.db.Model(&Models.UserModel{}).Updates(model)

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

	models := Models.ConvertUserEntitiesToModels(entities)

	for model := range models {

		queryResult := repository.db.Model(&Models.UserModel{}).Updates(model)

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

	model := Models.ConvertUserEntityToModel(entity)

	queryResult := repository.db.Model(&Models.UserModel{}).Delete(model, model.Id)

	if queryResult.Error != nil {
		return &DTOs.Result[bool]{
			Result: false,
			Errors: []error{queryResult.Error},
		}
	}

	return &DTOs.Result[bool]{Result: true}

}

func (repository *UserRepository) RemoveRange(entities []*Entities.User) *DTOs.Result[bool] {

	var errors []error

	models := Models.ConvertUserEntitiesToModels(entities)

	for _, model := range models {

		queryResult := repository.db.Model(&Models.UserModel{}).Delete(model, model.Id)

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

	var model *Models.UserModel

	queryResult := repository.db.First(model, "id = ?", id)

	if queryResult.Error != nil {
		return &DTOs.Result[*Entities.User]{
			Result: nil,
			Errors: []error{queryResult.Error},
		}
	}

	return &DTOs.Result[*Entities.User]{
		Result: Models.ConvertUserModelToEntity(model),
	}

}

func (repository *UserRepository) FindAll(paginationRequest *DTOs.PaginationRequest) *DTOs.Result[*DTOs.PaginationResponse[*Entities.User]] {

	offset := (paginationRequest.PageIndex - 1) * paginationRequest.PageSize

	var total int64

	var models []Models.UserModel

	countOfItem := repository.db.Model(&Models.UserModel{}).Count(&total)

	if countOfItem.Error != nil {
		return &DTOs.Result[*DTOs.PaginationResponse[*Entities.User]]{
			Result: nil,
			Errors: []error{countOfItem.Error},
		}
	}

	totalPages := int(total / int64(paginationRequest.PageSize))

	queryResult := repository.db.Model(&Models.UserModel{}).Limit(paginationRequest.PageSize).Offset(offset).Find(&models)

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
			Items:     Models.ConvertUserModelsToEntities(models),
			TotalItem: total,
			HasNext:   paginationRequest.PageIndex < totalPages,
			HasPrev:   paginationRequest.PageIndex > 1,
		},
	}

}

func (repository *UserRepository) IsExistById(id string) *DTOs.Result[bool] {

	return nil

}

func (repository *UserRepository) IsExistByUsername(username string) *DTOs.Result[bool] {

	var model *Models.UserModel

	queryResult := repository.db.First(model, "Username = ?", username)

	if queryResult.Error != nil || model == nil {
		return &DTOs.Result[bool]{
			Result: false,
			Errors: []error{queryResult.Error},
		}
	}

	return &DTOs.Result[bool]{Result: true}

}

func (repository *UserRepository) IsExistByPhoneNumber(phoneNumber string) *DTOs.Result[bool] {

	var model *Models.UserModel

	queryResult := repository.db.First(model, "PhoneNumber = ?", phoneNumber)

	if queryResult.Error != nil || model == nil {
		return &DTOs.Result[bool]{
			Result: false,
			Errors: []error{queryResult.Error},
		}
	}

	return &DTOs.Result[bool]{Result: true}

}

func (repository *UserRepository) IsExistByEmail(email string) *DTOs.Result[bool] {

	var model *Models.UserModel

	queryResult := repository.db.First(model, "Email = ?", email)

	if queryResult.Error != nil || model == nil {
		return &DTOs.Result[bool]{
			Result: false,
			Errors: []error{queryResult.Error},
		}
	}

	return &DTOs.Result[bool]{Result: true}

}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}
