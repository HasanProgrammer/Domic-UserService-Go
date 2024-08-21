package InfrastructureModel

import (
	"gorm.io/gorm"
	"time"
)

type EventModel struct {
	gorm.Model

	Id          string `gorm:"primaryKey"`
	Name        string
	Table       string
	Action      string
	Payload     string `json:"payload"`
	CreatedAt   time.Time
	CreatedBy   string
	CreatedRole string
}
