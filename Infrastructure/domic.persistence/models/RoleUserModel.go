package models

import (
	"domic.domain/role_user/entities"
)

type RoleUserModel struct {
	BaseModel

	RoleId string `gorm:"column:RoleId; not null"`
	UserId string `gorm:"column:UserId; not null"`

	User UserModel `gorm:"references:Id"`
	Role RoleModel `gorm:"references:Id"`
}

func ConvertRoleUserEntityToModel(roleUser *entities.RoleUser) *RoleUserModel {
	model := &RoleUserModel{}

	model.RoleId = roleUser.GetRoleId()
	model.UserId = roleUser.GetUserId()
	model.BaseModel.Id = roleUser.GetId()
	model.BaseModel.CreatedBy = roleUser.GetCreatedBy()
	model.BaseModel.CreatedAt = roleUser.GetCreatedAt()
	model.BaseModel.CreatedRole = roleUser.GetCreatedRole()
	model.BaseModel.UpdatedBy = roleUser.GetUpdatedBy()
	model.BaseModel.UpdatedAt = roleUser.GetUpdatedAt()
	model.BaseModel.UpdatedRole = roleUser.GetUpdatedRole()
	model.BaseModel.IsActive = roleUser.GetIsActive()

	return model
}

func ConvertRoleUserEntitiesToModels(roleUsers []*entities.RoleUser) []*RoleUserModel {

	var models []*RoleUserModel

	for _, roleUser := range roleUsers {

		model := &RoleUserModel{}

		model.RoleId = roleUser.GetRoleId()
		model.UserId = roleUser.GetUserId()
		model.BaseModel.Id = roleUser.GetId()
		model.BaseModel.CreatedBy = roleUser.GetCreatedBy()
		model.BaseModel.CreatedAt = roleUser.GetCreatedAt()
		model.BaseModel.CreatedRole = roleUser.GetCreatedRole()
		model.BaseModel.UpdatedBy = roleUser.GetUpdatedBy()
		model.BaseModel.UpdatedAt = roleUser.GetUpdatedAt()
		model.BaseModel.UpdatedRole = roleUser.GetUpdatedRole()
		model.BaseModel.IsActive = roleUser.GetIsActive()

		models = append(models, model)

	}

	return models

}
