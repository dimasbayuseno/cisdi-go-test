package logger

import (
	"context"
	"github.com/dimasbayuseno/cisdi-go-test/config"
	"github.com/rs/zerolog"
	"os"
	"sync"
)

var loggerInstance zerolog.Logger
var once sync.Once

func Log(ctx context.Context) *zerolog.Logger {
	once.Do(func() {
		multi := zerolog.MultiLevelWriter(os.Stdout)
		loggerInstance = zerolog.New(multi).With().Timestamp().Logger()

		if config.Get().Env == "dev" {
			loggerInstance = loggerInstance.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		}

	})

	newlogger := loggerInstance.With().Ctx(ctx).Logger()
	return &newlogger
}
