package Persistence

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type SqlContext struct {
	connectionString string
}

func (sqlContext *SqlContext) GetContext() *gorm.DB {
	db, err := gorm.Open(sqlserver.Open(sqlContext.connectionString), &gorm.Config{})

	if err != nil {
	}

	return db
}

func NewSqlContext(connectionString string) *SqlContext {
	return &SqlContext{connectionString: connectionString}
}
