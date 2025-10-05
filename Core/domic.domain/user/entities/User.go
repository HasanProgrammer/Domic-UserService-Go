package entities

import (
	"domic.domain/commons/contracts/interfaces"
	"domic.domain/commons/entities"
	"time"
)

type User struct {
	entities.BaseEntity

	firstName   string
	lastName    string
	username    string
	password    string
	email       string
	phoneNumber string
	imageUrl    string
	description string
}

func NewUser(idGenerator interfaces.IIdentityGenerator, firstName string, lastName string,
	username string, password string, email string, phoneNumber string, imageUrl string, description string,
	createdBy string, createdRole string,
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
	user.phoneNumber = phoneNumber
	user.imageUrl = imageUrl
	user.description = description

	//producing event

	user.BaseEntity.AppendEvent(entities.NewEvent(idGenerator, "UserCreated", "UserService", "user", "CREATE", "", nowTime))

	return user
}

func Assemble(id string, firstName string, lastName string,
	username string, password string, email string, phoneNumber string, imageUrl string, description string,
	createdBy string, createdRole string, createdAt time.Time,
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
	user.phoneNumber = phoneNumber
	user.imageUrl = imageUrl
	user.description = description

	return user

}

func (user *User) GetEvents() []*entities.Event {
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

func (user *User) GetPhoneNumber() string {
	return user.phoneNumber
}

func (user *User) GetImageUrl() string {
	return user.imageUrl
}

func (user *User) GetDescription() string {
	return user.description
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
