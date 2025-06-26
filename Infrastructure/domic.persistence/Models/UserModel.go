package Models

import (
	"domic.domain/User/Entities"
	"gorm.io/gorm"
	"time"
)

type UserModel struct {
	gorm.Model

	Id        string `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Username  string
	Password  string
	Email     string
	IsActive  bool

	//audit fields

	CreatedAt   time.Time
	CreatedBy   string
	CreatedRole string

	UpdatedAt   *time.Time
	UpdatedBy   *string
	UpdatedRole *string
}

func MapUserEntityToModel(user *Entities.User) *UserModel {
	return &UserModel{
		Id:          user.GetId(),
		FirstName:   user.GetFirstName(),
		LastName:    user.GetLastName(),
		Username:    user.GetUsername(),
		Password:    user.GetPassword(),
		Email:       user.GetEmail(),
		CreatedBy:   user.GetCreatedBy(),
		CreatedAt:   user.GetCreatedAt(),
		CreatedRole: user.GetCreatedRole(),
		UpdatedBy:   user.GetUpdatedBy(),
		UpdatedAt:   user.GetUpdatedAt(),
		UpdatedRole: user.GetUpdatedRole(),
		IsActive:    user.GetIsActive(),
	}
}

func MapUserModelToEntity(model *UserModel) *Entities.User {
	return Entities.AssembleUser(model.Id, model.FirstName, model.LastName, model.Username, model.Password,
		model.Email, model.CreatedBy, model.CreatedRole, model.CreatedAt, model.UpdatedBy, model.UpdatedRole, model.UpdatedAt,
	)
}

func MapUserModelsToEntity(models []UserModel) []*Entities.User {

	var users []*Entities.User

	for _, model := range models {

		users = append(users, Entities.AssembleUser(model.Id, model.FirstName, model.LastName, model.Username, model.Password,
			model.Email, model.CreatedBy, model.CreatedRole, model.CreatedAt, model.UpdatedBy, model.UpdatedRole, model.UpdatedAt,
		))

	}

	return users
}
