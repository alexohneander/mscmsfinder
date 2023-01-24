package database

import (
	"fmt"
	"mscmsfinder/config"
	"mscmsfinder/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ConnectDB connect to db
func ConnectDB() {
	var err error
	databaseFile := config.Config("DB_FILE")
	DB, err := gorm.Open(sqlite.Open(databaseFile), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&model.CMS{})
	fmt.Println("Database Migrated")
}
