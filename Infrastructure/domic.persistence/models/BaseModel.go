package models

import "time"

type BaseModel struct {
	Id          string     `gorm:"primaryKey;column:Id"`
	CreatedAt   time.Time  `gorm:"column:CreatedAt"`
	CreatedBy   string     `gorm:"column:CreatedBy"`
	CreatedRole string     `gorm:"column:CreatedRole"`
	UpdatedAt   *time.Time `gorm:"column:UpdatedAt"`
	UpdatedBy   *string    `gorm:"column:UpdatedBy"`
	UpdatedRole *string    `gorm:"column:UpdatedRole"`
	IsActive    bool       `gorm:"column:IsActive"`
}
