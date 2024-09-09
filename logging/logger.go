package logging

import (
	"github.com/dowling-john/DjanGopher/config"
	"log"
	"os"
)

func createFileLogger(fileName string) *log.Logger {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	return log.New(f, "", log.Ldate|log.Ltime)
}

// Init
// Initialize logging from configuration file
func Init(configuration config.LoggingConfiguration) *log.Logger {
	switch configuration.WriterType {
	case "console":
		return log.New(os.Stdout, "", log.Ldate|log.Ltime)
	case "file":
		return createFileLogger(configuration.FileName)
	default:
		return log.New(os.Stderr, "", log.Ldate|log.Ltime)
	}
}
