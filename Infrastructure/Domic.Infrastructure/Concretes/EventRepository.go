package InfrastructureConcrete

import (
	"Domic.Domain/Commons/DTOs"
	"Domic.Domain/Commons/Entities"
	InfrastructureModel "Domic.Persistence/Models"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func (eventRepository *EventRepository) Add(entity *DomainCommonEntity.Event[string], result chan DomainCommonDTO.Result[bool]) {

	queryChannel := make(chan DomainCommonDTO.Result[bool])

	go func() {

		queryResult := eventRepository.db.Create(entity)

		queryChannel <- DomainCommonDTO.Result[bool]{
			Error:  queryResult.Error,
			OutPut: queryResult.Error != nil,
		}

	}()

	result <- <-queryChannel

}

func (eventRepository *EventRepository) AddRange(entities []*DomainCommonEntity.Event[string], result chan DomainCommonDTO.Result[bool]) {

	queryChannel := make(chan DomainCommonDTO.Result[bool])

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

		queryChannel <- DomainCommonDTO.Result[bool]{
			Error:  queryResult.Error,
			OutPut: queryResult.Error != nil,
		}

	}()

	result <- <-queryChannel

}

func (eventRepository *EventRepository) Change(entity *DomainCommonEntity.Event[string], result chan DomainCommonDTO.Result[bool]) {

	//todo

}

func (eventRepository *EventRepository) Remove(entity *DomainCommonEntity.Event[string], result chan DomainCommonDTO.Result[bool]) {

	//todo

}

func (eventRepository *EventRepository) FindById(id string, result chan DomainCommonDTO.Result[*DomainCommonEntity.Event[string]]) {

	queryChannel := make(chan DomainCommonDTO.Result[*DomainCommonEntity.Event[string]])

	go func() {

		var user *DomainCommonEntity.Event[string]

		queryResult := eventRepository.db.First(user, "id = ?", id)

		queryChannel <- DomainCommonDTO.Result[*DomainCommonEntity.Event[string]]{
			Error:  queryResult.Error,
			OutPut: user,
		}

	}()

	result <- <-queryChannel

}

func (eventRepository *EventRepository) FindAll(paginationRequest *DomainCommonDTO.PaginationRequest, result chan DomainCommonDTO.PaginationResponse[*DomainCommonEntity.Event[string]]) {

}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{db: db}
}
