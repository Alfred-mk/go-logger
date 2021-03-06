package main

import (
	"os"
	"log"
    "time"
)

var (
	WarningLogger *log.Logger
    InfoLogger    *log.Logger
    ErrorLogger   *log.Logger
)

func init() {
    currentTime := time.Now()
 
	var logFile string

	// Create the dynamic file based on the current date
	logFile = "go-logger-" + currentTime.Format("2006-01-02") + ".log"

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }

    InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
    ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	InfoLogger.Println("Starting the application...")
	WarningLogger.Println("The application encountered an issue...")
	ErrorLogger.Println("The application failed to run...")
}
