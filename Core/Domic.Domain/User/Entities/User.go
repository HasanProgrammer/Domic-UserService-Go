package DomainUserEntity

import (
	"Domic.Domain/Commons/Entities"
	"errors"
	"time"
)

type User struct {
	events    []*DomainCommonEntity.Event
	id        string
	firstName string
	lastName  string
	username  string
	password  string
	email     string
	isActive  bool

	//audit fields

	createdAt   time.Time
	createdBy   string
	createdRole string
	updatedAt   *time.Time
	updatedBy   string
	updatedRole *string
}

func (u *User) GetEvents() []*DomainCommonEntity.Event {
	return u.events
}

func (u *User) GetId() string {
	return u.id
}

func (u *User) GetFirstName() string {
	return u.firstName
}

func (u *User) GetLastName() string {
	return u.lastName
}

func (u *User) GetUsername() string {
	return u.username
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) GetIsActive() bool {
	return u.isActive
}

func (u *User) GetCreatedAt() time.Time {
	return u.createdAt
}

func (u *User) GetCreatedBy() string {
	return u.createdBy
}

func (u *User) GetCreatedRole() string {
	return u.createdRole
}

func (u *User) GetUpdatedAt() *time.Time {
	return u.updatedAt
}

func (u *User) GetUpdatedBy() string {
	return u.updatedBy
}

func (u *User) GetUpdatedRole() *string {
	return u.updatedRole
}

func NewUser(id string, firstName string, lastName string, username string, password string, email string, createdBy string, createdRole string) (*User, error) {

	if len(firstName) >= 100 {
		return nil, errors.New("")
	}

	if len(lastName) >= 200 {
		return nil, errors.New("")
	}

	nowTime := time.Now()

	user := &User{
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
		DomainCommonEntity.NewEvent("", "UserCreated", "User", "Create", "", nowTime, createdBy, createdRole),
	)

	return user, nil
}
