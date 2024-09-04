package InfrastructureModel

import (
	"gorm.io/gorm"
	"time"
)

type EventModel struct {
	gorm.Model

	Id          string    `gorm:"primaryKey" gorm:"column:Id"`
	Name        string    `gorm:"column:Name"`
	Table       string    `gorm:"column:Table"`
	Action      string    `gorm:"column:Action"`
	Payload     string    `json:"payload" gorm:"column:Payload"`
	CreatedAt   time.Time `gorm:"column:CreatedAt"`
	CreatedBy   string    `gorm:"column:CreatedBy"`
	CreatedRole string    `gorm:"column:CreatedRole"`
}

func (model *EventModel) TableName() string {
	return "Events"
}
