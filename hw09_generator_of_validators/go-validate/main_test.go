package main

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
	"testing"
)

const TestModel string = `
package testmodel
type (
	TestSingle struct {
		TstIntMinMax     int ` + "`" + `validate:"min:18|max:50"` + "`" + `
		TstIntIn         int ` + "`" + `validate:"in:45,21,57,12"` + "`" + `
		TstStringLen     string ` + "`" + `condidate:"len:36"` + "`" + `
		TstStringRegexp  string
		TstStringIn      string ` + "`" + `validate:"in:admin,stuff"` + "`" + `
	}
)
type (
	TestSlice struct {
		TstIntMinMax     []int ` + "`" + `validate:"min:18|max:50"` + "`" + `
		TstIntIn         []int ` + "`" + `validate:"in:45,21,57,12"` + "`" + `
		TstStringLen     []string ` + "`" + `json:"len:36"` + "`" + `
		TstStringRegexp  []string ` + "`" + `validate:"regexp:^\\w+@\\w+\\.\\w+$"` + "`" + `
		TstStringIn      []string ` + "`" + `xml:"in:admin,stuff"` + "`" + `
	}
)`
const FineStructure string = `
package testmodel
type (
	TestSingle struct {
		TstIntMinMax     int ` + "`" + `validate:"min:18|max:50"` + "`" + `
		TstIntIn         int ` + "`" + `validate:"in:45,21,57,12"` + "`" + `
		TstStringLen     string ` + "`" + `validate:"len:36"` + "`" + `
		TstStringIn      string ` + "`" + `validate:"in:admin,stuff"` + "`" + `
	}
)`

const WrongStructure string = `
package testmodel
type (
	TestSingle struct {
		TstIntMinMax     int ` + "`" + `validate:"min:WRONG|max:#$&^%"` + "`" + `
		TstIntIn         int ` + "`" + `validate:"in:45,STRING,57,12.5"` + "`" + `
		TstStringLen     string ` + "`" + `validate:"len:"` + "`" + `
		TstStringIn      string ` + "`" + `validate:"in:admin,stuff,,,"` + "`" + `
	}
)`

const FineStructureExpcted string = "\n\tif ok, err = valMin([]int{int(u.TstIntMinMax)},18); !ok { res=append(res,ValidationError{\"TstIntMinMax\",fmt.Errorf(\"min:18\")})}\n\tif ok, err = valMax([]int{int(u.TstIntMinMax)},50); !ok { res=append(res,ValidationError{\"TstIntMinMax\",fmt.Errorf(\"max:50\")})}\n\tif ok, err = valInInt([]int{int(u.TstIntIn)},`45,21,57,12`); !ok { res=append(res,ValidationError{\"TstIntIn\",fmt.Errorf(\"in:45,21,57,12\")})}\n\tif ok, err = valLen([]string{string(u.TstStringLen)},36); !ok { res=append(res,ValidationError{\"TstStringLen\",fmt.Errorf(\"len:36\")})}\n\tif ok, err = valInString([]string{string(u.TstStringIn)},`admin,stuff`); !ok { res=append(res,ValidationError{\"TstStringIn\",fmt.Errorf(`in:admin,stuff`)})}"

const WrongStructureExpcted string = "\n\tif ok, err = valMin([]int{int(u.TstIntMinMax)},WRONG); !ok { res=append(res,ValidationError{\"TstIntMinMax\",fmt.Errorf(\"min:WRONG\")})}\n\tif ok, err = valMax([]int{int(u.TstIntMinMax)},#$&^%); !ok { res=append(res,ValidationError{\"TstIntMinMax\",fmt.Errorf(\"max:#$&^%\")})}\n\tif ok, err = valInInt([]int{int(u.TstIntIn)},`45,STRING,57,12.5`); !ok { res=append(res,ValidationError{\"TstIntIn\",fmt.Errorf(\"in:45,STRING,57,12.5\")})}\n\tif ok, err = valLen([]string{string(u.TstStringLen)},); !ok { res=append(res,ValidationError{\"TstStringLen\",fmt.Errorf(\"len:\")})}\n\tif ok, err = valInString([]string{string(u.TstStringIn)},`admin,stuff,,,`); !ok { res=append(res,ValidationError{\"TstStringIn\",fmt.Errorf(`in:admin,stuff,,,`)})}"

