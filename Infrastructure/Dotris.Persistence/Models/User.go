package InfrastructureModel

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model

	Id        string `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Username  string
	Password  string
	Email     string
	IsActive  bool
}
