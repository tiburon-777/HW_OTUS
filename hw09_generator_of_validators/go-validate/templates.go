package main

const validatorHeader string = `// Code generated by cool go-validate tool; DO NOT EDIT.
package models

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type ValidationError struct {
	Field	string
	Err	error
}
`

const validatorFunctionHeader string = `
func (u {{.StructName}}) Validate() ([]ValidationError, error) {
	res := []ValidationError{}
	var err error
	var ok bool`

const validatorFunctionFooter string = `
	log.Println(ok)
	return res, err
}
`

const validatorLen string = `
	if ok, err = valLen({{if .IsSlice}}u.{{.FieldName}}{{- else}}[]string{string(u.{{.FieldName}})}{{- end}},{{.ValidatorValue}}); !ok { res=append(res,ValidationError{"{{.FieldName}}",fmt.Errorf("len:{{.ValidatorValue}}")})}`
const validatorMin string = `
	if ok, err = valMin({{if .IsSlice}}u.{{.FieldName}}{{- else}}[]int{int(u.{{.FieldName}})}{{- end}},{{.ValidatorValue}}); !ok { res=append(res,ValidationError{"{{.FieldName}}",fmt.Errorf("min:{{.ValidatorValue}}")})}`
const validatorMax string = `
	if ok, err = valMax({{if .IsSlice}}u.{{.FieldName}}{{- else}}[]int{int(u.{{.FieldName}})}{{- end}},{{.ValidatorValue}}); !ok { res=append(res,ValidationError{"{{.FieldName}}",fmt.Errorf("max:{{.ValidatorValue}}")})}`
const validatorRegexp string = `
	if ok, err = valRegexp({{if .IsSlice}}u.{{.FieldName}}{{- else}}[]string{string(u.{{.FieldName}})}{{- end}},` + "`{{.ValidatorValue}}`" + `); !ok { res=append(res,ValidationError{"{{.FieldName}}",fmt.Errorf(` + "`regexp:{{.ValidatorValue}}`" + `)})}`
const validatorInStr string = `
	if ok, err = valInString({{if .IsSlice}}u.{{.FieldName}}{{- else}}[]string{string(u.{{.FieldName}})}{{- end}},` + "`{{.ValidatorValue}}`" + `); !ok { res=append(res,ValidationError{"{{.FieldName}}",fmt.Errorf(` + "`in:{{.ValidatorValue}}`" + `)})}`
const validatorInInt string = `
	if ok, err = valInInt({{if .IsSlice}}u.{{.FieldName}}{{- else}}[]int{int(u.{{.FieldName}})}{{- end}},` + "`{{.ValidatorValue}}`" + `); !ok { res=append(res,ValidationError{"{{.FieldName}}",fmt.Errorf("in:{{.ValidatorValue}}")})}`

const validatorFunctions string = `

func valLen(s []string, n int) (bool,error) {
	res := true
	for _,v :=range s {
		if len(v)!=n {
			res=false
			break
		}
	}
	return res,nil
}

func valMin(i []int, n int) (bool,error) {
	res := true
	for _,v :=range i {
		if v<n {
			res=false
			break
		}
	}
	return res,nil
}

func valMax(i []int, n int) (bool,error) {
	res := true
	for _,v :=range i {
		if v>n {
			res=false
			break
		}
	}
	return res,nil
}

func valRegexp(s []string, r string) (bool,error) {
	res := true
	var err error
	rg := regexp.MustCompile(r)
	for _,v :=range s {
		if !rg.MatchString(v) {
			res=false
			break
		}
	}
	return res,err
}

func valInString(s []string, r string) (bool,error) {
	i := false
	for _,k :=range s {
		i = false
		for _, v := range strings.Split(r, ",") {
			if k == v {
				i = true
			}
		}
		if !i { break }
	}
	return i, nil
}

func valInInt(s []int, r string) (bool,error) {
	i := false
	var err error
	var m int
	for _,k :=range s {
		i = false
		for _, v := range strings.Split(r, ",") {
			m,err=strconv.Atoi(v)
			if k == m { i = true }
		}
		if !i { break }
	}
	return i, err
}
`
