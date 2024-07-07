package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Element struct {
	Name       string    `yaml:"name"`
	Attributes []Element `yaml:"attributes"`
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
	fmt.Println(openDriveElements)
}
