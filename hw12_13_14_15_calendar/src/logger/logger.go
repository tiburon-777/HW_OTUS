package logger

import (
	"errors"
	amitralog "github.com/amitrai48/logger"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/src/config"
	"log"
	"os"
	"strings"
)

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Fatal(msg string)
}

type Log struct {
	Logger amitralog.Logger
}

func New(conf config.Config) (Log, error) {

	if conf.Logger.File == "" || !validLevel(conf.Logger.Level) {
		return Log{}, errors.New("invalid logger config")
	}

	c := amitralog.Configuration{
		EnableConsole:     !conf.Logger.MuteStdout,
		ConsoleLevel:      amitralog.Fatal,
		ConsoleJSONFormat: false,
		EnableFile:        true,
		FileLevel:         strings.ToLower(conf.Logger.Level),
		FileJSONFormat:    true,
		FileLocation:      conf.Logger.File,
	}

	if err := amitralog.NewLogger(c, amitralog.InstanceZapLogger); err != nil {
		log.Fatalf("Could not instantiate log %s", err.Error())
	}
	l := amitralog.WithFields(amitralog.Fields{"hw": "12"})
	return Log{Logger: l}, nil
}

func (l Log) Debug(format string, args ...interface{}) {
	l.Logger.Debugf(format, args)
}

func (l *Log) Info(format string, args ...interface{}) {
	l.Logger.Infof(format, args)
}

func (l *Log) Warn(format string, args ...interface{}) {
	l.Logger.Warnf(format, args)
}

func (l *Log) Error(format string, args ...interface{}) {
	l.Logger.Errorf(format, args)
}

func (l *Log) Fatal(format string, args ...interface{}) {
	l.Logger.Fatalf(format, args)
	os.Exit(2)
}

func validLevel(level string) bool {
	l := []string{"debug", "info", "warn", "error", "fatal"}
	for _, v := range l {
		if strings.ToLower(level) == v {
			return true
		}
	}
	return false
}
