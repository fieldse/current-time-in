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

	// File doesn't exist
	assert.NoFileExistsf(t, testLogFile, "test log file should not exist")
	testLogger := makeLogger(testLogFile)

	// Log a standard message
	exampleMsg := "example log message"
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

	// Log a debug message
	testLogger.Debug().Msg(("example debug message"))

	// Message and level should be in content
	data, _ = os.ReadFile(testLogFile)
	s = string(data)
	assert.Containsf(t, s, "\"debug\"", "log message should contain debug level")
	assert.Containsf(t, s, "example debug message", "log message should exists in log file")

	// Log a warn message
	warnMessage := "warning: something bad"
	testLogger.Warn().Msg((warnMessage))

	// Message and level should be in content
	data, _ = os.ReadFile(testLogFile)
	s = string(data)
	assert.Containsf(t, s, "\"warn\"", "log message should contain warn level")
	assert.Containsf(t, s, warnMessage, "log message should exists in log file")
}
