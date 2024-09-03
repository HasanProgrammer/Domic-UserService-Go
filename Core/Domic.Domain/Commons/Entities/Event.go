package DomainCommonEntity

import "time"

type Event struct {
	id          string
	name        string
	table       string
	action      string
	payload     string
	createdAt   time.Time
	createdBy   string
	createdRole string
}

func (e *Event) GetId() string {
	return e.id
}

func (e *Event) SetId(id string) {
	e.id = id
}

func (e *Event) GetName() string {
	return e.name
}

func (e *Event) SetName(name string) {
	e.name = name
}

func (e *Event) GetTable() string {
	return e.table
}

func (e *Event) SetTable(table string) {
	e.table = table
}

func (e *Event) GetAction() string {
	return e.action
}

func (e *Event) SetAction(action string) {
	e.action = action
}

func (e *Event) GetPayload() string {
	return e.payload
}

func (e *Event) SetPayload(payload string) {
	e.payload = payload
}

func (e *Event) GetCreatedAt() time.Time {
	return e.createdAt
}

func (e *Event) SetCreatedAt(createdAt time.Time) {
	e.createdAt = createdAt
}

func (e *Event) GetCreatedBy() string {
	return e.createdBy
}

func (e *Event) SetCreatedBy(createdBy string) {
	e.createdBy = createdBy
}

func (e *Event) GetCreatedRole() string {
	return e.createdRole
}

func (e *Event) SetCreatedRole(createdRole string) {
	e.createdRole = createdRole
}

func NewEvent(id string, name string, table string, action string, payload string, createdAt time.Time, createdBy string, createdRole string) *Event {
	return &Event{id: id, name: name, table: table, action: action, payload: payload, createdAt: createdAt, createdBy: createdBy, createdRole: createdRole}
}
