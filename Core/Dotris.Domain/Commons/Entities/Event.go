package DomainCommonEntity

import "time"

type Event struct {
	id          string
	name        string
	table       string
	action      string
	payload     string `json:"payload"`
	createdAt   time.Time
	createdBy   string
	createdRole string
}

func (e *Event) Id() string {
	return e.id
}

func (e *Event) SetId(id string) {
	e.id = id
}

func (e *Event) CreatedRole() string {
	return e.createdRole
}

func (e *Event) SetCreatedRole(createdRole string) {
	e.createdRole = createdRole
}

func (e *Event) CreatedBy() string {
	return e.createdBy
}

func (e *Event) SetCreatedBy(createdBy string) {
	e.createdBy = createdBy
}

func (e *Event) CreatedAt() time.Time {
	return e.createdAt
}

func (e *Event) SetCreatedAt(createdAt time.Time) {
	e.createdAt = createdAt
}

func (e *Event) Payload() string {
	return e.payload
}

func (e *Event) SetPayload(payload string) {
	e.payload = payload
}

func (e *Event) Action() string {
	return e.action
}

func (e *Event) SetAction(action string) {
	e.action = action
}

func (e *Event) Table() string {
	return e.table
}

func (e *Event) SetTable(table string) {
	e.table = table
}

func (e *Event) Name() string {
	return e.name
}

func (e *Event) SetName(name string) {
	e.name = name
}

func NewEvent(createdRole string, createdBy string, createdAt time.Time, payload string, action string, table string, name string, id string) *Event {
	return &Event{createdRole: createdRole, createdBy: createdBy, createdAt: createdAt, payload: payload, action: action, table: table, name: name, id: id}
}
