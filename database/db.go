package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectionWithDatabase() {
	godotenv.Load()
	stringConnection := os.Getenv("DATABASE_URL")
	DB, err = gorm.Open(postgres.Open(stringConnection))

	if err != nil {
		log.Panic("Error connecting to database")
	}

}