package logger

import (
	"os"

	"github.com/mattn/go-colorable"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// New returns a configured zerolog.Logger.
// If env == "development", use console writer; otherwise structured JSON.
func New(env string) zerolog.Logger {
	if env == "development" {
		cw := zerolog.ConsoleWriter{Out: colorable.NewColorableStdout()}
		cw.TimeFormat = "2006-01-02 15:04:05"
		logger := zerolog.New(cw).With().Timestamp().Logger()
		return logger
	}
	logger := log.Output(os.Stdout)
	return logger
}
