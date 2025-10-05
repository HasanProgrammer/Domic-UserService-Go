package entities

import (
	"domic.domain/commons/contracts/interfaces"
	"domic.domain/commons/entities"
	"time"
)

type RoleUser struct {
	entities.BaseEntity

	roleId string
	userId string
}

func NewRoleUser(idGenerator interfaces.IIdentityGenerator, roleId string,
	userId string, createdBy string, createdRole string,
) *RoleUser {

	id := idGenerator.GetRandom(4)
	nowTime := time.Now()

	roleUsereUser := &RoleUser{}

	roleUsereUser.BaseEntity.SetId(id)
	roleUsereUser.BaseEntity.SetCreatedBy(createdBy)
	roleUsereUser.BaseEntity.SetCreatedRole(createdRole)
	roleUsereUser.BaseEntity.SetCreatedAt(nowTime)

	roleUsereUser.roleId = roleId
	roleUsereUser.userId = userId

	//producing event

	return roleUsereUser
}

func Assemble(id string, roleId string, userId string, createdBy string, createdRole string, createdAt time.Time,
	updatedBy *string, updatedRole *string, updatedAt *time.Time,
) *RoleUser {

	roleUser := &RoleUser{}

	roleUser.BaseEntity.SetId(id)
	roleUser.BaseEntity.SetCreatedBy(createdBy)
	roleUser.BaseEntity.SetCreatedRole(createdRole)
	roleUser.BaseEntity.SetCreatedAt(createdAt)
	roleUser.BaseEntity.SetUpdatedBy(updatedBy)
	roleUser.BaseEntity.SetUpdatedRole(updatedRole)
	roleUser.BaseEntity.SetUpdatedAt(updatedAt)

	roleUser.roleId = roleId
	roleUser.userId = userId

	return roleUser

}

func (roleUser *RoleUser) GetUserId() string {
	return roleUser.userId
}

func (roleUser *RoleUser) GetRoleId() string {
	return roleUser.roleId
}
