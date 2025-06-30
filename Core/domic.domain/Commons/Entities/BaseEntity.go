package Entities

import "time"

type BaseEntity struct {
	events      []*Event
	id          string
	isActive    bool
	createdAt   time.Time
	createdBy   string
	createdRole string
	updatedAt   *time.Time
	updatedBy   *string
	updatedRole *string
}

func (b *BaseEntity) Events() []*Event {
	return b.events
}

func (b *BaseEntity) AppendEvent(event *Event) {
	b.events = append(b.events, event)
}

func (b *BaseEntity) Id() string {
	return b.id
}

func (b *BaseEntity) SetId(id string) {
	b.id = id
}

func (b *BaseEntity) IsActive() bool {
	return b.isActive
}

func (b *BaseEntity) SetIsActive(isActive bool) {
	b.isActive = isActive
}

func (b *BaseEntity) CreatedAt() time.Time {
	return b.createdAt
}

func (b *BaseEntity) SetCreatedAt(createdAt time.Time) {
	b.createdAt = createdAt
}

func (b *BaseEntity) CreatedBy() string {
	return b.createdBy
}

func (b *BaseEntity) SetCreatedBy(createdBy string) {
	b.createdBy = createdBy
}

func (b *BaseEntity) CreatedRole() string {
	return b.createdRole
}

func (b *BaseEntity) SetCreatedRole(createdRole string) {
	b.createdRole = createdRole
}

func (b *BaseEntity) UpdatedAt() *time.Time {
	return b.updatedAt
}

func (b *BaseEntity) SetUpdatedAt(updatedAt *time.Time) {
	b.updatedAt = updatedAt
}

func (b *BaseEntity) UpdatedBy() *string {
	return b.updatedBy
}

func (b *BaseEntity) SetUpdatedBy(updatedBy *string) {
	b.updatedBy = updatedBy
}

func (b *BaseEntity) UpdatedRole() *string {
	return b.updatedRole
}

func (b *BaseEntity) SetUpdatedRole(updatedRole *string) {
	b.updatedRole = updatedRole
}
