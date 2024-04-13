// Logging functions
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"time"
)

// Logs dir is at /logs
var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	DebugLogger *log.Logger
)

// Initialize the three loggers
func init() {
	// logs-YYYY-MM-DD.log
	logFile := path.Join("logs", fmt.Sprintf("logs-%s.log", time.Now().Format("2006-01-02")))
	InfoLogger = makeLogger(logFile, "[+]")
	ErrorLogger = makeLogger(logFile, "[ERROR]")
	DebugLogger = makeLogger(logFile, "[DEBUG]")
}

// makeLogger returns a new logger instance that writes to both STDOUT and logfile (log.txt)
func makeLogger(logFile string, prefix string) *log.Logger {
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("initialize logger failed: %v", err)
	}
	// Open multiwriter to log both to console and to file
	mw := io.MultiWriter(os.Stdout, file)

	return log.New(mw, prefix+" ", log.Ldate|log.Ltime|log.Lshortfile)
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
	ErrorLogger.Printf(fmt.Sprintf("%s: %v", msg, err.Error()))
}
