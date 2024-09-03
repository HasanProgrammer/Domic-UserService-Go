package InfrastructureConcrete

import (
	"Domic.Domain/Commons/DTOs"
	"Domic.Domain/Commons/Entities"
	"Domic.Persistence/Models"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func (eventRepository *EventRepository) Add(entity *DomainCommonEntity.Event, result chan DomainCommonDTO.Result[bool]) {

	go func() {

		queryResult := eventRepository.db.Create(entity)

		result <- DomainCommonDTO.Result[bool]{
			Error:  queryResult.Error,
			Result: queryResult.Error != nil,
		}

	}()

}

func (eventRepository *EventRepository) AddRange(entities []*DomainCommonEntity.Event, result chan DomainCommonDTO.Result[bool]) {

	go func() {

		var models []*InfrastructureModel.EventModel

		for _, entity := range entities {
			models = append(models, &InfrastructureModel.EventModel{
				Id:          entity.GetId(),
				Name:        entity.GetName(),
				Table:       entity.GetTable(),
				Action:      entity.GetAction(),
				Payload:     entity.GetPayload(),
				CreatedAt:   entity.GetCreatedAt(),
				CreatedBy:   entity.GetCreatedBy(),
				CreatedRole: entity.GetCreatedRole(),
			})
		}

		queryResult := eventRepository.db.CreateInBatches(models, len(entities))

		result <- DomainCommonDTO.Result[bool]{
			Error:  queryResult.Error,
			Result: queryResult.Error != nil,
		}

	}()

}

func (eventRepository *EventRepository) Change(entity *DomainCommonEntity.Event, result chan DomainCommonDTO.Result[bool]) {

	//todo

}

func (eventRepository *EventRepository) Remove(entity *DomainCommonEntity.Event, result chan DomainCommonDTO.Result[bool]) {

	//todo

}

func (eventRepository *EventRepository) FindById(id string, result chan DomainCommonDTO.Result[*DomainCommonEntity.Event]) {

	go func() {

		var user *DomainCommonEntity.Event

		queryResult := eventRepository.db.First(user, "id = ?", id)

		result <- DomainCommonDTO.Result[*DomainCommonEntity.Event]{
			Error:  queryResult.Error,
			Result: user,
		}

	}()

}

func (eventRepository *EventRepository) FindAll(paginationRequest *DomainCommonDTO.PaginationRequest,
	result chan DomainCommonDTO.PaginationResponse[*DomainCommonEntity.Event],
) {

}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{db: db}
}
