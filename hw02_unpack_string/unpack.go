package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

const slash = '\\'

func Unpack(str string) (string, error) {
	var b strings.Builder
	temp := []rune(str)
	for i := 0; i < len(temp); i++ {
		switch {
		// Ловим паттерн [\][letter][digit]
		case temp[i] == slash && i < len(temp)-2 && unicode.IsLetter(temp[i+1]) && unicode.IsDigit(temp[i+2]):
			c, _ := strconv.Atoi(string(temp[i+2]))
			b.WriteString(strings.Repeat(string(slash)+string(temp[i+1]), c))
			i += 2
		// Ловим паттерн [\][digit или \][digit]
		case temp[i] == slash && i < len(temp)-2 && (unicode.IsDigit(temp[i+1]) || temp[i+1] == slash) && unicode.IsDigit(temp[i+2]):
			c, _ := strconv.Atoi(string(temp[i+2]))
			b.WriteString(strings.Repeat(string(temp[i+1]), c))
			i += 2
		// Ловим паттерн [\][digit или \]
		case temp[i] == slash && i < len(temp)-1 && (unicode.IsDigit(temp[i+1]) || temp[i+1] == slash):
			b.WriteRune(temp[i+1])
			i++
		case unicode.IsLetter(temp[i]) && i < len(temp)-1 && unicode.IsDigit(temp[i+1]):
			c, _ := strconv.Atoi(string(temp[i+1]))
			b.WriteString(strings.Repeat(string(temp[i]), c))
			i++
			// Ловим паттерн [letter]...
		case unicode.IsLetter(temp[i]):
			b.WriteRune(temp[i])
		default:
			return "", ErrInvalidString
		}
	}
	return b.String(), nil
}
