// Logging functions
package main

import (
	"fmt"
	"log"
	"os"
)

const LOGFILE_NAME = "log.txt"

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	DebugLogger *log.Logger
)

// Initialize the three loggers
func init() {
	InfoLogger = makeLogger("[+]")
	ErrorLogger = makeLogger("[ERROR]")
	DebugLogger = makeLogger("[DEBUG]")
}

// makeLogger returns a new logger instance that writes to outfile (log.txt)
func makeLogger(prefix string) *log.Logger {
	file, err := os.OpenFile(LOGFILE_NAME, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("initialize logger failed: %v", err)
	}
	return log.New(file, prefix+" ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Log logs a standard info-level message to file
func Log(msg string, args ...interface{}) {
	InfoLogger.Printf(msg, args...)
}

// LogDebug logs a debug-level message to file
func LogDebug(msg string, args ...interface{}) {
	DebugLogger.Printf(msg, args...)
}

// LogError logs an error message to file
func LogError(err error, msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	ErrorLogger.Printf(fmt.Sprintf("%s: %w", msg, err.Error))
}
