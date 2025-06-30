package Models

import (
	"domic.domain/User/Entities"
	persistence "domic.persistence"
)

type UserModel struct {
	persistence.BaseModel

	FirstName string `gorm:"column:FirstName;type:varchar(80); not null"`
	LastName  string `gorm:"column:LastName;type:varchar(100); not null"`
	Username  string `gorm:"column:Username;type:varchar(20); not null"`
	Password  string `gorm:"column:Password;not null"`
	Email     string `gorm:"column:Email;not null"`
}

func MapUserEntityToModel(user *Entities.User) *UserModel {
	model := &UserModel{}

	model.FirstName = user.GetFirstName()
	model.LastName = user.GetLastName()
	model.Username = user.GetUsername()
	model.Password = user.GetPassword()
	model.Email = user.GetEmail()
	model.BaseModel.Id = user.GetId()
	model.BaseModel.CreatedBy = user.GetCreatedBy()
	model.BaseModel.CreatedAt = user.GetCreatedAt()
	model.BaseModel.CreatedRole = user.GetCreatedRole()
	model.BaseModel.UpdatedBy = user.GetUpdatedBy()
	model.BaseModel.UpdatedAt = user.GetUpdatedAt()
	model.BaseModel.UpdatedRole = user.GetUpdatedRole()
	model.BaseModel.IsActive = user.GetIsActive()

	return model
}

func MapUserEntitiesToModel(users []*Entities.User) []*UserModel {

	var models []*UserModel

	for _, user := range users {

		model := &UserModel{}

		model.FirstName = user.GetFirstName()
		model.LastName = user.GetLastName()
		model.Username = user.GetUsername()
		model.Password = user.GetPassword()
		model.Email = user.GetEmail()
		model.BaseModel.Id = user.GetId()
		model.BaseModel.CreatedBy = user.GetCreatedBy()
		model.BaseModel.CreatedAt = user.GetCreatedAt()
		model.BaseModel.CreatedRole = user.GetCreatedRole()
		model.BaseModel.UpdatedBy = user.GetUpdatedBy()
		model.BaseModel.UpdatedAt = user.GetUpdatedAt()
		model.BaseModel.UpdatedRole = user.GetUpdatedRole()
		model.BaseModel.IsActive = user.GetIsActive()

		models = append(models, model)

	}

	return models

}

func MapUserModelToEntity(model *UserModel) *Entities.User {
	return Entities.Assemble(model.Id, model.FirstName, model.LastName, model.Username, model.Password,
		model.Email, model.CreatedBy, model.CreatedRole, model.CreatedAt, model.UpdatedBy, model.UpdatedRole, model.UpdatedAt,
	)
}

func MapUserModelsToEntity(models []UserModel) []*Entities.User {

	var users []*Entities.User

	for _, model := range models {

		userEntity := Entities.Assemble(model.Id, model.FirstName, model.LastName, model.Username, model.Password,
			model.Email, model.CreatedBy, model.CreatedRole, model.CreatedAt, model.UpdatedBy, model.UpdatedRole,
			model.UpdatedAt,
		)

		users = append(users, userEntity)

	}

	return users

}
