package entities

import (
	"domic.domain/commons/contracts/interfaces"
	"domic.domain/commons/entities"
	"time"
)

type Role struct {
	entities.BaseEntity

	name string
}

func NewRole(idGenerator interfaces.IIdentityGenerator, name string,
	createdBy string, createdRole string,
) *Role {

	id := idGenerator.GetRandom(4)
	nowTime := time.Now()

	role := &Role{}

	role.BaseEntity.SetId(id)
	role.BaseEntity.SetCreatedBy(createdBy)
	role.BaseEntity.SetCreatedRole(createdRole)
	role.BaseEntity.SetCreatedAt(nowTime)

	role.name = name

	//producing event

	role.BaseEntity.AppendEvent(
		entities.NewEvent(idGenerator, "RoleCreated", "UserService", "role", "CREATE", "", nowTime),
	)

	return role
}

func Assemble(id string, name string, createdBy string, createdRole string, createdAt time.Time,
	updatedBy *string, updatedRole *string, updatedAt *time.Time,
) *Role {

	role := &Role{}

	role.BaseEntity.SetId(id)
	role.BaseEntity.SetCreatedBy(createdBy)
	role.BaseEntity.SetCreatedRole(createdRole)
	role.BaseEntity.SetCreatedAt(createdAt)
	role.BaseEntity.SetUpdatedBy(updatedBy)
	role.BaseEntity.SetUpdatedRole(updatedRole)
	role.BaseEntity.SetUpdatedAt(updatedAt)

	role.name = name

	return role

}

func (role *Role) GetName() string {
	return role.name
}
