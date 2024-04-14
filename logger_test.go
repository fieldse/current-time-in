// Tests for logging functions
package main

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testLogFile = "test.log"

var testLogger = AppLogger{}.New(testLogFile)

// Remove any files if already existing
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

	// Log message. File should exist
	testLogger.Info("example log message")
	require.FileExistsf(t, testLogFile, "test log file should exist")

	// Read content
	data, err := os.ReadFile(testLogFile)
	assert.Nilf(t, err, "read file error %v", err)
	assert.NotEmpty(t, data)

	// Message should be in content, and have prefix
	var s = string(data)
	assert.Truef(t, strings.HasPrefix(s, "[+]"), "log message should start with prefix")
	assert.Containsf(t, s, "example log message", "log message should exists in log file")

	// Test debug message
	testLogger.Debug("example debug message")
	data, _ = os.ReadFile(testLogFile)
	s = string(data)
	assert.Containsf(t, s, "example log message", "original log message should still exists in log file")
	assert.Containsf(t, s, "[DEBUG]", "debug prefix should exist in log file")
	assert.Containsf(t, s, "example debug message", "debug message should exist in log file")

}
