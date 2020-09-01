package config

import (
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Logger struct {
		File       string
		Level      string
		MuteStdout bool
	}
	Storage struct {
		InMemory bool
		SQLHost  string
		SQLPort  string
		SQLDbase string
		SQLUser  string
		SQLPass  string
	}
}

// Confita может быти и хороша, но она не возвращает ошибки, если не может распарсить файл в структуру. Мне не нравится такая "молчаливость".
func NewConfig(configFile string) (Config, error) {
	f, err := os.Open(configFile)
	if err != nil {
		return Config{}, err
	}
	defer f.Close()
	s, err := ioutil.ReadAll(f)
	if err != nil {
		return Config{}, err
	}
	var config Config
	_, err = toml.Decode(string(s), &config)
	return config, err
}
