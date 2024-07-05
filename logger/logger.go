package logger

import (
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/spf13/viper"
)

// Logger is an interface for logging messages.
//
// It provides methods for logging messages at different levels, as well as
// attaching additional fields to the log message.
//
// Please note that the fields are key-value pairs, where the key is a string
// and the value is an interface{}.
type Logger interface {
	Debug(msg string, fields ...interface{})
	Info(msg string, fields ...interface{})
	Warn(msg string, fields ...interface{})
	Error(err error, msg string, fields ...interface{})
	Fatal(err error, msg string, fields ...interface{})
}

// logger is an implementation of Logger interface.
//
// It uses zerolog library to write logs to stdout.
type logger struct {
	log zerolog.Logger
}

// Debug logs a message at debug level.
func (l *logger) Debug(msg string, fields ...interface{}) {
	l.log.Debug().Fields(fields).Msg(msg)
}

// Info logs a message at info level.
func (l *logger) Info(msg string, fields ...interface{}) {
	l.log.Info().Fields(fields).Msg(msg)
}

// Warn logs a message at warn level.
func (l *logger) Warn(msg string, fields ...interface{}) {
	l.log.Warn().Fields(fields).Msg(msg)
}

// Error logs a message at error level.
func (l *logger) Error(err error, msg string, fields ...interface{}) {
	l.log.Error().Err(err).Fields(fields).Msg(msg)
}

// Fatal logs a message at fatal level.
func (l *logger) Fatal(err error, msg string, fields ...interface{}) {
	l.log.Fatal().Err(err).Fields(fields).Msg(msg)
}

var (
	once sync.Once
	log  *logger
)

// Get returns a singleton instance of Logger.
//
// It is configured to write logs to stdout in a colorized, human-friendly format.
//
// Logging is performed at the level specified by the log.level environment variable.
func Get() Logger {
	once.Do(func() {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		zerolog.TimeFieldFormat = time.RFC3339

		output := zerolog.ConsoleWriter{
			Out:        os.Stdout,
			NoColor:    false,
			TimeFormat: time.RFC3339,
		}

		log = &logger{
			log: zerolog.New(output).
				Level(zerolog.Level(viper.GetInt("log.level"))).
				With().
				Timestamp().
				Logger(),
		}
	})

	return log
}
