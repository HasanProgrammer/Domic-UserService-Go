package Entities

import (
	"domic.domain/Commons/Contracts/Interfaces"
	"domic.domain/Commons/Entities"
	"time"
)

type User struct {
	Entities.BaseEntity

	firstName string
	lastName  string
	username  string
	password  string
	email     string
}

func New(idGenerator Interfaces.IIdentityGenerator, firstName string, lastName string,
	username string, password string, email string, createdBy string, createdRole string,
) *User {

	id := idGenerator.GetRandom(4)
	nowTime := time.Now()

	user := &User{}

	user.BaseEntity.SetId(id)
	user.BaseEntity.SetCreatedBy(createdBy)
	user.BaseEntity.SetCreatedRole(createdRole)
	user.BaseEntity.SetCreatedAt(nowTime)

	user.firstName = firstName
	user.lastName = lastName
	user.username = username
	user.password = password
	user.email = email

	//producing event

	user.BaseEntity.AppendEvent(Entities.NewEvent(id, "UserCreated", "UserService", "User", "CREATE", "", nowTime))

	return user
}

func Assemble(id string, firstName string, lastName string,
	username string, password string, email string, createdBy string, createdRole string, createdAt time.Time,
	updatedBy *string, updatedRole *string, updatedAt *time.Time,
) *User {

	user := &User{}

	user.BaseEntity.SetId(id)
	user.BaseEntity.SetCreatedBy(createdBy)
	user.BaseEntity.SetCreatedRole(createdRole)
	user.BaseEntity.SetCreatedAt(createdAt)
	user.BaseEntity.SetUpdatedBy(updatedBy)
	user.BaseEntity.SetUpdatedRole(updatedRole)
	user.BaseEntity.SetUpdatedAt(updatedAt)

	user.firstName = firstName
	user.lastName = lastName
	user.username = username
	user.password = password
	user.email = email

	return user

}

func (user *User) GetEvents() []*Entities.Event {
	return user.BaseEntity.Events()
}

func (user *User) GetId() string {
	return user.BaseEntity.Id()
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
	return user.BaseEntity.IsActive()
}

func (user *User) GetCreatedAt() time.Time {
	return user.BaseEntity.CreatedAt()
}

func (user *User) GetCreatedBy() string {
	return user.BaseEntity.CreatedBy()
}

func (user *User) GetCreatedRole() string {
	return user.BaseEntity.CreatedRole()
}

func (user *User) GetUpdatedAt() *time.Time {
	return user.BaseEntity.UpdatedAt()
}

func (user *User) GetUpdatedBy() *string {
	return user.BaseEntity.UpdatedBy()
}

func (user *User) GetUpdatedRole() *string {
	return user.BaseEntity.UpdatedRole()
}
