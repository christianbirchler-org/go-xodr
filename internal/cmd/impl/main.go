package main

import (
	"bytes"
	"flag"
	"go/format"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"text/template"
	"unicode"
)

type Element struct {
	Name string `yaml:"name"`
	//Attributes []Element `yaml:"attributes"`
}

type OpenDriveElements struct {
	Version     string    `yaml:"version"`
	Description string    `yaml:"description"`
	Elements    []Element `yaml:"elements"`
}

func main() {
	in := flag.String("in", "", "YAML file to parse")
	templateInput := flag.String("tmpl", "", "template file")
	out := flag.String("out", "", "generated Go file")
	flag.Parse()

	yamlFile, err := os.ReadFile(*in)
	if err != nil {
		log.Fatal(err)
	}

	tmplFile, err := os.ReadFile(*templateInput)
	if err != nil {
		log.Fatal(err)
	}

	var openDriveElements OpenDriveElements
	err = yaml.Unmarshal(yamlFile, &openDriveElements)
	if err != nil {
		log.Fatal(err)
	}

	elNamesToGoNames(&openDriveElements)

	tmpl, err := template.New("OpenDRIVE Elements").Parse(string(tmplFile))
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, openDriveElements)
	if err != nil {
		log.Fatal(err)
	}

	fmtSrc, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(*out, fmtSrc, 0644)
	if err != nil {
		log.Fatal(err)
	}

}

func elNamesToGoNames(od *OpenDriveElements) {
	for elPos, element := range od.Elements {
		capitalized := make([]rune, 0, len(element.Name))
		preWasTick := false
		for pos, char := range element.Name {
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
		od.Elements[elPos] = Element{
			Name: string(capitalized),
		}
	}

}
