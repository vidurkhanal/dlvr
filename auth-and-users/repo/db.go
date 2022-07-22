package repo

import (
	"auth-and-users/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var dbError error

func Connect(connectionString string) {
	DB, dbError = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")
}

func Migrate() {
	DB.AutoMigrate(&models.User{})
	log.Println("Database Migration Completed!")
}
