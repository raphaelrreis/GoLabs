package logger

import (
	"log"
	"os"
)

var Logger = log.New(os.Stdout, "GoLabs: ", log.LstdFlags)

func Info(msg string) {
	Logger.Println("[INFO]", msg)
}

func Error(msg string) {
	Logger.Println("[ERROR]", msg)
}
