package Entities

import (
	"domic.domain/Commons/Contracts/Interfaces"
	"domic.domain/Commons/Entities"
	"time"
)

type User[TIdentity any] struct {
	events    []*Entities.Event[TIdentity]
	id        TIdentity
	firstName string
	lastName  string
	username  string
	password  string
	email     string
	isActive  bool

	//audit fields

	createdAt   time.Time
	createdBy   TIdentity
	createdRole string

	updatedAt   *time.Time
	updatedBy   *TIdentity
	updatedRole *string
}

func NewUser(idGenerator Interfaces.IIdentityGenerator, firstName string, lastName string,
	username string, password string, email string, createdBy string, createdRole string,
) *User[string] {

	id := idGenerator.GetRandom(4)
	nowTime := time.Now()

	user := &User[string]{
		id:          id,
		firstName:   firstName,
		lastName:    lastName,
		username:    username,
		password:    password,
		email:       email,
		createdAt:   nowTime,
		createdBy:   createdBy,
		createdRole: createdRole,
	}

	//producing event

	user.events = append(user.events,
		Entities.NewEvent[string](id, "UserCreated", "UserService", "User", "CREATE", "", nowTime, createdBy,
			createdRole,
		),
	)

	return user
}

func (user *User[string]) GetEvents() []*Entities.Event[string] {
	return user.events
}
