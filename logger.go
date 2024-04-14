// Logging functions
package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"
)

type AppLogger struct {
	*slog.Logger
}

// TeeWriter writes to stdout and to file
// see: https://rednafi.com/go/structured_logging_with_slog/
type TeeWriter struct {
	stdout *os.File
	file   *os.File
}

func (t *TeeWriter) Write(p []byte) (n int, err error) {
	n, err = t.stdout.Write(p)
	if err != nil {
		return n, err
	}
	n, err = t.file.Write(p)
	return n, err
}

// New creates a new AppLogger, which will write to stdout and logfile
func (l AppLogger) New(logfile string) *AppLogger {
	file, _ := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY|os.O_SYNC, 0666)
	writer := &TeeWriter{
		stdout: os.Stdout,
		file:   file,
	}
	handler := slog.NewTextHandler(writer, nil)
	logger := slog.New(handler)
	return &AppLogger{logger}
}

// logfileName is a dated logfile name in format
func logfileName() string {
	return fmt.Sprintf("logs-%s.log", time.Now().Format("2006-01-02"))
}
