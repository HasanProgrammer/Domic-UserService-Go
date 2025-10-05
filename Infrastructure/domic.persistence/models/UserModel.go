package models

import (
	"domic.domain/user/entities"
)

type UserModel struct {
	BaseModel

	FirstName   string `gorm:"column:FirstName;type:varchar(80); not null"`
	LastName    string `gorm:"column:LastName;type:varchar(100); not null"`
	Username    string `gorm:"column:Username;type:varchar(20); not null"`
	Password    string `gorm:"column:Password;not null"`
	Email       string `gorm:"column:Email;not null"`
	PhoneNumber string `gorm:"column:PhoneNumber;not null"`
	Description string `gorm:"column:Description;not null"`
	ImageUrl    string `gorm:"column:Description;not null"`
}

func ConvertUserEntityToModel(user *entities.User) *UserModel {
	model := &UserModel{}

	model.FirstName = user.GetFirstName()
	model.LastName = user.GetLastName()
	model.Username = user.GetUsername()
	model.Password = user.GetPassword()
	model.Email = user.GetEmail()
	model.PhoneNumber = user.GetPhoneNumber()
	model.Description = user.GetDescription()
	model.ImageUrl = user.GetImageUrl()
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

func ConvertUserEntitiesToModels(users []*entities.User) []*UserModel {

	var models []*UserModel

	for _, user := range users {

		model := &UserModel{}

		model.FirstName = user.GetFirstName()
		model.LastName = user.GetLastName()
		model.Username = user.GetUsername()
		model.Password = user.GetPassword()
		model.Email = user.GetEmail()
		model.PhoneNumber = user.GetPhoneNumber()
		model.Description = user.GetDescription()
		model.ImageUrl = user.GetImageUrl()
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

func ConvertUserModelToEntity(model *UserModel) *entities.User {
	return entities.Assemble(model.Id, model.FirstName, model.LastName, model.Username, model.Password,
		model.Email, model.PhoneNumber, model.ImageUrl, model.Description, model.CreatedBy, model.CreatedRole, model.CreatedAt, model.UpdatedBy, model.UpdatedRole, model.UpdatedAt,
	)
}

func ConvertUserModelsToEntities(models []UserModel) []*entities.User {

	var users []*entities.User

	for _, model := range models {

		userEntity := entities.Assemble(model.Id, model.FirstName, model.LastName, model.Username, model.Password,
			model.Email, model.PhoneNumber, model.ImageUrl, model.Description, model.CreatedBy, model.CreatedRole,
			model.CreatedAt, model.UpdatedBy, model.UpdatedRole, model.UpdatedAt,
		)

		users = append(users, userEntity)

	}

	return users

}
