package Entities

import "time"

type Event struct {
	id      string
	name    string
	service string
	table   string
	action  string
	payload string

	//audit

	createdAt time.Time
	updatedAt *string
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

func NewEvent(id string, name string, service string, table string, action string, payload string,
	createdAt time.Time,
) *Event {

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
