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

var (
	Logger *AppLogger
)

type AppLogger struct {
	logfile     string
	infoLogger  *log.Logger
	errorLogger *log.Logger
	debugLogger *log.Logger
}

func init() {
	f := path.Join("logs", logfileName())
	Logger = AppLogger{}.New(f)
}

// logfileName is a dated logfile name in format "logs-[YYYY-MM-DD].log"
func logfileName() string {
	return fmt.Sprintf("logs-%s.log", time.Now().Format("2006-01-02"))
}

// New initializes the logger with info, debug, and error loggers.
func (l AppLogger) New(logFile string) *AppLogger {
	return &AppLogger{
		logfile:     logFile,
		infoLogger:  makeLogger(logFile, "[+]"),
		errorLogger: makeLogger(logFile, "[ERROR]"),
		debugLogger: makeLogger(logFile, "[DEBUG]"),
	}
}

// Info logs a standard info-level message to file
func (l *AppLogger) Info(msg string, args ...interface{}) {
	l.infoLogger.Printf(msg, args...)
}

// Debug logs a debug-level message to file
func (l *AppLogger) Debug(msg string, args ...interface{}) {
	l.debugLogger.Printf(msg, args...)
}

// Error logs an error message to file
func (l *AppLogger) Error(err error, msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	l.errorLogger.Printf(fmt.Sprintf("%s: %v", msg, err.Error()))
}

// makeLogger returns a new log.Logger instance that writes to both STDOUT and logfile
func makeLogger(logFile string, prefix string) *log.Logger {
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("initialize logger failed: %v", err)
	}
	// Open multiwriter to log both to console and to file
	mw := io.MultiWriter(os.Stdout, file)

	return log.New(mw, prefix+" ", log.Ldate|log.Ltime|log.Lshortfile)
}
