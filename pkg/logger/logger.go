// Logging functions
package logger

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/rs/zerolog"
)

var (
	Logger         zerolog.Logger
	LogFile        string
	LOGS_DIRECTORY string
)

const LOGGER_LEVEL_DEBUG = zerolog.DebugLevel

func init() {
	LogFile = path.Join("..", "..", "logs", logfileCurrentDate())
	Logger = makeLogger(LogFile)
	LOGS_DIRECTORY, _ = getLogsDir()
}

// makeLogger returns a new log.Logger instance that writes to both STDOUT and logfile
func makeLogger(logFile string) zerolog.Logger {
	file, err := os.OpenFile(
		logFile,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	if err != nil {
		panic(err)
	}

	// Multi-out logger
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	output := zerolog.MultiLevelWriter(consoleWriter, file)

	return zerolog.New(output).Level(LOGGER_LEVEL_DEBUG).
		With().
		Timestamp().
		Caller().
		Logger()
}

// logfileCurrentDate returns logfile name in format "logs-[YYYY-MM-DD].log"
func logfileCurrentDate() string {
	return fmt.Sprintf("logs-%s.log", time.Now().Format(time.DateOnly))
}

// getLogsDir attempts to resolve the log directory
func getLogsDir() (string, error) {
	curDir, err := os.Getwd()                     // current working directory
	rootDir := filepath.Dir(filepath.Dir(curDir)) // get grandparent directory of current working directory
	fp := path.Join(rootDir, "logs")
	if err != nil {
		return "", err
	}
	return fp, nil
}
