package main

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"log"
	"os"
	"reflect"
	"strings"
	"text/template"
)

type TemplateVars struct {
	StructName     string
	FieldName      string
	FieldType      string
	ValidatorValue string
	IsSlice        bool
}

const str string = "string"

func main() {
	buf := new(bytes.Buffer)
	buf.WriteString(validatorHeader)
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, os.Args[1], nil, parser.ParseComments)
	//node, err := parser.ParseFile(fset, "models/models.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatal("не удалось открыть файл", err.Error())
	}
	// Перебираем корневые узлы
	for _, f := range node.Decls {
		gd, ok := f.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec := range gd.Specs {
			// Ищем структуры
			t, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			s, ok := t.Type.(*ast.StructType)
			if !ok {
				continue
			}
			// Если все проверки пройдены, это структура
			// Итерируем ее поля
			tmpBuf := new(bytes.Buffer)
			tmplV := TemplateVars{StructName: t.Name.Name}
			z := template.Must(template.New("validatorFunctionHeader").Parse(validatorFunctionHeader))
			if z.Execute(tmpBuf, tmplV) != nil {
				log.Fatal("ошибка сборки шаблона:")
			}
			ok, err = iterStructFields("", s, tmpBuf)

			if err != nil {
				log.Fatal("ошибка в ходе перебора полей структуры:", err.Error())
			}
			tmpBuf.Write([]byte(validatorFunctionFooter))
			if ok {
				if _, err := tmpBuf.WriteTo(buf); err != nil {
					log.Fatal("ошибка перекладывания в буфер:", err.Error())
				}
			}
		}
	}
	buf.WriteString(validatorFunctions)
	f, err := os.Create(strings.Split(os.Args[1], ".")[0] + "_validation_generated.go")
	if err != nil {
		log.Fatal("не удалось создать/открыть файл:", err.Error())
	}
	defer func() {
		if f.Close() != nil {
			panic("Не удается захлопнуть файл!")
		}
	}()
	_, err = buf.WriteTo(f)
	if err != nil {
		log.Fatal("не удалось записать буфер в файл:", err.Error())
	}
}

func iterStructFields(name string, s *ast.StructType, buf io.Writer) (bool, error) {
	var isValidated bool
	for _, field := range s.Fields.List {
		// Рекурсивный вызов, в случае если поле является структурой
		switch field.Type.(type) {
		case *ast.StructType:
			if _, err := iterStructFields(field.Names[0].Name+".", field.Type.(*ast.StructType), buf); err != nil {
				return false, err
			}
		}
		if len(field.Names) == 0 {
			continue
		}
		// Достаем тэг поля
		tag, ok := getTagString(field, "validate")
		if !ok {
			continue
		}
		isValidated = true
		//Ищем комбинации
		for _, comb := range strings.Split(tag, "|") {
			k := strings.Split(comb, ":")

			isSlice, fieldType := getFieldType(field)
			tmplV := TemplateVars{FieldName: name + field.Names[0].Name, FieldType: fieldType, ValidatorValue: k[1], IsSlice: isSlice}

			z := &template.Template{}
			switch k[0] {
			case "min":
				z = template.Must(template.New("validatorMin").Parse(validatorMin))
			case "max":
				z = template.Must(template.New("validatorMax").Parse(validatorMax))
			case "in":
				switch fieldType {
				case "int":
					z = template.Must(template.New("validatorInInt").Parse(validatorInInt))
				case str:
					z = template.Must(template.New("validatorInStr").Parse(validatorInStr))
				}
			case "len":
				z = template.Must(template.New("validatorLen").Parse(validatorLen))
			case "regexp":
				z = template.Must(template.New("validatorRegexp").Parse(validatorRegexp))
			default:
				log.Fatal("Неизвестный параметр тега validate")
			}
			err := z.Execute(buf, tmplV)
			if err != nil {
				return false, err
			}
		}
	}
	return isValidated, nil
}

func getTagString(f *ast.Field, tag string) (string, bool) {
	if f.Tag == nil {
		return "", false
	}
	t := reflect.StructTag(f.Tag.Value[1 : len(f.Tag.Value)-1])
	v := t.Get(tag)
	if v == "" {
		return "", false
	}
	return v, true
}

func getFieldType(field *ast.Field) (bool, string) {
	var isSlice bool
	var fieldType string
	// Эту конструкцию не пропускал линтер gocritic, под предлогом "typeSwitchVar: 2 cases can benefit from type switch with assignment". Я не вижу тут возможности срабатывания обеих веток. тип ast.Expr не может привестись и к *ast.Ident и к *ast.ArrayType одновременно. Пришлось, отключить старикашку критика.
	switch field.Type.(type) {
	case *ast.Ident:
		isSlice = false
		fieldType = field.Type.(*ast.Ident).Name
		if field.Type.(*ast.Ident).Obj != nil {
			t, ok := field.Type.(*ast.Ident).Obj.Decl.(*ast.TypeSpec)
			if !ok {
				return false, ""
			}
			s, ok := t.Type.(*ast.Ident)
			if !ok {
				return false, ""
			}
			fieldType = s.Name
		}
	case *ast.ArrayType:
		isSlice = true
		fieldType = field.Type.(*ast.ArrayType).Elt.(*ast.Ident).Name
	}
	switch {
	case strings.Contains(fieldType, "int"):
		return isSlice, "int"
	case strings.Contains(fieldType, "float"):
		return isSlice, "float"
	case fieldType == str:
		return isSlice, str
	}
	return false, ""
}
