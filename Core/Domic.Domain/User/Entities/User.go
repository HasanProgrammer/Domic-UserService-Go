package DomainUserEntity

import (
	"Domic.Domain/Commons/Consts"
	"Domic.Domain/Commons/Contracts"
	"Domic.Domain/Commons/Entities"
	"encoding/json"
	"errors"
	"time"
)

type User struct {
	events []*DomainCommonEntity.Event

	id        string
	firstName string
	lastName  string
	username  string
	password  string
	email     string
	isActive  bool
	version   string

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

func (u *User) Active(idGenerator DomainCommonContract.IGlobalIdentityGenerator, updatedBy string, updatedRole string) error {

	nowTime := time.Now()

	u.isActive = false
	u.updatedAt = &nowTime
	u.updatedBy = updatedBy
	u.updatedRole = &updatedRole

	eventPayload, err := json.Marshal(u)

	if err != nil {
		return err
	}

	u.events = append(u.events,
		DomainCommonEntity.NewEvent(
			idGenerator.Generate(),
			"UserActived",
			"User",
			DomainCommonConst.UPDATE,
			string(eventPayload),
			nowTime,
			updatedBy,
			updatedRole,
		),
	)

	return nil
}

func NewUser(idGenerator DomainCommonContract.IGlobalIdentityGenerator, serializer DomainCommonContract.ISerializer, firstName string, lastName string,
	username string, password string, email string, createdBy string, createdRole string,
) (*User, error) {

	if len(firstName) >= 100 {
		return nil, errors.New("فیلد نام باید کمتر از 100 عبارت داشته باشد")
	}

	if len(lastName) >= 200 {
		return nil, errors.New("فیلد نام خانوادگی باید کمتر از 200 عبارت داشته باشد")
	}

	nowTime := time.Now()

	user := &User{
		id:          idGenerator.Generate(),
		firstName:   firstName,
		lastName:    lastName,
		username:    username,
		password:    password,
		email:       email,
		version:     idGenerator.Generate(),
		createdAt:   nowTime,
		createdBy:   createdBy,
		createdRole: createdRole,
	}

	//producing event

	eventPayload, err := serializer.Serialize(struct {
		Id          string     `json:"id"`
		FirstName   string     `json:"firstName"`
		LastName    string     `json:"lastName"`
		Username    string     `json:"username"`
		Password    string     `json:"password"`
		Email       string     `json:"email"`
		IsActive    bool       `json:"isActive"`
		CreatedAt   time.Time  `json:"createdAt"`
		CreatedBy   string     `json:"createdBy"`
		CreatedRole string     `json:"createdRole"`
		UpdatedAt   *time.Time `json:"updatedAt"`
		UpdatedBy   string     `json:"updatedBy"`
		UpdatedRole *string    `json:"updatedRole"`
	}{
		Id:          user.id,
		FirstName:   user.firstName,
		LastName:    user.lastName,
		Username:    user.username,
		Password:    user.password,
		Email:       user.email,
		IsActive:    user.isActive,
		CreatedAt:   user.createdAt,
		CreatedBy:   user.createdBy,
		CreatedRole: user.createdRole,
		UpdatedAt:   user.updatedAt,
		UpdatedBy:   user.updatedBy,
		UpdatedRole: user.updatedRole,
	})

	if err != nil {
		return nil, err
	}

	user.events = append(user.events,
		DomainCommonEntity.NewEvent(
			idGenerator.Generate(),
			"UserCreated",
			"User",
			DomainCommonConst.CREATE,
			eventPayload,
			nowTime,
			createdBy,
			createdRole,
		),
	)

	return user, nil
}
