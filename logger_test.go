// Tests for logging functions
package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testLogFile = "test.log"

func setup() {
	cleanup()
}

func cleanup() {
	fmt.Println("cleaning up test logfiles...")
	os.Remove(testLogFile)
}

func Test_Log(t *testing.T) {
	setup()
	defer cleanup()

	const (
		exampleMsg   = "example log message"
		infoMessage  = "example Info message"
		debugMessage = "example debug message"
		warnMessage  = "warning: something bad"
	)

	// File doesn't exist
	assert.NoFileExistsf(t, testLogFile, "test log file should not exist")
	testLogger := makeLogger(testLogFile)

	// Log a standard message
	testLogger.Log().Msg(exampleMsg)

	// File should now exist
	assert.FileExistsf(t, testLogFile, "test log file should exist")

	// Read content
	data, err := os.ReadFile(testLogFile)
	require.Nilf(t, err, "read file error %v", err)

	// Message should be in content
	s := string(data)
	require.NotEmpty(t, data)
	assert.Containsf(t, s, exampleMsg, "log message should exists in log file")

	// Log an info level message
	testLogger.Info().Msg(infoMessage)

	// Message and level should be in content
	data, _ = os.ReadFile(testLogFile)
	s = string(data)
	assert.Containsf(t, s, "\"info\"", "log message should contain info level")
	assert.Containsf(t, s, infoMessage, "log message should exists in log file")

	// Log a debug level message
	testLogger.Debug().Msg(debugMessage)

	// Message and level should be in content
	data, _ = os.ReadFile(testLogFile)
	s = string(data)
	assert.Containsf(t, s, "\"debug\"", "log message should contain debug level")
	assert.Containsf(t, s, debugMessage, "log message should exists in log file")

	// Log a warn level message
	testLogger.Warn().Msg(warnMessage)

	// Message and level should be in content
	data, _ = os.ReadFile(testLogFile)
	s = string(data)
	assert.Containsf(t, s, "\"warn\"", "log message should contain warn level")
	assert.Containsf(t, s, warnMessage, "log message should exists in log file")
}
