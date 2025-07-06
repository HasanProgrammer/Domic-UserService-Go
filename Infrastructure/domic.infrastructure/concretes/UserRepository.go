package concretes

import (
	"domic.domain/commons/dtos"
	"domic.domain/user/entities"
	"domic.persistence/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (repository *UserRepository) Add(entity *entities.User) *dtos.Result[bool] {

	dataModel := models.ConvertUserEntityToModel(entity)

	queryResult := repository.db.Model(&models.UserModel{}).Create(dataModel)

	if queryResult.Error != nil {
		return &dtos.Result[bool]{
			Result: false,
			Errors: []error{queryResult.Error},
		}
	}

	return &dtos.Result[bool]{
		Result: true,
	}

}

func (repository *UserRepository) AddRange(entities []*entities.User) *dtos.Result[bool] {

	dataModels := models.ConvertUserEntitiesToModels(entities)

	queryResult := repository.db.Model(&models.UserModel{}).CreateInBatches(dataModels, len(entities))

	if queryResult.Error != nil {
		return &dtos.Result[bool]{
			Result: false,
			Errors: []error{queryResult.Error},
		}
	}

	return &dtos.Result[bool]{
		Result: false,
	}

}

func (repository *UserRepository) Change(entity *entities.User) *dtos.Result[bool] {

	dataModel := models.ConvertUserEntityToModel(entity)

	queryResult := repository.db.Model(&models.UserModel{}).Updates(dataModel)

	if queryResult.Error != nil {
		return &dtos.Result[bool]{
			Result: false,
			Errors: []error{queryResult.Error},
		}
	}

	return &dtos.Result[bool]{
		Result: false,
	}

}

func (repository *UserRepository) ChangeRange(entities []*entities.User) *dtos.Result[bool] {

	var errors []error

	for model := range models.ConvertUserEntitiesToModels(entities) {

		queryResult := repository.db.Model(&models.UserModel{}).Updates(model)

		if queryResult.Error != nil {
			errors = append(errors, queryResult.Error)
		}

	}

	if len(errors) > 0 {
		return &dtos.Result[bool]{
			Result: false,
			Errors: errors,
		}
	}

	return &dtos.Result[bool]{Result: true}

}

func (repository *UserRepository) Remove(entity *entities.User) *dtos.Result[bool] {

	dataModel := models.ConvertUserEntityToModel(entity)

	queryResult := repository.db.Model(&models.UserModel{}).Delete(dataModel, dataModel.Id)

	if queryResult.Error != nil {
		return &dtos.Result[bool]{
			Result: false,
			Errors: []error{queryResult.Error},
		}
	}

	return &dtos.Result[bool]{Result: true}

}

func (repository *UserRepository) RemoveRange(entities []*entities.User) *dtos.Result[bool] {

	var errors []error

	for _, model := range models.ConvertUserEntitiesToModels(entities) {

		queryResult := repository.db.Model(&models.UserModel{}).Delete(model, model.Id)

		if queryResult.Error != nil {
			errors = append(errors, queryResult.Error)
		}

	}

	if len(errors) > 0 {
		return &dtos.Result[bool]{
			Result: false,
			Errors: errors,
		}
	}

	return &dtos.Result[bool]{
		Result: true,
	}

}

func (repository *UserRepository) FindById(id string) *dtos.Result[*entities.User] {

	var model *models.UserModel

	queryResult := repository.db.First(model, "id = ?", id)

	if queryResult.Error != nil {
		return &dtos.Result[*entities.User]{
			Result: nil,
			Errors: []error{queryResult.Error},
		}
	}

	return &dtos.Result[*entities.User]{
		Result: models.ConvertUserModelToEntity(model),
	}

}

func (repository *UserRepository) FindAll(paginationRequest *dtos.PaginationRequest) *dtos.Result[*dtos.PaginationResponse[*entities.User]] {

	offset := (paginationRequest.PageIndex - 1) * paginationRequest.PageSize

	var total int64

	var dataModels []models.UserModel

	countOfItem := repository.db.Model(&models.UserModel{}).Count(&total)

	if countOfItem.Error != nil {
		return &dtos.Result[*dtos.PaginationResponse[*entities.User]]{
			Result: nil,
			Errors: []error{countOfItem.Error},
		}
	}

	totalPages := int(total / int64(paginationRequest.PageSize))

	queryResult := repository.db.Model(&models.UserModel{}).Limit(paginationRequest.PageSize).Offset(offset).Find(&dataModels)

	if queryResult.Error != nil {
		return &dtos.Result[*dtos.PaginationResponse[*entities.User]]{
			Result: nil,
			Errors: []error{queryResult.Error},
		}
	}

	return &dtos.Result[*dtos.PaginationResponse[*entities.User]]{
		Result: &dtos.PaginationResponse[*entities.User]{
			PageSize:  paginationRequest.PageSize,
			PageIndex: paginationRequest.PageIndex,
			Items:     models.ConvertUserModelsToEntities(dataModels),
			TotalItem: total,
			HasNext:   paginationRequest.PageIndex < totalPages,
			HasPrev:   paginationRequest.PageIndex > 1,
		},
	}

}

func (repository *UserRepository) IsExistById(id string) *dtos.Result[bool] {

	var model *models.UserModel

	queryResult := repository.db.First(model, "Id = ?", id)

	if queryResult.Error != nil || model == nil {
		return &dtos.Result[bool]{
			Result: false,
			Errors: []error{queryResult.Error},
		}
	}

	return &dtos.Result[bool]{Result: true}

}

func (repository *UserRepository) IsExistByUsername(username string) *dtos.Result[bool] {

	var model *models.UserModel

	queryResult := repository.db.First(model, "Username = ?", username)

	if queryResult.Error != nil || model == nil {
		return &dtos.Result[bool]{
			Result: false,
			Errors: []error{queryResult.Error},
		}
	}

	return &dtos.Result[bool]{Result: true}

}

func (repository *UserRepository) IsExistByPhoneNumber(phoneNumber string) *dtos.Result[bool] {

	var model *models.UserModel

	queryResult := repository.db.First(model, "PhoneNumber = ?", phoneNumber)

	if queryResult.Error != nil || model == nil {
		return &dtos.Result[bool]{
			Result: false,
			Errors: []error{queryResult.Error},
		}
	}

	return &dtos.Result[bool]{Result: true}

}

func (repository *UserRepository) IsExistByEmail(email string) *dtos.Result[bool] {

	var model *models.UserModel

	queryResult := repository.db.First(model, "Email = ?", email)

	if queryResult.Error != nil || model == nil {
		return &dtos.Result[bool]{
			Result: false,
			Errors: []error{queryResult.Error},
		}
	}

	return &dtos.Result[bool]{Result: true}

}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}
