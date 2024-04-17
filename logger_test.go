// Tests for logging functions
package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/suite"
)

const testLogFile = "test.log"

// LoggerTests is a suite of tests for logging functions
type LoggerTests struct {
	*suite.Suite
	testLogger zerolog.Logger
	logFile    string
}

const (
	exampleMsg   = "example log message"
	infoMessage  = "example Info message"
	debugMessage = "example debug message"
	warnMessage  = "warning: something bad"
)

// TestLoggerSuite runs all tests
func TestLoggerSuite(t *testing.T) {
	suite.Run(t, new(LoggerTests))
}

// TestLogFileExists test the logfile is created correctly
func (t *LoggerTests) TestLogFileExists() {

	foo := "foo-test.log"
	defer os.Remove(foo)

	// File doesn't exist
	t.Assert().NoFileExistsf(foo, "test log file should not exist")

	// Make logger -file should now exist
	makeLogger(foo)
	t.Require().FileExists(foo)
}

// TestLogDefault tests logging an Default-level message
func (t *LoggerTests) TestLogMsg() {

	// Log a standard message
	t.testLogger.Log().Msg(exampleMsg)

	// Message and level should be in content
	s, _ := t.readTestFile()
	t.Assert().Containsf(t, s, infoMessage, "log message should exists in log file")
}

// TestLogInfo tests logging an info-level message
func (t *LoggerTests) TestLogInfo() {

	// Log an info level message
	t.testLogger.Info().Msg(infoMessage)

	// Message and level should be in content
	s, _ := t.readTestFile()
	t.Assert().Containsf(t, s, "\"info\"", "log message should contain info level")
	t.Assert().Containsf(t, s, infoMessage, "log message should exists in log file")
}

// TestLogInfo tests logging an debug-level message
func (t *LoggerTests) TestLogDebug() {

	// Log an debug level message
	t.testLogger.Debug().Msg(debugMessage)

	// Message and level should be in content
	s, _ := t.readTestFile()
	t.Assert().Containsf(t, s, "\"debug\"", "log message should contain debug level")
	t.Assert().Containsf(t, s, debugMessage, "log message should exists in log file")
}

// TestLogWarn tests logging an warn-level message
func (t *LoggerTests) TestLogWarn() {

	// Log a warn level message
	t.testLogger.Warn().Msg(warnMessage)

	// Message and level should be in content
	s, _ := t.readTestFile()
	t.Assert().Containsf(t, s, "\"warn\"", "log message should contain warn level")
	t.Assert().Containsf(t, s, warnMessage, "log message should exists in log file")
}

// Setup the test logger
func (t *LoggerTests) SetupSuite() {
	t.logFile = testLogFile
	t.testLogger = makeLogger(testLogFile)
}

// Cleanup test files
func (t *LoggerTests) TearDownSuite() {
	fmt.Println("cleaning up test logfiles...")
	os.Remove(testLogFile)
}

// readTestFile reads the content of the test logfile
func (t *LoggerTests) readTestFile() (string, error) {
	data, err := os.ReadFile(t.logFile)
	return string(data), err
}
