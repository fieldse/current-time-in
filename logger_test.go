// tests for app logger
package main

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func Test_AppLogger(t *testing.T) {
	assert.IsType(t, &logrus.Logger{}, AppLogger)

	AppLogger.Info("This is an info message")
	AppLogger.Warn("This is a Warn message")
	AppLogger.Debug("This is a debug message")
	AppLogger.Debugln("This is a debug line")
}
