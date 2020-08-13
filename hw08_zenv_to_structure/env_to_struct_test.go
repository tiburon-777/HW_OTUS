package hw08_zenv_to_structure

import (
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"testing"
)

func TestSlice2struct(t *testing.T) {
	mp := []string{"ENV1=fsfsdvsvssdfsd fsf fsf sds", "ENV2=12345", "ENV3=aefrsgrgdtgtdhn"}
	var st struct {
		ENV1 string
		ENV2 int
		ENV3 string
	}
	err := Slice2struct(mp, &st)
	require.Equal(t, struct {
		ENV1 string
		ENV2 int
		ENV3 string
	}{"fsfsdvsvssdfsd fsf fsf sds", 12345, "aefrsgrgdtgtdhn"}, st)
	require.NoError(t, err)
}

func TestEnv2struct(t *testing.T) {
	type env struct {
		ENV1 string
		ENV2 bool
		ENV3 int
	}
	exp := env{"какая-то строка", true, 1234567}
	if os.Setenv("ENV1", "какая-то строка") != nil {
		log.Fatal("не удалось установить ENV")
	}
	if os.Setenv("ENV2", "true") != nil {
		log.Fatal("не удалось установить ENV")
	}
	if os.Setenv("ENV3", "1234567") != nil {
		log.Fatal("не удалось установить ENV")
	}

	var res env
	Env2struct(&res)
	require.Equal(t, exp, res)
}
