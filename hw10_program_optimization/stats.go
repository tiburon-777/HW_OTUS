package hw10_program_optimization //nolint:golint,stylecheck

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
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

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	u, err := getUsers(r)
	if err != nil {
		return nil, fmt.Errorf("get users error: %s", err)
	}
	return countDomains(u, domain)
}

type Users []User

func getUsers(r io.Reader) (Users, error) {
	result := make([]User, 10000)
	content, err := ioutil.ReadAll(r)
	if err != nil {
		return result, err
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		var user User
		if err = json.Unmarshal([]byte(line), &user); err != nil {
			return result, err
		}
		result = append(result,user)
	}
	return result, err
}

func countDomains(u Users, domain string) (DomainStat, error) {
	result := make(DomainStat)
	i:=0
	for _, user := range u {
		if strings.Contains(user.Email,"."+domain) {
			result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])]++
			i++
		}
	}
	return result, nil
}
