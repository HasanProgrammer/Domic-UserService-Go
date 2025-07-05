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

	model := Models.ConvertEventEntityToModel(entity)

	queryResult := repository.db.Model(&Models.EventModel{}).Create(model)

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

	models := Models.ConvertEventEntitiesToModels(entities)

	queryResult := repository.db.Model(&Models.EventModel{}).CreateInBatches(models, len(entities))

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

	model := Models.ConvertEventEntityToModel(entity)

	queryResult := repository.db.Model(&Models.EventModel{}).Updates(model)

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

	models := Models.ConvertEventEntitiesToModels(entities)

	for model := range models {

		queryResult := repository.db.Model(&Models.EventModel{}).Updates(model)

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

	model := Models.ConvertEventEntityToModel(entity)

	queryResult := repository.db.Model(&Models.EventModel{}).Delete(model, model.Id)

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

	models := Models.ConvertEventEntitiesToModels(entities)

	for _, model := range models {

		queryResult := repository.db.Model(&Models.EventModel{}).Delete(model, model.Id)

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

	var model *Models.EventModel

	queryResult := repository.db.First(model, "id = ?", id)

	if queryResult.Error != nil {
		return &DTOs.Result[*Entities.Event]{
			Result: nil,
			Errors: []error{queryResult.Error},
		}
	}

	return &DTOs.Result[*Entities.Event]{
		Result: Models.ConvertEventModelToEntity(model),
	}

}

func (repository *EventRepository) FindAll(paginationRequest *DTOs.PaginationRequest) *DTOs.Result[*DTOs.PaginationResponse[*Entities.Event]] {

	offset := (paginationRequest.PageIndex - 1) * paginationRequest.PageSize

	var total int64

	var models []Models.EventModel

	countOfItem := repository.db.Model(&Models.EventModel{}).Count(&total)

	if countOfItem.Error != nil {
		return &DTOs.Result[*DTOs.PaginationResponse[*Entities.Event]]{
			Result: nil,
			Errors: []error{countOfItem.Error},
		}
	}

	totalPages := int(total / int64(paginationRequest.PageSize))

	queryResult := repository.db.Model(&Models.EventModel{}).Limit(paginationRequest.PageSize).Offset(offset).Find(&models)

	if queryResult.Error != nil {
		return &DTOs.Result[*DTOs.PaginationResponse[*Entities.Event]]{
			Result: nil,
			Errors: []error{queryResult.Error},
		}
	}

	return &DTOs.Result[*DTOs.PaginationResponse[*Entities.Event]]{
		Result: &DTOs.PaginationResponse[*Entities.Event]{
			PageSize:  paginationRequest.PageSize,
			PageIndex: paginationRequest.PageIndex,
			Items:     Models.ConvertEventModelsToEntities(models),
			TotalItem: total,
			HasNext:   paginationRequest.PageIndex < totalPages,
			HasPrev:   paginationRequest.PageIndex > 1,
		},
	}

}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{db: db}
}
