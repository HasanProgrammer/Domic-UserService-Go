package Entities

import "time"

type Event[TIdentity any] struct {
	id          string
	name        string
	service     string
	table       string
	action      string
	payload     string
	createdAt   time.Time
	createdBy   TIdentity
	createdRole string
}

func (e *Event[TIdentity]) GetId() string {
	return e.id
}

func (e *Event[TIdentity]) GetName() string {
	return e.name
}

func (e *Event[TIdentity]) GetService() string {
	return e.service
}

func (e *Event[TIdentity]) GetTable() string {
	return e.table
}

func (e *Event[TIdentity]) GetAction() string {
	return e.action
}

func (e *Event[TIdentity]) GetPayload() string {
	return e.payload
}

func (e *Event[TIdentity]) GetCreatedAt() time.Time {
	return e.createdAt
}

func (e *Event[TIdentity]) GetCreatedBy() TIdentity {
	return e.createdBy
}

func (e *Event[TIdentity]) GetCreatedRole() string {
	return e.createdRole
}

func NewEvent[TIdentity any](id string, name string, service string, table string, action string, payload string,
	createdAt time.Time, createdBy TIdentity, createdRole string,
) *Event[TIdentity] {

	return &Event[TIdentity]{
		id:          id,
		name:        name,
		service:     service,
		table:       table,
		action:      action,
		payload:     payload,
		createdAt:   createdAt,
		createdBy:   createdBy,
		createdRole: createdRole,
	}

}
