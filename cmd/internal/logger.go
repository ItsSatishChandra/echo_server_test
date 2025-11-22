package internal

import (
	"log"
	"os"
	"path/filepath"
)

var LOG_FILE_LOCATION string = "data/echo_server.log"
var LOGGER_FILE *os.File

func InitLogFile(fileLocation string) (file *os.File) {
	var fileLocationToUse string
	if fileLocation != "" {
		fileLocationToUse = fileLocation
	} else {
		fileLocationToUse = LOG_FILE_LOCATION
	}
	logFile, err := createLogFile(fileLocationToUse)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	LOGGER_FILE = logFile
	return logFile
}

func createLogFile(path string) (*os.File, error) {
	dir := filepath.Dir(path)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, err
		}
	}

	return os.OpenFile(path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
}

func SetLogFileLocation(fileLocation string) {
	logFile := InitLogFile(fileLocation)
	log.SetFlags(log.LstdFlags | log.LUTC)
	if logFile != nil {
		log.SetOutput(logFile)
	}
}

func EchoLogger(source string, logData string) {
	log.Println(logData)
}

func Echo(source string, input string) string {
	EchoLogger(source, input)
	return input
}
