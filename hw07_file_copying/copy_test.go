package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	t.Run("Не существует исходный файл", func(t *testing.T) {
		err := Copy("testdata/novalid.txt", "out.txt", 0, 0)
		require.Equal(t, err, ErrNoInFile)
	})

	t.Run("Исходный файл является папкой", func(t *testing.T) {
		err := Copy("testdata", "out.txt", 0, 0)

		require.Equal(t, err, ErrThisIsDirectory)
	})

	t.Run("Офсет больше размера исходного файла", func(t *testing.T) {
		err := Copy("testdata/input.txt", "out.txt", 8000, 0)

		require.Equal(t, err, ErrOffsetExceedsFileSize)
	})

}
