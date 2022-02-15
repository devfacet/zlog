// zlog
// For the full copyright and license information, please view the LICENSE.txt file.

// Package zlog provides abstraction layer for zerolog logging library.
package zlog

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

const (
	defTimeFormat = "2006-01-02T15:04:05.000"
	defLocation   = "UTC"
)

var (
	defLogger *zerolog.Logger
)

// Logger returns a new logger by the given arguments.
func Logger(level, format, output, location string) *zerolog.Logger {
	// Check location
	zerolog.TimeFieldFormat = defTimeFormat
	if location == "" {
		location = defLocation
	}
	if l, err := time.LoadLocation(location); err == nil {
		zerolog.TimestampFunc = func() time.Time {
			return time.Now().In(l)
		}
	}

	// Init the logger
	var logger zerolog.Logger

	// Output
	var writer *os.File
	switch output {
	case "stderr":
		writer = os.Stderr
	default:
		writer = os.Stdout
	}

	// Format
	switch format {
	case "console":
		logger = zerolog.New(zerolog.ConsoleWriter{Out: writer, TimeFormat: defTimeFormat})
	default:
		logger = zerolog.New(writer)
	}

	// Level
	switch level {
	case "debug":
		logger = logger.Level(zerolog.DebugLevel)
	case "error":
		logger = logger.Level(zerolog.ErrorLevel)
	case "fatal":
		logger = logger.Level(zerolog.FatalLevel)
	case "info":
		logger = logger.Level(zerolog.InfoLevel)
	case "panic":
		logger = logger.Level(zerolog.PanicLevel)
	case "warning":
		logger = logger.Level(zerolog.WarnLevel)
	default:
		logger = logger.Level(zerolog.InfoLevel)
	}

	logger = logger.With().Timestamp().Logger()
	return &logger
}

// DefaultLogger returns the default logger.
func DefaultLogger() *zerolog.Logger {
	if defLogger == nil {
		defLogger = Logger("", "", "", "")
	}
	return defLogger
}

// SetDefaultLogger sets the default logger.
func SetDefaultLogger(l *zerolog.Logger) {
	defLogger = l
}
