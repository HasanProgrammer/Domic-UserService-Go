package Models

import "time"

type EventModel struct {
	Id        string    `gorm:"column:Id"`
	Name      string    `gorm:"column:Name"`
	Service   string    `gorm:"column:Service"`
	Table     string    `gorm:"column:Table"`
	Action    string    `gorm:"column:Action"`
	Payload   string    `gorm:"column:Payload"`
	CreatedAt time.Time `gorm:"column:CreatedAt"`
	UpdatedAt *string   `gorm:"column:UpdatedAt"`
	IsActive  bool      `gorm:"column:IsActive"`
}
