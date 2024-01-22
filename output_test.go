package go_logging

import (
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestOutput(t *testing.T) {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})

	file, _ := os.OpenFile("application.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)

	logger.SetOutput(file)

	logger.Info("Test Logger")
	logger.Warn("Test Logger Warn")
	logger.Error("Test Logger Error")
}
