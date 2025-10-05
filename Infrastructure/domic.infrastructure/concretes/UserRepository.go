package concretes

import (
	"context"
	"domic.domain/commons/dtos"
	"domic.domain/user/entities"
	"domic.persistence/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (repository *UserRepository) Add(entity *entities.User, context context.Context) *dtos.Result[bool] {

	select {
	case <-context.Done():
		return &dtos.Result[bool]{Errors: []error{context.Err()}, Result: false}
	default:
	}

	dataModel := models.ConvertUserEntityToModel(entity)

	queryResult := repository.db.Model(&models.UserModel{}).Create(dataModel).WithContext(context)

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

func (repository *UserRepository) AddRange(entities []*entities.User, context context.Context) *dtos.Result[bool] {

	select {
	case <-context.Done():
		return &dtos.Result[bool]{Errors: []error{context.Err()}, Result: false}
	default:
	}

	dataModels := models.ConvertUserEntitiesToModels(entities)

	queryResult := repository.db.Model(&models.UserModel{}).CreateInBatches(dataModels, len(entities)).WithContext(context)

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

func (repository *UserRepository) Change(entity *entities.User, context context.Context) *dtos.Result[bool] {

	select {
	case <-context.Done():
		return &dtos.Result[bool]{Errors: []error{context.Err()}, Result: false}
	default:
	}

	dataModel := models.ConvertUserEntityToModel(entity)

	queryResult := repository.db.Model(&models.UserModel{}).Updates(dataModel).WithContext(context)

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

func (repository *UserRepository) ChangeRange(entities []*entities.User, context context.Context) *dtos.Result[bool] {

	select {
	case <-context.Done():
		return &dtos.Result[bool]{Errors: []error{context.Err()}, Result: false}
	default:
	}

	var errors []error

	for model := range models.ConvertUserEntitiesToModels(entities) {

		queryResult := repository.db.Model(&models.UserModel{}).Updates(model).WithContext(context)

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

func (repository *UserRepository) Remove(entity *entities.User, context context.Context) *dtos.Result[bool] {

	select {
	case <-context.Done():
		return &dtos.Result[bool]{Errors: []error{context.Err()}, Result: false}
	default:
	}

	dataModel := models.ConvertUserEntityToModel(entity)

	queryResult := repository.db.Model(&models.UserModel{}).Delete(dataModel, dataModel.Id).WithContext(context)

	if queryResult.Error != nil {
		return &dtos.Result[bool]{
			Result: false,
			Errors: []error{queryResult.Error},
		}
	}

	return &dtos.Result[bool]{Result: true}

}

func (repository *UserRepository) RemoveRange(entities []*entities.User, context context.Context) *dtos.Result[bool] {

	var errors []error

	for _, model := range models.ConvertUserEntitiesToModels(entities) {

		queryResult := repository.db.Model(&models.UserModel{}).Delete(model, model.Id).WithContext(context)

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

func (repository *UserRepository) FindById(id string, context context.Context) *dtos.Result[*entities.User] {

	var model *models.UserModel

	queryResult := repository.db.First(model, "id = ?", id).WithContext(context)

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

func (repository *UserRepository) FindAll(paginationRequest *dtos.PaginationRequest, context context.Context) *dtos.Result[*dtos.PaginationResponse[*entities.User]] {

	offset := (paginationRequest.PageIndex - 1) * paginationRequest.PageSize

	var total int64

	var dataModels []models.UserModel

	countOfItem := repository.db.Model(&models.UserModel{}).Count(&total).WithContext(context)

	if countOfItem.Error != nil {
		return &dtos.Result[*dtos.PaginationResponse[*entities.User]]{
			Result: nil,
			Errors: []error{countOfItem.Error},
		}
	}

	totalPages := int(total / int64(paginationRequest.PageSize))

	queryResult := repository.db.Model(&models.UserModel{}).Limit(paginationRequest.PageSize).Offset(offset).Find(&dataModels).WithContext(context)

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

func (repository *UserRepository) IsExistById(id string, context context.Context) *dtos.Result[bool] {

	var model *models.UserModel

	queryResult := repository.db.First(model, "GetId = ?", id).WithContext(context)

	if queryResult.Error != nil || model == nil {
		return &dtos.Result[bool]{
			Result: false,
			Errors: []error{queryResult.Error},
		}
	}

	return &dtos.Result[bool]{Result: true}

}

func (repository *UserRepository) IsExistByUsername(username string, context context.Context) *dtos.Result[bool] {

	var model *models.UserModel

	queryResult := repository.db.First(model, "Username = ?", username).WithContext(context)

	if queryResult.Error != nil || model == nil {
		return &dtos.Result[bool]{
			Result: false,
			Errors: []error{queryResult.Error},
		}
	}

	return &dtos.Result[bool]{Result: true}

}

func (repository *UserRepository) IsExistByPhoneNumber(phoneNumber string, context context.Context) *dtos.Result[bool] {

	var model *models.UserModel

	queryResult := repository.db.First(model, "PhoneNumber = ?", phoneNumber).WithContext(context)

	if queryResult.Error != nil || model == nil {
		return &dtos.Result[bool]{
			Result: false,
			Errors: []error{queryResult.Error},
		}
	}

	return &dtos.Result[bool]{Result: true}

}

func (repository *UserRepository) IsExistByEmail(email string, context context.Context) *dtos.Result[bool] {

	var model *models.UserModel

	queryResult := repository.db.First(model, "Email = ?", email).WithContext(context)

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
