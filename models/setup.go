package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Setup database
	dsn := "host=localhost user=metheuser password=mysecretpassword dbname=mydb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("There is an error connecting to the database")
		return
	}
	db.AutoMigrate(&Item{})

	DB = db
}
