package DomainUserEntity

import (
	"Dotris.Domain/Commons/Entities"
	"errors"
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
}

func (user *User) Id() string {
	return user.id
}

func (user *User) Events() []*DomainCommonEntity.Event {
	return user.events
}

func (user *User) Change(FirstName string, LastName string, Username string, Password string, Email string) {
	user.firstName = FirstName
	user.lastName = LastName
	user.username = Username
	user.password = Password
	user.email = Email
}

func (user *User) ChangeStatus(IsActive bool) {
	user.isActive = IsActive
}

func NewUser(Id string, FirstName string, LastName string, Username string, Password string, email string) (*User, error) {

	var e error

	//validations

	if len(FirstName) >= 100 {
		e = errors.New("FirstName must be less than 100 characters")
	}

	if len(LastName) >= 100 {
		e = errors.New("LastName must be less than 100 characters")
	}

	return &User{
		id:        Id,
		firstName: FirstName,
		lastName:  LastName,
		username:  Username,
		password:  Password,
		email:     email,
		isActive:  true,
	}, e

}
