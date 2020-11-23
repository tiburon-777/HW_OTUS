package config

import (
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

func New(configFile string, str interface{}) error {
	f, err := os.Open(configFile)
	if err != nil {
		return err
	}
	defer f.Close()
	s, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	_, err = toml.Decode(string(s), str)
	return err
}
