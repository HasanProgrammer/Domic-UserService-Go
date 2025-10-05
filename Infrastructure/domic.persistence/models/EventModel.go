package models

import (
	"domic.domain/commons/entities"
	"time"
)

type EventModel struct {
	Id        string     `gorm:"column:GetId"`
	Name      string     `gorm:"column:Name"`
	Service   string     `gorm:"column:Service"`
	Table     string     `gorm:"column:Table"`
	Action    string     `gorm:"column:Action"`
	Payload   string     `gorm:"column:Payload"`
	CreatedAt time.Time  `gorm:"column:GetCreatedAt"`
	UpdatedAt *time.Time `gorm:"column:GetUpdatedAt"`
	IsActive  bool       `gorm:"column:GetIsActive"`
}

func ConvertEventEntityToModel(event *entities.Event) *EventModel {

	return &EventModel{
		Id:        event.GetId(),
		Name:      event.GetName(),
		Service:   event.GetService(),
		Table:     event.GetTable(),
		Action:    event.GetAction(),
		Payload:   event.GetPayload(),
		IsActive:  event.GetIsActive(),
		CreatedAt: event.GetCreatedAt(),
		UpdatedAt: event.GetUpdatedAt(),
	}

}

func ConvertEventEntitiesToModels(events []*entities.Event) []*EventModel {

	var models []*EventModel

	for _, event := range events {

		model := &EventModel{
			Id:      event.GetId(),
			Name:    event.GetName(),
			Table:   event.GetTable(),
			Service: event.GetService(),
		}

		models = append(models, model)

	}

	return models

}

func ConvertEventModelToEntity(model *EventModel) *entities.Event {
	return entities.Assemble(model.Id, model.Name, model.Service, model.Table, model.Action, model.Payload,
		model.CreatedAt, model.UpdatedAt, model.IsActive,
	)
}

func ConvertEventModelsToEntities(models []EventModel) []*entities.Event {

	var events []*entities.Event

	for _, model := range models {

		eventEntity := entities.Assemble(model.Id, model.Name, model.Service, model.Table, model.Action, model.Payload,
			model.CreatedAt, model.UpdatedAt, model.IsActive,
		)

		events = append(events, eventEntity)

	}

	return events

}
