// Package logger with Logrus
package main

import "github.com/sirupsen/logrus"

const LOG_LEVEL = logrus.DebugLevel

var AppLogger *logrus.Logger

func init() {
	// TODO -- make this support logging to file
	AppLogger = logrus.New()
	AppLogger.SetLevel(LOG_LEVEL)
	AppLogger.SetFormatter(&logrus.TextFormatter{
		// DisableQuote: true,
	})
}
