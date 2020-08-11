package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRunCmd(t *testing.T) {
	t.Run("Команда не существует", func(t *testing.T) {
		code := RunCmd([]string{"sffdsfsvs"}, Environment{"E1": "val1"})
		require.Equal(t, code, -1)
	})
	t.Run("Команда выполнилась", func(t *testing.T) {
		code := RunCmd([]string{"ls"}, Environment{"E1": "val1"})
		require.Equal(t, code, 0)
	})
}
