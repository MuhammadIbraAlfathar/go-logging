package go_logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLogger(t *testing.T) {
	logger := logrus.New()

	logger.Println("Hello Logger")
	fmt.Println("Test")
}
