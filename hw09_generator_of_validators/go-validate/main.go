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
	//node, err := parser.ParseFile(fset, os.Args[1], nil, parser.ParseComments)
	node, err := parser.ParseFile(fset, "models/models.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatal("не удалось открыть файл", err.Error())
	}
	// Перебираем ноды
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

			tmplV := TemplateVars{StructName: t.Name.Name}
			z := template.Must(template.New("validatorFunctionHeader").Parse(validatorFunctionHeader))
			if z.Execute(buf, tmplV) != nil {
				log.Fatal("ошибка сборки шаблона:", err)
			}
			err = iterStructFields(s, buf)
			if err != nil {
				log.Fatal("ошибка в ходе перебора полей структуры:", err.Error())
			}
			buf.WriteString(validatorFunctionFooter)
		}
	}
	buf.WriteString(validatorFunctions)
	f, err := os.Create("models/models_validation_generated.go")
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
	// Тут нужно открыть файл, записать в него байтбуфер и закрыть файд
}

func iterStructFields(s *ast.StructType, buf io.Writer) error {
	for _, field := range s.Fields.List {
		// Рекурсивный вызов, в случае если поле является структурой
		//switch field.Type.(type) {
		//case *ast.StructType:
		//	if err:=iterStructFields(field.Type.(*ast.StructType), buf); err!=nil { log.Fatal("Структура не распарсилась в рекурсии:", err.Error()) }
		//}
		if len(field.Names) == 0 {
			continue
		}
		// Достаем тэг поля
		if field.Tag == nil {
			continue
		}
		tag, ok := getTagString(field, "validate")
		if !ok {
			continue
		}
		//Ищем комбинации
		for _, comb := range strings.Split(tag, "|") {
			k := strings.Split(comb, ":")

			isSlice, fieldType := getFieldType(field)
			tmplV := TemplateVars{FieldName: field.Names[0].Name, FieldType: fieldType, ValidatorValue: k[1], IsSlice: isSlice}

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
				return err
			}
		}
	}
	return nil
}

func getTagString(f *ast.Field, tag string) (string, bool) {
	t := reflect.StructTag(f.Tag.Value[1 : len(f.Tag.Value)-1])
	v := t.Get(tag)
	if v == "" {
		return "", false
	}
	return v, true
}

func getFieldType(field *ast.Field) (bool, string) {
	var fieldSlice bool
	var fieldType string
	switch field.Type.(type) {
	case *ast.Ident:
		fieldSlice = false
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
		fieldSlice = true
		fieldType = field.Type.(*ast.ArrayType).Elt.(*ast.Ident).Name
	}
	switch {
	case strings.Contains(fieldType, "int"):
		return fieldSlice, "int"
	case strings.Contains(fieldType, "float"):
		return fieldSlice, "float"
	case fieldType == str:
		return fieldSlice, str
	}
	return false, ""
}
