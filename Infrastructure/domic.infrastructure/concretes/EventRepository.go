package concretes

import (
	"domic.domain/commons/dtos"
	"domic.domain/commons/entities"
	"domic.persistence/models"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func (repository *EventRepository) Add(entity *entities.Event) *dtos.Result[bool] {

	dataModel := models.ConvertEventEntityToModel(entity)

	queryResult := repository.db.Model(&models.EventModel{}).Create(dataModel)

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

func (repository *EventRepository) AddRange(entities []*entities.Event) *dtos.Result[bool] {

	dataModels := models.ConvertEventEntitiesToModels(entities)

	queryResult := repository.db.Model(&models.EventModel{}).CreateInBatches(dataModels, len(entities))

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

func (repository *EventRepository) Change(entity *entities.Event) *dtos.Result[bool] {

	dataModel := models.ConvertEventEntityToModel(entity)

	queryResult := repository.db.Model(&models.EventModel{}).Updates(dataModel)

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

func (repository *EventRepository) ChangeRange(entities []*entities.Event) *dtos.Result[bool] {

	var errors []error

	for model := range models.ConvertEventEntitiesToModels(entities) {

		queryResult := repository.db.Model(&models.EventModel{}).Updates(model)

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

func (repository *EventRepository) Remove(entity *entities.Event) *dtos.Result[bool] {

	dataModel := models.ConvertEventEntityToModel(entity)

	queryResult := repository.db.Model(&models.EventModel{}).Delete(dataModel, dataModel.Id)

	if queryResult.Error != nil {
		return &dtos.Result[bool]{
			Result: false,
			Errors: []error{queryResult.Error},
		}
	}

	return &dtos.Result[bool]{Result: true}

}

func (repository *EventRepository) RemoveRange(entities []*entities.Event) *dtos.Result[bool] {

	var errors []error

	for _, model := range models.ConvertEventEntitiesToModels(entities) {

		queryResult := repository.db.Model(&models.EventModel{}).Delete(model, model.Id)

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

func (repository *EventRepository) FindById(id string) *dtos.Result[*entities.Event] {

	var model *models.EventModel

	queryResult := repository.db.First(model, "id = ?", id)

	if queryResult.Error != nil {
		return &dtos.Result[*entities.Event]{
			Result: nil,
			Errors: []error{queryResult.Error},
		}
	}

	return &dtos.Result[*entities.Event]{
		Result: models.ConvertEventModelToEntity(model),
	}

}

func (repository *EventRepository) FindAll(paginationRequest *dtos.PaginationRequest) *dtos.Result[*dtos.PaginationResponse[*entities.Event]] {

	offset := (paginationRequest.PageIndex - 1) * paginationRequest.PageSize

	var total int64

	var dataModels []models.EventModel

	countOfItem := repository.db.Model(&models.EventModel{}).Count(&total)

	if countOfItem.Error != nil {
		return &dtos.Result[*dtos.PaginationResponse[*entities.Event]]{
			Result: nil,
			Errors: []error{countOfItem.Error},
		}
	}

	totalPages := int(total / int64(paginationRequest.PageSize))

	queryResult := repository.db.Model(&models.EventModel{}).Limit(paginationRequest.PageSize).Offset(offset).Find(&dataModels)

	if queryResult.Error != nil {
		return &dtos.Result[*dtos.PaginationResponse[*entities.Event]]{
			Result: nil,
			Errors: []error{queryResult.Error},
		}
	}

	return &dtos.Result[*dtos.PaginationResponse[*entities.Event]]{
		Result: &dtos.PaginationResponse[*entities.Event]{
			PageSize:  paginationRequest.PageSize,
			PageIndex: paginationRequest.PageIndex,
			Items:     models.ConvertEventModelsToEntities(dataModels),
			TotalItem: total,
			HasNext:   paginationRequest.PageIndex < totalPages,
			HasPrev:   paginationRequest.PageIndex > 1,
		},
	}

}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{db: db}
}
