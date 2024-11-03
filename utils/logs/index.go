package utils

import (
	"log"
	"os"
)

var Logger *log.Logger

func init() {
	Logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func LogInfo(message string) {
	Logger.Println("INFO: " + message)
}

func LogError(message string) {
	Logger.Println("ERROR: " + message)
}
