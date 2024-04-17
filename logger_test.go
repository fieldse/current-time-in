// Tests for logging functions
package main

import (
	"log"
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/suite"
)

const (
	TESTING_LOGFILE = "test.log"
	exampleMsg      = "example log message"
	infoMessage     = "example Info message"
	debugMessage    = "example debug message"
	warnMessage     = "warning: something bad"
)

// LoggerTests is a suite of tests for logging functions
type LoggerTests struct {
	suite.Suite
	testLogger zerolog.Logger
	logFile    string
}

// Run all tests
func TestLoggerSuite(t *testing.T) {
	log.Println("TestLoggerSuite()")
	suite.Run(t, new(LoggerTests))
}

// TestLogFileExists test the logfile is created correctly
func (t *LoggerTests) TestLogFileExists() {
	log.Println("TestLogFileExists()")

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
	log.Println("TestLogMsg()")
	msg := exampleMsg

	// Log a standard message
	t.testLogger.Log().Msg(msg)

	// Message and level should be in content
	s := t.readTestFile()
	t.Assert().Containsf(s, msg, "log message should exists in log file")
}

// TestLogInfo tests logging an info-level message
func (t *LoggerTests) TestLogInfo() {
	log.Println("TestLogInfo()")

	msg := infoMessage
	lvl := `"level":"info"`

	// Log an info level message
	t.testLogger.Info().Msg(infoMessage)

	// Message and level should be in content
	s := t.readTestFile()
	t.Assert().Containsf(s, lvl, "log message should contain level %s", lvl)
	t.Assert().Containsf(s, msg, "log message should exists in log file")
}

// TestLogInfo tests logging an debug-level message
func (t *LoggerTests) TestLogDebug() {
	log.Println("TestLogDebug()")
	msg := debugMessage
	lvl := `"level":"debug"`

	// Log an debug level message
	t.testLogger.Debug().Msg(debugMessage)

	// Message and level should be in content
	s := t.readTestFile()
	t.Assert().Containsf(s, lvl, "log message should contain level %s", lvl)
	t.Assert().Containsf(s, msg, "log message should exists in log file")
}

// TestLogWarn tests logging an warn-level message
func (t *LoggerTests) TestLogWarn() {
	log.Println("TestLogWarn()")
	msg := warnMessage
	lvl := `"level":"warn"`

	// Log a warn level message
	t.testLogger.Warn().Msg(warnMessage)

	// Message and level should be in content
	s := t.readTestFile()
	t.Assert().Containsf(s, lvl, "log message should contain level %s", lvl)
	t.Assert().Containsf(s, msg, "log message should exists in log file")
}

// Setup the test logger
func (t *LoggerTests) SetupSuite() {
	log.Println("SetupSuite()")
	t.logFile = TESTING_LOGFILE
	t.testLogger = makeLogger(TESTING_LOGFILE)
}

// Cleanup test files
func (t *LoggerTests) TearDownSuite() {
	log.Println("TearDownSuite()")
	os.Remove(TESTING_LOGFILE)
}

// readTestFile reads the content of the test logfile
func (t *LoggerTests) readTestFile() string {
	log.Println("readTestFile()")
	data, err := os.ReadFile(t.logFile)
	if err != nil {
		log.Fatalf("readTestFile() error: %v", err)
	}
	return string(data)
}
