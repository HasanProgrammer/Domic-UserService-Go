package models

import "time"

type BaseModel struct {
	Id          string     `gorm:"primaryKey;column:GetId"`
	CreatedAt   time.Time  `gorm:"column:GetCreatedAt"`
	CreatedBy   string     `gorm:"column:GetCreatedBy"`
	CreatedRole string     `gorm:"column:GetCreatedRole"`
	UpdatedAt   *time.Time `gorm:"column:GetUpdatedAt"`
	UpdatedBy   *string    `gorm:"column:GetUpdatedBy"`
	UpdatedRole *string    `gorm:"column:GetUpdatedRole"`
	IsActive    bool       `gorm:"column:GetIsActive"`
}
