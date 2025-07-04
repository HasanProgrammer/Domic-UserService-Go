package Concrete

import (
	"domic.domain/Commons/DTOs"
	"domic.domain/Commons/Entities"
	"domic.persistence/Models"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func (repository *EventRepository) Add(entity *Entities.Event) *DTOs.Result[bool] {

	userModel := Models.MapUserEntityToModel(entity)

	queryResult := repository.db.Model(&Models.UserModel{}).Create(userModel)

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

func (repository *EventRepository) AddRange(entities []*Entities.Event) *DTOs.Result[bool] {

	models := Models.MapUserEntitiesToModel(entities)

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

func (repository *EventRepository) Change(entity *Entities.Event) *DTOs.Result[bool] {

	model := Models.MapUserEntityToModel(entity)

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

func (repository *EventRepository) ChangeRange(entities []*Entities.Event) *DTOs.Result[bool] {

	var errors []error

	models := Models.MapUserEntitiesToModel(entities)

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

func (repository *EventRepository) Remove(entity *Entities.Event) *DTOs.Result[bool] {

	model := Models.MapUserEntityToModel(entity)

	queryResult := repository.db.Model(&Models.UserModel{}).Delete(model, model.Id)

	if queryResult.Error != nil {
		return &DTOs.Result[bool]{
			Result: false,
			Errors: []error{queryResult.Error},
		}
	}

	return &DTOs.Result[bool]{Result: true}

}

func (repository *EventRepository) RemoveRange(entities []*Entities.Event) *DTOs.Result[bool] {

	var errors []error

	models := Models.MapUserEntitiesToModel(entities)

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

func (repository *EventRepository) FindById(id string) *DTOs.Result[*Entities.Event] {

	var userModel *Models.UserModel

	queryResult := repository.db.First(userModel, "id = ?", id)

	if queryResult.Error != nil {
		return &DTOs.Result[*Entities.Event]{
			Result: nil,
			Errors: []error{queryResult.Error},
		}
	}

	return &DTOs.Result[*Entities.Event]{
		Result: Models.MapUserModelToEntity(userModel),
	}

}

func (repository *EventRepository) FindAll(paginationRequest *DTOs.PaginationRequest) *DTOs.Result[*DTOs.PaginationResponse[*Entities.Event]] {

}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{db: db}
}
