package database

import (
	"log"

	"affan/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
  
func DatabaseConnection() *gorm.DB {
	dsn := "postgresql://postgres:suZzBmTIDjicSQJC55sA@containers-us-west-64.railway.app:5640/railway" // Change to your postgres database url
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database : ", err)
	}

	db.AutoMigrate(models.Photo{}, models.User{})
	println("DB connected")
	return db
}