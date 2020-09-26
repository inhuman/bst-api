package log

import (
	"github.com/rs/zerolog"
	"os"
)

func NewLogger() *zerolog.Logger {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	l := zerolog.New(os.Stdout).With().Logger()
	return &l
}
