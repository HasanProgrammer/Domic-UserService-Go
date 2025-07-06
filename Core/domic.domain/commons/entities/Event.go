package entities

import (
	"domic.domain/commons/contracts/interfaces"
	"time"
)

type Event struct {
	id      string
	name    string
	service string
	table   string
	action  string
	payload string

	//audit

	createdAt time.Time
	updatedAt *time.Time
	isActive  bool
}

func (e *Event) GetId() string {
	return e.id
}

func (e *Event) GetName() string {
	return e.name
}

func (e *Event) GetService() string {
	return e.service
}

func (e *Event) GetTable() string {
	return e.table
}

func (e *Event) GetAction() string {
	return e.action
}

func (e *Event) GetPayload() string {
	return e.payload
}

func (e *Event) GetCreatedAt() time.Time {
	return e.createdAt
}

func (e *Event) GetUpdatedAt() *time.Time {
	return e.updatedAt
}

func (e *Event) GetIsActive() bool {
	return e.isActive
}

func NewEvent(idGenerator interfaces.IIdentityGenerator, name string, service string, table string, action string,
	payload string, createdAt time.Time,
) *Event {

	id := idGenerator.GetRandom(4)

	return &Event{
		id:        id,
		name:      name,
		service:   service,
		table:     table,
		action:    action,
		payload:   payload,
		createdAt: createdAt,
	}

}

func Assemble(id string, name string, service string, table string, action string, payload string, createdAt time.Time,
	updatedAt *time.Time, isActive bool,
) *Event {

	return &Event{
		id:        id,
		name:      name,
		service:   service,
		table:     table,
		action:    action,
		payload:   payload,
		createdAt: createdAt,
		updatedAt: updatedAt,
		isActive:  isActive,
	}

}
