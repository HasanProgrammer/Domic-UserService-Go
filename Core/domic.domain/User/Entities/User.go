package Entities

import (
	"domic.domain/Commons/Contracts/Interfaces"
	"domic.domain/Commons/Entities"
	"time"
)

type User struct {
	events    []*Entities.Event
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
	updatedBy   *string
	updatedRole *string
}

func NewUser(idGenerator Interfaces.IIdentityGenerator, firstName string, lastName string,
	username string, password string, email string, createdBy string, createdRole string,
) *User {

	id := idGenerator.GetRandom(4)
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
		Entities.NewEvent(id, "UserCreated", "UserService", "User", "CREATE", "", nowTime),
	)

	return user
}

func AssembleUser(id string, firstName string, lastName string,
	username string, password string, email string, createdBy string, createdRole string, createdAt time.Time,
	updatedBy *string, updatedRole *string, updatedAt *time.Time,
) *User {

	nowTime := time.Now()

	user := &User{
		id:          id,
		firstName:   firstName,
		lastName:    lastName,
		username:    username,
		password:    password,
		email:       email,
		createdBy:   createdBy,
		createdRole: createdRole,
		createdAt:   nowTime,
		updatedBy:   updatedBy,
		updatedRole: updatedRole,
		updatedAt:   updatedAt,
	}

	return user
}

func (user *User) GetEvents() []*Entities.Event {
	return user.events
}

func (user *User) GetId() string {
	return user.id
}

func (user *User) GetFirstName() string {
	return user.firstName
}

func (user *User) GetLastName() string {
	return user.lastName
}

func (user *User) GetUsername() string {
	return user.username
}

func (user *User) GetPassword() string {
	return user.password
}

func (user *User) GetEmail() string {
	return user.email
}

func (user *User) GetIsActive() bool {
	return user.isActive
}

func (user *User) GetCreatedAt() time.Time {
	return user.createdAt
}

func (user *User) GetCreatedBy() string {
	return user.createdBy
}

func (user *User) GetCreatedRole() string {
	return user.createdRole
}

func (user *User) GetUpdatedAt() *time.Time {
	return user.updatedAt
}

func (user *User) GetUpdatedBy() *string {
	return user.updatedBy
}

func (user *User) GetUpdatedRole() *string {
	return user.updatedRole
}
