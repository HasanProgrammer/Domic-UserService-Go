package Models

import (
	"domic.domain/Commons/Entities"
	"time"
)

type EventModel struct {
	Id        string    `gorm:"column:Id"`
	Name      string    `gorm:"column:Name"`
	Service   string    `gorm:"column:Service"`
	Table     string    `gorm:"column:Table"`
	Action    string    `gorm:"column:Action"`
	Payload   string    `gorm:"column:Payload"`
	CreatedAt time.Time `gorm:"column:CreatedAt"`
	UpdatedAt *string   `gorm:"column:UpdatedAt"`
	IsActive  bool      `gorm:"column:IsActive"`
}

func MapEventEntityToModel(event *Entities.Event) *EventModel {

	return &EventModel{}

}

func MapEventEntitiesToModel(events []*Entities.Event) []*EventModel {

	var models []*EventModel

	for _, event := range events {

		model := &EventModel{}

		models = append(models, model)

	}

	return models

}

func MapEventModelToEntity(model *EventModel) *Entities.Event {
	return Entities.Assemble(model.Id, model.FirstName, model.LastName, model.Username, model.Password,
		model.Email, model.CreatedBy, model.CreatedRole, model.CreatedAt, model.UpdatedBy, model.UpdatedRole, model.UpdatedAt,
	)
}

func MapUserModelsToEntity(models []UserModel) []*Entities.Event {

	var users []*Entities.Event

	for _, model := range models {

		userEntity := Entities.Assemble(model.Id, model.FirstName, model.LastName, model.Username, model.Password,
			model.Email, model.CreatedBy, model.CreatedRole, model.CreatedAt, model.UpdatedBy, model.UpdatedRole,
			model.UpdatedAt,
		)

		users = append(users, userEntity)

	}

	return users

}