func TestGetTagString(T *testing.T) {
	var tags = []string{"min:18|max:50", "in:45,21,57,12", "", "", "in:admin,stuff", "min:18|max:50", "in:45,21,57,12", "", `regexp:^\w+@\w+\.\w+$`, ""}
	var oks = []bool{true, true, false, false, true, true, true, false, true, false}
	fset := token.NewFileSet()
	node, _ := parser.ParseFile(fset, "", []byte(TestModel), parser.ParseComments)
	for _, f := range node.Decls {
		gd, ok := f.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec := range gd.Specs {
			t, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			s, ok := t.Type.(*ast.StructType)
			if !ok {
				continue
			}
			for _, field := range s.Fields.List {
				T.Run("Single int field "+field.Names[0].Name, func(t *testing.T) {
					tagString, ok := getTagString(field, "validate")
					require.Equal(t, oks[0], ok)
					require.Equal(t, tags[0], tagString)
					tags = tags[1:]
					oks = oks[1:]
				})
			}
		}
	}
}

func TestGetFieldType(T *testing.T) {
	fset := token.NewFileSet()
	node, _ := parser.ParseFile(fset, "", []byte(TestModel), parser.ParseComments)
	for _, f := range node.Decls {
		gd, ok := f.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec := range gd.Specs {
			t, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			s, ok := t.Type.(*ast.StructType)
			if !ok {
				continue
			}
			if t.Name.Name == "TestSingle" {
				for _, field := range s.Fields.List {
					if strings.Index(field.Names[0].Name, "Int") > 0 {
						T.Run("Single int field "+field.Names[0].Name, func(t *testing.T) {
							isSlice, fType := getFieldType(field)
							require.Equal(t, false, isSlice)
							require.Equal(t, "int", fType)
						})
					}
					if strings.Index(field.Names[0].Name, "String") > 0 {
						T.Run("Single string field "+field.Names[0].Name, func(t *testing.T) {
							isSlice, fType := getFieldType(field)
							require.Equal(t, false, isSlice)
							require.Equal(t, "string", fType)
						})
					}
				}
			}
			if t.Name.Name == "TestSlice" {
				for _, field := range s.Fields.List {
					if strings.Index(field.Names[0].Name, "Int") > 0 {
						T.Run("Slice int field "+field.Names[0].Name, func(t *testing.T) {
							isSlice, fType := getFieldType(field)
							require.Equal(t, true, isSlice)
							require.Equal(t, "int", fType)
						})
					}
					if strings.Index(field.Names[0].Name, "String") > 0 {
						T.Run("Slice string field "+field.Names[0].Name, func(t *testing.T) {
							isSlice, fType := getFieldType(field)
							require.Equal(t, true, isSlice)
							require.Equal(t, "string", fType)
						})
					}
				}
			}
		}
	}
}

func TestIterStructFields(T *testing.T) {
	fset := token.NewFileSet()
	node, _ := parser.ParseFile(fset, "", []byte(FineStructure), parser.ParseComments)
	for _, f := range node.Decls {
		gd, ok := f.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec := range gd.Specs {
			t, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			s, ok := t.Type.(*ast.StructType)
			if !ok {
				continue
			}
			T.Run("Fine structure", func(t *testing.T) {
				buf := bytes.NewBufferString("")
				err := iterStructFields(s, buf)
				require.Equal(t, FineStructureExpcted, buf.String())
				require.NoError(t, err)
			})
		}
	}

	fset1 := token.NewFileSet()
	node1, _ := parser.ParseFile(fset1, "", []byte(WrongStructure), parser.ParseComments)
	for _, f1 := range node1.Decls {
		gd1, ok := f1.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec1 := range gd1.Specs {
			t, ok := spec1.(*ast.TypeSpec)
			if !ok {
				continue
			}
			s1, ok := t.Type.(*ast.StructType)
			if !ok {
				continue
			}
			T.Run("Wrong structure", func(t *testing.T) {
				buf1 := bytes.NewBufferString("")
				err := iterStructFields(s1, buf1)
				require.Equal(t, WrongStructureExpcted, buf1.String())
				require.NoError(t, err)
			})
		}
	}

}
