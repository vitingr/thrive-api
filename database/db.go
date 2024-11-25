package database

import (
	"log"
	"main/models"
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

func setup(db *gorm.DB) {
	db.AutoMigrate(&models.Follower{}, &models.Group{}, &models.Like{}, &models.Post{}, &models.User{})
}

func ConnectionWithDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	stringConnection := os.Getenv("DATABASE_URL")
	if stringConnection == "" {
		log.Fatal("DATABASE_URL is not set in the environment")
	}

	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN: stringConnection,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		PrepareStmt: false,
		Logger:      logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	setup(DB)

	database, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get raw SQLDB from GORM: %v", err)
	}

	database.SetMaxOpenConns(999)
	database.SetMaxIdleConns(999)
	database.SetConnMaxLifetime(time.Minute * 15)

	if err := database.Ping(); err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	} else {
		log.Println("Database connection established successfully.")
	}
}
