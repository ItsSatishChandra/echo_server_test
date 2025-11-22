package echo

import (
	"log"
	"os"
)

var LOG_FILE_LOCATION string = "/app/data/echo_server.log"
var LOGGER_FILE *os.File

func InitLogFile(fileLocation string) {
	var fileLocationToUse string
	if fileLocation != "" {
		fileLocationToUse = fileLocation
	} else {
		fileLocationToUse = LOG_FILE_LOCATION
	}
	logFile, err := os.OpenFile(fileLocationToUse, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	LOGGER_FILE = logFile
}

func SetLogFileLocation(fileLocation string) {
	InitLogFile(fileLocation)
	log.SetFlags(log.LstdFlags | log.LUTC)
	log.SetOutput(LOGGER_FILE)
}

func EchoLogger(logData string) {
	log.Println(logData)
}

func Echo(input string) string {
	EchoLogger(input)
	return input
}
