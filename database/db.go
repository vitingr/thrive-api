package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
)

func ConnectionWithDatabase() {
    godotenv.Load()
    stringConnection := os.Getenv("DATABASE_URL")
    var err error
    DB, err = gorm.Open(postgres.Open(stringConnection), &gorm.Config{
        PrepareStmt: false,
        Logger:      logger.Default.LogMode(logger.Info),
    })

    if err != nil {
        log.Panic("Error connecting to database")
    }
}
