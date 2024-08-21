package DomainUserEntity

import (
	"errors"
)

type User struct {
	id        string
	firstName string
	lastName  string
	username  string
	password  string
	email     string
	isActive  bool
}

func (user *User) GetId() string {
	return user.id
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

		firstName: FirstName,
		lastName:  LastName,
		username:  Username,
		password:  Password,
		email:     email,
		isActive:  true,
	}, e

}
