package InfrastructureModel

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model

	Id        string `gorm:"primaryKey" gorm:"column:Id"`
	FirstName string `gorm:"column:FirstName"`
	LastName  string `gorm:"column:LastName"`
	Username  string `gorm:"column:Username"`
	Password  string `gorm:"column:Password"`
	Email     string `gorm:"column:Email"`
	IsActive  bool   `gorm:"column:IsActive"`
}

func (model *UserModel) TableName() string {
	return "Users"
}
