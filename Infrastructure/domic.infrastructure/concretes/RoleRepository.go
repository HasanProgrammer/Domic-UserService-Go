package concretes

import (
	"context"
	"domic.domain/commons/dtos"
	"domic.domain/role/entities"
	"domic.persistence/models"
	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

func (repository *RoleRepository) Add(entity *entities.Role, context context.Context) *dtos.Result[bool] {

	dataModel := models.ConvertRoleEntityToModel(entity)

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

func (repository *RoleRepository) AddRange(entities []*entities.Role, context context.Context) *dtos.Result[bool] {

	dataModels := models.ConvertRoleEntitiesToModels(entities)

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

func (repository *RoleRepository) Change(entity *entities.Role, context context.Context) *dtos.Result[bool] {

	dataModel := models.ConvertRoleEntityToModel(entity)

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

func (repository *RoleRepository) ChangeRange(entities []*entities.Role, context context.Context) *dtos.Result[bool] {

	var errors []error

	for model := range models.ConvertRoleEntitiesToModels(entities) {

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

func (repository *RoleRepository) Remove(entity *entities.Role, context context.Context) *dtos.Result[bool] {

	dataModel := models.ConvertRoleEntityToModel(entity)

	queryResult := repository.db.Model(&models.UserModel{}).Delete(dataModel, dataModel.Id).WithContext(context)

	if queryResult.Error != nil {
		return &dtos.Result[bool]{
			Result: false,
			Errors: []error{queryResult.Error},
		}
	}

	return &dtos.Result[bool]{Result: true}

}

func (repository *RoleRepository) RemoveRange(entities []*entities.Role, context context.Context) *dtos.Result[bool] {

	var errors []error

	for _, model := range models.ConvertRoleEntitiesToModels(entities) {

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

func (repository *RoleRepository) FindById(id string, context context.Context) *dtos.Result[*entities.Role] {

	var model *models.RoleModel

	queryResult := repository.db.First(model, "id = ?", id).WithContext(context)

	if queryResult.Error != nil {
		return &dtos.Result[*entities.Role]{
			Result: nil,
			Errors: []error{queryResult.Error},
		}
	}

	return &dtos.Result[*entities.Role]{
		Result: models.ConvertRoleModelToEntity(model),
	}

}

func (repository *RoleRepository) FindAll(paginationRequest *dtos.PaginationRequest, context context.Context) *dtos.Result[*dtos.PaginationResponse[*entities.Role]] {

	offset := (paginationRequest.PageIndex - 1) * paginationRequest.PageSize

	var total int64

	var dataModels []models.RoleModel

	countOfItem := repository.db.Model(&models.UserModel{}).Count(&total).WithContext(context)

	if countOfItem.Error != nil {
		return &dtos.Result[*dtos.PaginationResponse[*entities.Role]]{
			Result: nil,
			Errors: []error{countOfItem.Error},
		}
	}

	totalPages := int(total / int64(paginationRequest.PageSize))

	queryResult := repository.db.Model(&models.UserModel{}).Limit(paginationRequest.PageSize).Offset(offset).Find(&dataModels).WithContext(context)

	if queryResult.Error != nil {
		return &dtos.Result[*dtos.PaginationResponse[*entities.Role]]{
			Result: nil,
			Errors: []error{queryResult.Error},
		}
	}

	return &dtos.Result[*dtos.PaginationResponse[*entities.Role]]{
		Result: &dtos.PaginationResponse[*entities.Role]{
			PageSize:  paginationRequest.PageSize,
			PageIndex: paginationRequest.PageIndex,
			Items:     models.ConvertRoleModelsToEntities(dataModels),
			TotalItem: total,
			HasNext:   paginationRequest.PageIndex < totalPages,
			HasPrev:   paginationRequest.PageIndex > 1,
		},
	}

}

func (repository *RoleRepository) IsExistById(id string, context context.Context) *dtos.Result[bool] {

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

func (repository *RoleRepository) IsExistByUsername(username string, context context.Context) *dtos.Result[bool] {

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

func (repository *RoleRepository) IsExistByPhoneNumber(phoneNumber string, context context.Context) *dtos.Result[bool] {

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

func (repository *RoleRepository) IsExistByEmail(email string, context context.Context) *dtos.Result[bool] {

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

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{db: db}
}
