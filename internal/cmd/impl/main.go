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

type Attribute struct {
	Name        string `yaml:"name"`
	Type        string `yaml:"type"`
	Use         string `yaml:"use"`
	Description string `yaml:"description"`
}

type Child struct {
	Name string `yaml:"name"`
}

type Element struct {
	Name       string      `yaml:"name"`
	XmlTag     string      `yaml:"xmlTag"`
	Children   []Child     `yaml:"children"`
	Attributes []Attribute `yaml:"attributes"`
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

	for elPos, element := range openDriveElements.Elements {
		children := make([]Child, len(element.Children))
		attributes := make([]Attribute, len(element.Attributes))
		// create go names for children structs
		for childPos, child := range element.Children {
			children[childPos] = Child{
				Name: goName(child.Name),
			}
		}
		for attrPos, attr := range element.Attributes {
			attributes[attrPos] = Attribute{
				Name:        goName(attr.Name),
				Type:        attr.Type,
				Use:         attr.Use,
				Description: attr.Description,
			}
		}
		openDriveElements.Elements[elPos] = Element{
			Name:       goName(element.Name),
			Children:   children,
			Attributes: attributes,
		}
	}

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
