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

func (roleUser *RoleUser) GetEvents() []*entities.Event {
	return roleUser.BaseEntity.Events()
}

func (roleUser *RoleUser) GetId() string {
	return roleUser.BaseEntity.Id()
}

func (roleUser *RoleUser) GetUserId() string {
	return roleUser.userId
}

func (roleUser *RoleUser) GetRoleId() string {
	return roleUser.roleId
}

func (roleUser *RoleUser) GetIsActive() bool {
	return roleUser.BaseEntity.IsActive()
}

func (roleUser *RoleUser) GetCreatedAt() time.Time {
	return roleUser.BaseEntity.CreatedAt()
}

func (roleUser *RoleUser) GetCreatedBy() string {
	return roleUser.BaseEntity.CreatedBy()
}

func (roleUser *RoleUser) GetCreatedRole() string {
	return roleUser.BaseEntity.CreatedRole()
}

func (roleUser *RoleUser) GetUpdatedAt() *time.Time {
	return roleUser.BaseEntity.UpdatedAt()
}

func (roleUser *RoleUser) GetUpdatedBy() *string {
	return roleUser.BaseEntity.UpdatedBy()
}

func (roleUser *RoleUser) GetUpdatedRole() *string {
	return roleUser.BaseEntity.UpdatedRole()
}
