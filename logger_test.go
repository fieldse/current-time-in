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

func cleanup() {
	fmt.Println("cleaning up test logfiles...")
	os.Remove(testLogFile)
}

func Test_Log(t *testing.T) {
	defer cleanup()

	// File doesn't exist
	assert.NoFileExistsf(t, testLogFile, "test log file should not exist")
	testLogger := makeLogger(testLogFile)

	// Log message. File should exist
	testLogger.Log().Msg(("example log message"))
	assert.FileExistsf(t, testLogFile, "test log file should exist")

	// Read content
	data, err := os.ReadFile(testLogFile)
	require.Nilf(t, err, "read file error %v", err)
	require.NotEmpty(t, data)

	// Message should be in content, and have prefix
	var s = string(data)
	assert.Truef(t, strings.HasPrefix(s, "[TEST]"), "log message should start with prefix")
	assert.Containsf(t, s, "example log message", "log message should exists in log file")
}
