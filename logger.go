// Logging functions
package main

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/rs/zerolog"
)

var (
	Logger  zerolog.Logger
	LogFile string
)

const YYYYMMDD string = "2006-01-02"
const LOGGER_LEVEL_DEBUG = zerolog.DebugLevel

func init() {
	LogFile = logfileCurrentDate()
	Logger = makeLogger(path.Join("logs", LogFile))
}

// logfileCurrentDate returns logfile name in format "logs-[YYYY-MM-DD].log"
func logfileCurrentDate() string {
	return fmt.Sprintf("logs-%s.log", time.Now().Format(YYYYMMDD))
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

	defer file.Close()

	return zerolog.New(file).Level(LOGGER_LEVEL_DEBUG).With().Timestamp().Logger()
}
