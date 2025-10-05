package models

import (
	"domic.domain/role/entities"
)

type RoleModel struct {
	BaseModel

	Name string `gorm:"column:Name;type:varchar(80); not null"`

	//relations
	RoleUsers []RoleUserModel `gorm:"foreignKey:RoleId"`
}

func ConvertRoleEntityToModel(role *entities.Role) *RoleModel {
	model := &RoleModel{}

	model.Name = role.GetName()
	model.BaseModel.Id = role.GetId()
	model.BaseModel.CreatedBy = role.GetCreatedBy()
	model.BaseModel.CreatedAt = role.GetCreatedAt()
	model.BaseModel.CreatedRole = role.GetCreatedRole()
	model.BaseModel.UpdatedBy = role.GetUpdatedBy()
	model.BaseModel.UpdatedAt = role.GetUpdatedAt()
	model.BaseModel.UpdatedRole = role.GetUpdatedRole()
	model.BaseModel.IsActive = role.GetIsActive()

	return model
}

func ConvertRoleEntitiesToModels(roles []*entities.Role) []*RoleModel {

	var models []*RoleModel

	for _, role := range roles {

		model := &RoleModel{}

		model.Name = role.GetName()
		model.BaseModel.Id = role.GetId()
		model.BaseModel.CreatedBy = role.GetCreatedBy()
		model.BaseModel.CreatedAt = role.GetCreatedAt()
		model.BaseModel.CreatedRole = role.GetCreatedRole()
		model.BaseModel.UpdatedBy = role.GetUpdatedBy()
		model.BaseModel.UpdatedAt = role.GetUpdatedAt()
		model.BaseModel.UpdatedRole = role.GetUpdatedRole()
		model.BaseModel.IsActive = role.GetIsActive()

		models = append(models, model)

	}

	return models

}

func ConvertRoleModelToEntity(model *RoleModel) *entities.Role {
	return entities.Assemble(model.Id, model.Name, model.CreatedBy, model.CreatedRole, model.CreatedAt,
		model.UpdatedBy, model.UpdatedRole, model.UpdatedAt,
	)
}

func ConvertRoleModelsToEntities(models []RoleModel) []*entities.Role {

	var roles []*entities.Role

	for _, model := range models {

		userEntity := entities.Assemble(model.Id, model.Name, model.CreatedBy, model.CreatedRole,
			model.CreatedAt, model.UpdatedBy, model.UpdatedRole, model.UpdatedAt,
		)

		roles = append(roles, userEntity)

	}

	return roles

}
