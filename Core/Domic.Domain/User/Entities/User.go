package DomainUserEntity

import (
	"Domic.Domain/Commons/Entities"
	"errors"
	"time"
)

type User[TIdentity any] struct {
	events    []*DomainCommonEntity.Event[TIdentity]
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
	updatedBy   TIdentity
	updatedRole *string
}

func (u *User[TIdentity]) Events() []*DomainCommonEntity.Event[TIdentity] {
	return u.events
}

func (u *User[TIdentity]) Id() TIdentity {
	return u.id
}

func (u *User[TIdentity]) FirstName() string {
	return u.firstName
}

func (u *User[TIdentity]) LastName() string {
	return u.lastName
}

func (u *User[TIdentity]) Username() string {
	return u.username
}

func (u *User[TIdentity]) Password() string {
	return u.password
}

func (u *User[TIdentity]) Email() string {
	return u.email
}

func (u *User[TIdentity]) IsActive() bool {
	return u.isActive
}

func NewUser[TIdentity any](id TIdentity, firstName string, lastName string, username string, password string, email string, createdBy TIdentity, createdRole string) (*User[TIdentity], error) {

	if len(firstName) >= 100 {
		return nil, errors.New("")
	}

	if len(lastName) >= 200 {
		return nil, errors.New("")
	}

	nowTime := time.Now()

	user := &User[TIdentity]{
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
		DomainCommonEntity.NewEvent[TIdentity]("", "UserCreated", "User", "Create", "", nowTime, createdBy, createdRole),
	)

	return user, nil
}
