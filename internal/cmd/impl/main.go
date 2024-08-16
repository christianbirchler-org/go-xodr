package main

import (
	"bytes"
	"encoding/xml"
	"github.com/iancoleman/strcase"
	"go/format"
	"log"
	"os"
	"text/template"
	"unicode"
)

func main() {

	file, err := os.ReadFile("C:\\Users\\birch\\repositories\\go-xodr\\xsd_schema\\OpenDRIVE_Core.xsd")
	if err != nil {
		panic(err)
	}

	coreSchema := new(CoreSchema)
	err = xml.Unmarshal(file, coreSchema)
	if err != nil {
		panic(err)
	}

	fnMap := template.FuncMap{
		"mapPrimitives":             mapPrimitives,
		"formatStructDocumentation": formatStructDocumentation,
		"toCamel":                   strcase.ToCamel,
	}

	tmpl := template.New("elements.tmpl").Funcs(fnMap)
	tmpl, err = tmpl.ParseFiles("C:\\Users\\birch\\repositories\\go-xodr\\elements.tmpl")
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, coreSchema)
	if err != nil {
		panic(err)
	}

	fmtSrc, err := format.Source(buf.Bytes())
	if err != nil {
		fmtSrc = buf.Bytes()
	}

	err = os.WriteFile("els.go", fmtSrc, 0644)
	if err != nil {
		log.Fatal(err)
	}

}
func formatStructDocumentation(doc string) string {
	// TODO
	return "TODO: Doc formatting needs to be implemented!"
}

var primitives = map[string]string{
	"xs:double": "float64",
	"xs:int":    "int",
	"xs:string": "string",
}

func mapPrimitives(t string) string {
	goType := primitives[t]
	if goType == "" {
		return strcase.ToCamel(t)
	}
	return goType
}

func goName(n string) string {
	capitalized := make([]rune, 0, len(n))
	preWasTick := false
	for pos, char := range n {
		switch {
		case pos == 0 || preWasTick:
			bigChar := unicode.ToUpper(char)
			capitalized = append(capitalized, bigChar)
			preWasTick = false
		case unicode.IsSpace(char):
			continue
		case unicode.IsPunct(char):
			preWasTick = true
			continue
		case unicode.IsSymbol(char):
			continue
		default:
			capitalized = append(capitalized, char)
		}
	}
	return string(capitalized)
}
