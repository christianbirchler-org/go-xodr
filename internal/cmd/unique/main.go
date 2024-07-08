package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
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

	file, err := os.ReadFile("elements.yml")
	if err != nil {
		log.Fatal(err)
	}

	var openDriveElements OpenDriveElements
	err = yaml.Unmarshal(file, &openDriveElements)
	if err != nil {
		log.Fatal(err)
	}

	hashMap := make(map[string]bool, len(openDriveElements.Elements))
	cnt := 0
	var oe OpenDriveElements
	oe.Version = openDriveElements.Version
	oe.Description = openDriveElements.Description
	for _, el := range openDriveElements.Elements {
		if !hashMap[el.Name] {
			cnt += 1
			hashMap[el.Name] = true
			oe.Elements = append(oe.Elements, el)
		}
	}
	fmt.Printf("%d unique elements", cnt)
	fmt.Println(oe.Elements)

	out, err := yaml.Marshal(&oe)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("unique-elements.yml", out, 0644)
	if err != nil {
		log.Fatal(err)
	}

}
