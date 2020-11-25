package config

import (
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

func New(configFile string, str interface{}) error {
	if configFile == "" {
		return ApplyEnvVars(str, "APP")
	}
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
	if err != nil {
		return err
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
		return errors.New("not a pointer value")
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
				return errors.Wrap(err, "could not parse to int")
			}
			v.SetInt(int64(envI))
		case reflect.String:
			v.SetString(env)
		case reflect.Bool:
			envB, err := strconv.ParseBool(env)
			if err != nil {
				return errors.Wrap(err, "could not parse bool")
			}
			v.SetBool(envB)
		}
	}
	if v.Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			if err := applyEnvVar(v.Field(i).Addr(), v.Type(), i, prefix+fName+"_"); err != nil {
				return errors.Wrap(err, "could not apply env var")
			}
		}
	}
	return nil
}
