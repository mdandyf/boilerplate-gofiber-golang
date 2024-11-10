package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var log zerolog.Logger

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	output := zerolog.ConsoleWriter{Out: os.Stdout}
	output.FormatTimestamp = func(i interface{}) string {
		return ""
	}
	log = zerolog.New(output).With().Timestamp().Logger()
}

// Info logs an informational message
func Info(msg string) {
	log.Info().Msg(msg)
}

// Error logs an error message
func Error(msg string) {
	log.Error().Msg(msg)
}

// Debug logs a debug message
func Debug(msg string) {
	log.Debug().Msg(msg)
}

// Warn logs a warning message
func Warn(msg string) {
	log.Warn().Msg(msg)
}