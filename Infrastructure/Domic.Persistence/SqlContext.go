package Persistence

import (
	"Domic.Persistence/Models"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type SqlContext struct {
	db *gorm.DB
}

func (sqlContext *SqlContext) GetContext() *gorm.DB {
	return sqlContext.db
}

func NewSqlContext(connectionString string) *SqlContext {
	db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})

	db.AutoMigrate(&InfrastructureModel.EventModel{})
	db.AutoMigrate(&InfrastructureModel.UserModel{})

	if err != nil {
	}

	return &SqlContext{db: db}
}
