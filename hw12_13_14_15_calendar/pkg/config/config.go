package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/BurntSushi/toml"
)

func New(configFile string, str interface{}) error {
	if configFile == "" {
		err := ApplyEnvVars(str, "APP")
		if err != nil {
			return fmt.Errorf("can't apply envvars to config :%w", err)
		}
		return nil
	}
	f, err := os.Open(configFile)
	if err != nil {
		return fmt.Errorf("can't open config file: %w", err)
	}
	defer f.Close()
	s, err := ioutil.ReadAll(f)
	if err != nil {
		return fmt.Errorf("can't read content of the config file : %w", err)
	}
	_, err = toml.Decode(string(s), str)
	if err != nil {
		return fmt.Errorf("can't parce config file : %w", err)
	}
	return nil
}

// Пришлось немного модифицировать пакет github.com/mxschmitt/golang-env-struct. В исходной реализации используется поиск по тегам, а у моих структур тегов быть не должно. Пришлось переделать функцию, чтобы она искала по именам полей.
// Тест дополнен проверкой заполнения структуры из переменных окружения.
func ApplyEnvVars(c interface{}, prefix string) error {
	return applyEnvVar(reflect.ValueOf(c), reflect.TypeOf(c), -1, prefix)
}

func applyEnvVar(v reflect.Value, t reflect.Type, counter int, prefix string) error {
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("not a pointer value")
	}
	f := reflect.StructField{}
	if counter != -1 {
		f = t.Field(counter)
	}
	v = reflect.Indirect(v)
	fName := strings.ToUpper(f.Name)
	env := os.Getenv(prefix + fName)
	if env != "" {
		switch v.Kind() {
		case reflect.Int:
			envI, err := strconv.Atoi(env)
			if err != nil {
				return fmt.Errorf("could not parse to int: %w", err)
			}
			v.SetInt(int64(envI))
		case reflect.String:
			v.SetString(env)
		case reflect.Bool:
			envB, err := strconv.ParseBool(env)
			if err != nil {
				return fmt.Errorf("could not parse bool: %w", err)
			}
			v.SetBool(envB)
		}
	}
	if v.Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			if err := applyEnvVar(v.Field(i).Addr(), v.Type(), i, prefix+fName+"_"); err != nil {
				return fmt.Errorf("could not apply env var: %w", err)
			}
		}
	}
	return nil
}
