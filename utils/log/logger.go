package log

import (
	"os"

	"github.com/agusheryanto182/go-schedule/models/config"
	"github.com/rs/zerolog"
)

func NewLogger(c *config.Global) *zerolog.Logger {
	var logger zerolog.Logger
	logger = zerolog.New(os.Stdout)
	logger.Level(c.Log.Level)
	logger = logger.With().Timestamp().Logger()

	return &logger
}
