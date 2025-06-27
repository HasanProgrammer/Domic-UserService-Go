package Models

import (
	"domic.domain/User/Entities"
	"gorm.io/gorm"
	"time"
)

type UserModel struct {
	gorm.Model

	Id        string `gorm:"primaryKey"`
	FirstName string `gorm:"column:FirstName;type:varchar(80); not null"`
	LastName  string `gorm:"column:LastName;type:varchar(100); not null"`
	Username  string `gorm:"column:Username;type:varchar(20); not null"`
	Password  string `gorm:"column:Password;not null"`
	Email     string `gorm:"column:Email;not null"`
	IsActive  bool   `gorm:"column:IsActive"`

	//audit fields

	CreatedAt   time.Time `gorm:"column:CreatedAt"`
	CreatedBy   string    `gorm:"column:CreatedBy"`
	CreatedRole string    `gorm:"column:CreatedRole"`

	UpdatedAt   *time.Time `gorm:"column:UpdatedAt"`
	UpdatedBy   *string    `gorm:"column:UpdatedBy"`
	UpdatedRole *string    `gorm:"column:UpdatedRole"`
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

func MapUserEntitiesToModel(users []*Entities.User) []*UserModel {

	var models []*UserModel

	for _, user := range users {

		model := &UserModel{
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

		models = append(models, model)

	}

	return models

}

func MapUserModelToEntity(model *UserModel) *Entities.User {
	return Entities.AssembleUser(model.Id, model.FirstName, model.LastName, model.Username, model.Password,
		model.Email, model.CreatedBy, model.CreatedRole, model.CreatedAt, model.UpdatedBy, model.UpdatedRole, model.UpdatedAt,
	)
}

func MapUserModelsToEntity(models []UserModel) []*Entities.User {

	var users []*Entities.User

	for _, model := range models {

		userEntity := Entities.AssembleUser(model.Id, model.FirstName, model.LastName, model.Username, model.Password,
			model.Email, model.CreatedBy, model.CreatedRole, model.CreatedAt, model.UpdatedBy, model.UpdatedRole,
			model.UpdatedAt,
		)

		users = append(users, userEntity)

	}

	return users

}
