package config

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
)

type Config struct {
	Logger struct {
		File       string
		Level      string
		MuteStdout bool
	}
	Storage struct {
		In_memory bool
		Sql_host  string
		Sql_port  string
		Sql_dbase string
		Sql_user  string
		Sql_pass  string
	}
}

// Confita может быти и хороша, но она не возвращает ошибки, если не может распарсить файл в структуру. Мне не нравится такая "молчаливость"
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
