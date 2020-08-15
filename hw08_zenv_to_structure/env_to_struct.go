package hw08_zenv_to_structure //nolint:golint,stylecheck
import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func Slice2struct(mp []string, str interface{}) error {
	v := reflect.ValueOf(str)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("%T is not a pointer", str)
	}
	v = v.Elem()
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("%T is not a pointer to struct", str)
	}

	for _, e := range mp {
		st := strings.Split(e, "=")
		var key string = st[0]
		var value interface{} = st[1]
		f := v.FieldByName(key)
		if f.IsValid() && f.CanSet() {
			switch v := f.Interface().(type) {
			case int:
				val, err := strconv.Atoi(value.(string))
				if err != nil {
					return err
				}
				x := int64(val)
				if !f.OverflowInt(x) {
					f.SetInt(x)
				}
			case string:
				f.SetString(value.(string))
			case bool:
				f.SetBool(value.(string) == "true")
			case []interface{}:
				input := reflect.ValueOf(value)
				f.Set(input)
			default:
				return fmt.Errorf("i don't know how to parse type %T", v)
			}
		}
	}
	return nil
}

func Env2struct(s interface{}) {
	env := os.Environ()
	if err := Slice2struct(env, s); err != nil {
		log.Fatal("ошибка", err.Error())
	}
}
