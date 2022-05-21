package logger

import (
	"log"
	"os"
)

type Logger struct {
	ErrorLogger *log.Logger
	InfoLogger  *log.Logger
}

func NewLogger() *Logger {
	logg := new(Logger)
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logg.ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	logg.InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime)
	return logg
}
