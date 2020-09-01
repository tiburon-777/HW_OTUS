package logger

import (
	"github.com/stretchr/testify/require"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/src/config"
	"io/ioutil"
	oslog "log"
	"os"
	"strings"
	"testing"
)

func TestLoggerLogic(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "log.")
	if err != nil {
		oslog.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	conf := config.Config{Logger: struct {
		File       string
		Level      string
		MuteStdout bool
	}{File: tmpfile.Name(), Level: "warn", MuteStdout: false}}
	log, err := New(conf)
	if err != nil {
		oslog.Fatal(err)
	}

	t.Run("Messages arround the level", func(t *testing.T) {
		log.Debug("debug message")
		log.Error("error message")

		res, err := ioutil.ReadAll(tmpfile)
		if err != nil {
			oslog.Fatal(err)
		}
		require.Less(t, strings.Index(string(res), "debug message"), 0)
		require.Greater(t, strings.Index(string(res), "error message"), 0)
	})
}

func TestLoggerNegative(t *testing.T) {
	t.Run("Bad file name", func(t *testing.T) {
		conf := config.Config{Logger: struct {
			File       string
			Level      string
			MuteStdout bool
		}{File: "", Level: "debug", MuteStdout: true}}
		_, err := New(conf)
		require.Error(t, err, "invalid logger config")
	})

	t.Run("Bad level", func(t *testing.T) {
		conf := config.Config{Logger: struct {
			File       string
			Level      string
			MuteStdout bool
		}{File: "asdafad", Level: "wegretryjt", MuteStdout: true}}
		_, err := New(conf)
		require.Error(t, err, "invalid logger config")
	})
}
