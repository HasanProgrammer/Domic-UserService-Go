package DomainCommonContract

import "time"

type BaseEntity[TIdentity any] struct {
	id          TIdentity
	createdBy   string
	createdAt   time.Time
	createdRole string
	updatedBy   *string    //nullable
	updatedAt   *time.Time //nullable
	updateRole  *string    //nullable
	version     string
}
