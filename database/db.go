package database

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

func ConnectionWithDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	stringConnection := os.Getenv("DATABASE_URL")
	if stringConnection == "" {
		log.Fatal("DATABASE_URL is not set in the environment")
	}

	DB, err = gorm.Open(postgres.Open(stringConnection), &gorm.Config{
		PrepareStmt: true, 
		Logger:      logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	database, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get raw SQLDB from GORM: %v", err)
	}

  database.Exec("DEALLOCATE ALL")

	database.SetMaxOpenConns(20)
	database.SetMaxIdleConns(10)
	database.SetConnMaxLifetime(time.Hour) 

	if err := database.Ping(); err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	} else {
		log.Println("Database connection established successfully.")
	}
}
