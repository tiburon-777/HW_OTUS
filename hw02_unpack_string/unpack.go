package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

const slash = `\`

func Unpack(str string) (string, error) {
	var b strings.Builder
	temp := []rune(str)
	for i := 0; i < len(temp); i++ {
		if unicode.IsLetter(temp[i]) {
			// Ловим паттерн [letter][digit]
			if i < len(temp)-1 && unicode.IsDigit(temp[i+1]) {
				c, err := strconv.Atoi(string(temp[i+1]))
				if err != nil {
					return "", err
				}
				b.WriteString(strings.Repeat(string(temp[i]), c))
				i++
				// Ловим паттерн [letter]...
			} else {
				b.WriteRune(temp[i])
			}
		} else if string(temp[i]) == slash {
			// Ловим паттерн [\][letter][digit]
			if i < len(temp)-2 && unicode.IsLetter(temp[i+1]) && unicode.IsDigit(temp[i+2]) {
				c, err := strconv.Atoi(string(temp[i+2]))
				if err != nil {
					return "", err
				}
				b.WriteString(strings.Repeat(slash+string(temp[i+1]), c))
				i += 2
				// Ловим паттерн [\][digit или \][digit]
			} else if i < len(temp)-2 && (unicode.IsDigit(temp[i+1]) || string(temp[i+1]) == slash) && unicode.IsDigit(temp[i+2]) {
				c, err := strconv.Atoi(string(temp[i+2]))
				if err != nil {
					return "", err
				}
				b.WriteString(strings.Repeat(string(temp[i+1]), c))
				i += 2
				// Ловим паттерн [\][digit или \]
			} else if i < len(temp)-1 && (unicode.IsDigit(temp[i+1]) || string(temp[i+1]) == slash) {
				b.WriteRune(temp[i+1])
				i++
			} else {
				return "", ErrInvalidString
			}
		} else {
			return "", ErrInvalidString
		}
	}
	return b.String(), nil
}
