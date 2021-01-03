package utils

import (
	"io"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

// Logger is our zerolog logger
var Logger zerolog.Logger

// GetLogLevel returns a zerolog representation of the given string log level
func GetLogLevel(level string) zerolog.Level {
	lvl := strings.ToLower(level)
	switch lvl {
	case "trace":
		return zerolog.TraceLevel
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	case "panic":
		return zerolog.PanicLevel
	default:
		return zerolog.InfoLevel
	}
}

// SetupLogger configures the zerolog logger based on the config
func SetupLogger(config BotConfig) {
	zerolog.TimeFieldFormat = time.RFC3339
	lvl := GetLogLevel(config.LogLevel)

	var writers []io.Writer

	if config.LogConsole {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	}

	if len(config.LogFile) > 0 {
		// writers = append(writers, zerolog.)
	}

	multi := io.MultiWriter(writers...)
	Logger = zerolog.New(multi).Level(lvl).With().Timestamp().Logger()

	Logger.Info().Msg("Logger configured.")
}
