package main

import (
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
	file, err := os.ReadFile("unique-elements.yml")
	if err != nil {
		log.Fatal(err)
	}

	var openDriveElements OpenDriveElements
	err = yaml.Unmarshal(file, &openDriveElements)
	if err != nil {
		log.Fatal(err)
	}

	elNamesToGoNames(&openDriveElements)

	tmpl, err := template.New("elements").Parse(
		`{{range .Elements}}
type {{ .Name }} struct {
}
{{end}}`)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(os.Stdout, openDriveElements)
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
