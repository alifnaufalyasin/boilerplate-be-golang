package utils

import (
	"os"
	"runtime/debug"
	"time"

	"github.com/rs/zerolog"
)

type Loggers struct {
	Logger    zerolog.Logger
	goVersion string
}

func NewLogger(level string) Loggers {
	writers := zerolog.MultiLevelWriter(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		NoColor:    false,
		TimeFormat: time.RFC1123,
	})
	Logger := zerolog.New(writers)
	switch level {
	case "debug":
		Logger.Level(zerolog.DebugLevel)
	case "info":
		Logger.Level(zerolog.InfoLevel)
	case "warn":
		Logger.Level(zerolog.WarnLevel)
	case "error":
		Logger.Level(zerolog.ErrorLevel)
	case "fatal":
		Logger.Level(zerolog.FatalLevel)
	default:
		Logger.Level(zerolog.InfoLevel)
	}

	buildInfo, _ := debug.ReadBuildInfo()

	return Loggers{
		Logger:    Logger,
		goVersion: buildInfo.GoVersion,
	}
}

func (l *Loggers) Fatal(err error, msg string) {
	l.Logger.Fatal().Timestamp().Caller().Err(err).Msg(msg)
}

func (l *Loggers) Fatalf(err error, msg string, v ...interface{}) {
	l.Logger.Fatal().Timestamp().Caller().Err(err).Msgf(msg, v...)
}

func (l *Loggers) Error(err error, msg string) {
	l.Logger.Error().Timestamp().Caller().Err(err).Msg(msg)
}

func (l *Loggers) Errorf(err error, msg string, v ...interface{}) {
	l.Logger.Error().Timestamp().Caller().Err(err).Msgf(msg, v...)
}

func (l *Loggers) Info(msg string) {
	l.Logger.Info().Timestamp().Caller().Msg(msg)
}

func (l *Loggers) Infof(msg string, v ...interface{}) {
	l.Logger.Info().Timestamp().Caller().Msgf(msg, v...)
}

func (l *Loggers) Debug(msg string) {
	l.Logger.Debug().Timestamp().Caller().Msg(msg)
}

func (l *Loggers) Debugf(msg string, v ...interface{}) {
	l.Logger.Debug().Timestamp().Caller().Msgf(msg, v...)
}
