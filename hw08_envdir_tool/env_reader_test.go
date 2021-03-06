package main

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestReadDir(t *testing.T) {
	t.Run("Нет файлов в директории", func(t *testing.T) {
		os.RemoveAll("testdata/env/no")
		if err := os.Mkdir("testdata/env/no", os.ModePerm); err != nil {
			return
		}
		env, err := ReadDir("testdata/env/no")
		require.Equal(t, env, Environment{})
		require.NoError(t, err)
		os.RemoveAll("testdata/env/no")
	})

	t.Run("Проверка = и ; в именах", func(t *testing.T) {
		if err := os.RemoveAll("testdata/env/no"); err != nil {
			return
		}
		if err := os.Mkdir("testdata/env/no", os.ModePerm); err != nil {
			return
		}
		f1, err := os.Create("testdata/env/no/TES=T1")
		if err != nil {
			return
		}
		f2, err := os.Create("testdata/env/no/TES;T2")
		if err != nil {
			return
		}
		env, err := ReadDir("testdata/env/no")
		if err != nil {
			return
		}
		require.Equal(t, env, Environment{})
		require.NoError(t, err)
		if err = f1.Close(); err != nil {
			return
		}
		if err = f2.Close(); err != nil {
			return
		}
		if err = os.RemoveAll("testdata/env/no"); err != nil {
			return
		}
	})
}

func TestReadFile(t *testing.T) {
	t.Run("Файл не существует", func(t *testing.T) {
		str, err := ReadFile("dsvfsdfdfdv")
		require.Equal(t, str, "")
		require.Error(t, err)
	})
	t.Run("Файл пуст", func(t *testing.T) {
		str, err := ReadFile("testdata/env/UNSET")
		require.Equal(t, str, "")
		require.NoError(t, err)
	})
}

func TestExtractEnv(t *testing.T) {

	t.Run("Zero escaping", func(t *testing.T) {
		require.Equal(t, "zero_escape\nwith new line", ExtractEnv("zero_escape\x00with new line"))
	})

	t.Run("Quotes", func(t *testing.T) {
		require.Equal(t, `"quotes"`, ExtractEnv("\"quotes\""))
	})

	t.Run("Pre spacing", func(t *testing.T) {
		require.Equal(t, "     	pre_spased", ExtractEnv("     	pre_spased"))
	})

	t.Run("Post spacing", func(t *testing.T) {
		require.Equal(t, "post_spased", ExtractEnv("post_spased      "))
	})

	t.Run("Multy spacing", func(t *testing.T) {
		require.Equal(t, "  	 multy spased", ExtractEnv("  	 multy spased      "))
	})
}
