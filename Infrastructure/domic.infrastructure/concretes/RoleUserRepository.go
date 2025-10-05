package concretes

import (
	"context"
	"domic.domain/commons/dtos"
	"domic.domain/role_user/entities"
	"domic.persistence/models"
	"gorm.io/gorm"
)

type RoleUserRepository struct {
	db *gorm.DB
}

func (repository *RoleUserRepository) Add(entity *entities.RoleUser, context context.Context) *dtos.Result[bool] {

	dataModel := models.ConvertRoleUserEntityToModel(entity)

	queryResult := repository.db.Model(&models.RoleUserModel{}).Create(dataModel).WithContext(context)

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

func (repository *RoleUserRepository) AddRange(entities []*entities.RoleUser, context context.Context) *dtos.Result[bool] {

	dataModels := models.ConvertRoleUserEntitiesToModels(entities)

	queryResult := repository.db.Model(&models.RoleUserModel{}).CreateInBatches(dataModels, len(entities)).WithContext(context)

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

func (repository *RoleUserRepository) Change(entity *entities.RoleUser, context context.Context) *dtos.Result[bool] {

	dataModel := models.ConvertRoleUserEntityToModel(entity)

	queryResult := repository.db.Model(&models.RoleUserModel{}).Updates(dataModel).WithContext(context)

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

func (repository *RoleUserRepository) ChangeRange(entities []*entities.RoleUser, context context.Context) *dtos.Result[bool] {

	var errors []error

	for model := range models.ConvertRoleUserEntitiesToModels(entities) {

		queryResult := repository.db.Model(&models.RoleUserModel{}).Updates(model).WithContext(context)

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

func (repository *RoleUserRepository) Remove(entity *entities.RoleUser, context context.Context) *dtos.Result[bool] {

	dataModel := models.ConvertRoleUserEntityToModel(entity)

	queryResult := repository.db.Model(&models.RoleUserModel{}).Delete(dataModel, dataModel.Id).WithContext(context)

	if queryResult.Error != nil {
		return &dtos.Result[bool]{
			Result: false,
			Errors: []error{queryResult.Error},
		}
	}

	return &dtos.Result[bool]{Result: true}

}

func (repository *RoleUserRepository) RemoveRange(entities []*entities.RoleUser, context context.Context) *dtos.Result[bool] {

	var errors []error

	for _, model := range models.ConvertRoleUserEntitiesToModels(entities) {

		queryResult := repository.db.Model(&models.RoleUserModel{}).Delete(model, model.Id).WithContext(context)

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

func (repository *RoleUserRepository) FindById(id string, context context.Context) *dtos.Result[*entities.RoleUser] {

	return nil

}

func (repository *RoleUserRepository) FindAll(paginationRequest *dtos.PaginationRequest, context context.Context) *dtos.Result[*dtos.PaginationResponse[*entities.RoleUser]] {

	return nil

}

func NewRoleUserRepository(db *gorm.DB) *RoleUserRepository {
	return &RoleUserRepository{db: db}
}
