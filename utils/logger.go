package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func InitLogger() {
	// Output ke terminal dan file (opsional)
	Logger.SetOutput(os.Stdout)

	// Set format ke JSON agar mudah dibaca sistem monitoring
	Logger.SetFormatter(&logrus.JSONFormatter{})

	// Set level logging (DEBUG, INFO, WARN, ERROR, FATAL)
	Logger.SetLevel(logrus.DebugLevel)
}
