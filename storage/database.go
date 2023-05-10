package storage

import (
	"fmt"
	"log"
	"os"

	"github.com/chicho69-cesar/dio-planner-back/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connectToDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("Error loading .env file: %s", err))
	}

	dsn := os.Getenv("DB_CONNECTION_STRING")
	db, dbError := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbError != nil {
		log.Panicf("Error connecting to database: %s", dbError)
	}

	DB = db
	return db
}

func performMigrations(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
		&models.Event{},
		&models.Guest{},
	)
}

func InitializeDB() *gorm.DB {
	db := connectToDB()
	performMigrations(db)
	return db
}
