package logging

import (
	"github.com/dowling-john/DjanGopher/config"
	"io"
	"log"
	"log/slog"
	"os"
)

// getHandlerType
func getHandlerType(configuration config.LoggingConfiguration, writer io.Writer) (handler slog.Handler) {
	switch configuration.HandlerType {
	case "json":
		return slog.NewJSONHandler(writer, nil)
	default:
		return slog.NewTextHandler(writer, nil)
	}
}

func getWriter(configuration config.LoggingConfiguration) (writer io.Writer) {
	switch configuration.WriterType {
	case "file":
		file, err := os.OpenFile(configuration.FileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal(err)
		}
		return file
	default:
		return os.Stdout
	}
}

// Init
// Initialize logging from configuration file
func Init(configuration config.LoggingConfiguration) (logger *slog.Logger) {
	return slog.New(
		getHandlerType(configuration, getWriter(configuration)),
	)
}
