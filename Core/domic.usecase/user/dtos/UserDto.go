package dtos

import "time"

type RoleDto struct {
	Id   string
	Name string
}

type PermissionDto struct {
	Id   string
	Name string
}

type UserDto struct {
	Id          string
	ImageUrl    string
	FirstName   string
	LastName    string
	Username    string
	PhoneNumber string
	Email       string
	Description string
	IsActive    bool
	CreatedAt   time.Time
	Roles       []RoleDto
	Permissions []PermissionDto
}
