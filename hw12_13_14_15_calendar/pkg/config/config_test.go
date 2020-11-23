package config

import (
	"github.com/stretchr/testify/require"
	"io/ioutil"
	oslog "log"
	"os"
	"testing"
)

var _ = func() bool {
	testing.Init()
	return true
}()

func TestNewConfig(t *testing.T) {

	badfile, err := ioutil.TempFile("", "conf.")
	if err != nil {
		oslog.Fatal(err)
	}
	defer os.Remove(badfile.Name())
	badfile.WriteString(`aefSD
sadfg
RFABND FYGUMG
V`)
	badfile.Sync()

	goodfile, err := ioutil.TempFile("", "conf.")
	if err != nil {
		oslog.Fatal(err)
	}
	defer os.Remove(goodfile.Name())
	goodfile.WriteString(`[storage]
inMemory = true
SQLHost = "localhost"`)
	goodfile.Sync()

	t.Run("No such file", func(t *testing.T) {
		var c Calendar
		e := New("adfergdth", &c)
		require.Equal(t, Calendar{}, c)
		require.Error(t, e)
	})

	t.Run("Bad file", func(t *testing.T) {
		var c Calendar
		e := New(badfile.Name(), &c)
		require.Equal(t, Calendar{}, c)
		require.Error(t, e)
	})

	t.Run("TOML reading", func(t *testing.T) {
		var c Calendar
		e := New(goodfile.Name(), &c)
		require.Equal(t, true, c.Storage.InMemory)
		require.Equal(t, "localhost", c.Storage.SQLHost)
		require.NoError(t, e)
	})

}
