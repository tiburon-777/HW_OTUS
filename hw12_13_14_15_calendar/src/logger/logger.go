package logger

import (
	"errors"
	"log"
	"os"
	"strings"

	amitralog "github.com/amitrai48/logger"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/src/config"
)

type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
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

func (l Log) Debugf(format string, args ...interface{}) {
	l.Logger.Debugf(format, args)
}

func (l *Log) Infof(format string, args ...interface{}) {
	l.Logger.Infof(format, args)
}

func (l *Log) Warnf(format string, args ...interface{}) {
	l.Logger.Warnf(format, args)
}

func (l *Log) Errorf(format string, args ...interface{}) {
	l.Logger.Errorf(format, args)
}

func (l *Log) Fatalf(format string, args ...interface{}) {
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
