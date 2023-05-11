package storage

import (
	"fmt"
	"log"
	"os"
	"database/sql"
	
	"github.com/chicho69-cesar/dio-planner-back/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
  _ "github.com/lib/pq"
)

var DB *gorm.DB
var PostgresDB *sql.DB

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

func connectToPostgresDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("Error loading .env file: %s", err))
	}

	dsn := os.Getenv("DB_CONNECTION_STRING")

	db, dbError := sql.Open("postgres", dsn)
	if dbError != nil {
		log.Panicf("Error connecting to database: %s", dbError)
	}

	PostgresDB = db
	return db
}

func performMigrations(db *gorm.DB) {
	db.AutoMigrate(
		&models.Event{},
		&models.Grade{},
		&models.Guest{},
		&models.Memory{},
		&models.Purchase{},
		&models.Todo{},
		&models.User{},
	)
}

func InitializeDB() *gorm.DB {
	db := connectToDB()
	performMigrations(db)
	return db
}

func InitializePostgresDB() *sql.DB {
	db := connectToPostgresDB()
	return db
}

func ClosePostgresDB() {
	PostgresDB.Close()
}
