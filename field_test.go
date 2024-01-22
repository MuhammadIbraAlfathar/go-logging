package go_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "ibra").Info("Hello world")

	logger.WithField("email", "test@gmail.com").
		WithField("address", "Jl. Diponogoro").
		Info("Test")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "Ibra",
		"email":    "test@gmail.com",
		"name":     "Ibra Alfathar",
	}).Info("Test Logger")
}
