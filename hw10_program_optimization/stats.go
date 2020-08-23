package hw10_program_optimization //nolint:golint,stylecheck

import (
	"bufio"
	"io"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	result := make(DomainStat)
	var err error
	var user User
	e := bufio.NewReader(r)
	for {
		line, _, err := e.ReadLine()
		if err == io.EOF {
			break
		}
		if err = json.Unmarshal(line, &user); err != nil {
			return nil, err
		}
		if strings.Contains(user.Email, "."+domain) {
			result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])]++
		}
	}
	return result, err
}
