package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type Environment map[string]string

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	e := make(map[string]string)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return e, err
	}
	for _, file := range files {
		if !file.IsDir() && file.Mode().IsRegular() {
			val, err := ReadFile(dir + "/" + file.Name())
			if err != nil {
				return nil, err
			}
			// TODO: Не правильно!
			key := file.Name()
			key = strings.ReplaceAll(key, `=`, ``)
			key = strings.ReplaceAll(key, `;`, ``)
			e[key] = ExtractEnv(val)
		}
	}

	return e, nil
}

func ReadFile(filePath string) (string, error) {
	f, err := os.Open(filePath)
	defer func() {
		if err := f.Close(); err != nil {
			return
		}
	}()
	if err != nil {
		return "", err
	}
	reader := bufio.NewReader(f)
	b, _, err := reader.ReadLine()
	if err != nil && err != io.EOF {
		return "", err
	}

	return string(b), nil
}

func ExtractEnv(text string) string {
	text = strings.TrimSpace(text)
	r := regexp.MustCompile("\"+")
	text = r.ReplaceAllString(text, "")
	// Удаляем все после /000
	i := strings.IndexByte(text, 0)
	if i > 0 {
		text = text[:i]
	}

	return text
}
