package logger

import (
	"errors"
	"log"
	"os"
	"strings"

	amitralog "github.com/amitrai48/logger"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/config"
)

type Interface interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}

type Logger struct {
	Logger amitralog.Logger
}

var validLevel = map[string]bool{"debug": true, "info":true, "warn": true, "error": true, "fatal": true}

func New(conf config.Config) (Interface, error) {
	if conf.Logger.File == "" || !validLevel[strings.ToLower(conf.Logger.Level)] {
		return nil, errors.New("invalid logger config")
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
	return l, nil
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.Logger.Debugf(format, args)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.Logger.Infof(format, args)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.Logger.Warnf(format, args)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Logger.Errorf(format, args)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.Logger.Fatalf(format, args)
	os.Exit(2)
}
