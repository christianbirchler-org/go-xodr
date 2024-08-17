package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"go/format"
	"log"
	"os"
	"text/template"
	"unicode"

	"github.com/iancoleman/strcase"
)

func main() {
	xsdPath := flag.String("xsd", ".", "path to xsd_schema directory")
	flag.Parse()

	coreFile, err := os.ReadFile(*xsdPath+"/OpenDRIVE_Core.xsd")
	if err != nil {
		panic(err)
	}
	coreSchema := new(CoreSchema)
	err = xml.Unmarshal(coreFile, coreSchema)
	if err != nil {
		panic(err)
	}
	junctionFile, err := os.ReadFile(*xsdPath+"/OpenDRIVE_Junction.xsd")
	if err != nil {
		panic(err)
	}
	junctionSchema := new(JunctionSchema)
	err = xml.Unmarshal(junctionFile, junctionSchema)
	if err != nil {
		panic(err)
	}
	laneFile, err := os.ReadFile(*xsdPath+"/OpenDRIVE_Lane.xsd")
	if err != nil {
		panic(err)
	}
	laneSchema := new(LaneSchema)
	err = xml.Unmarshal(laneFile, laneSchema)
	if err != nil {
		panic(err)
	}
	objectFile, err := os.ReadFile(*xsdPath+"/OpenDRIVE_Object.xsd")
	if err != nil {
		panic(err)
	}
	objectSchema := new(ObjectSchema)
	err = xml.Unmarshal(objectFile, objectSchema)
	if err != nil {
		panic(err)
	}
	railroadFile, err := os.ReadFile(*xsdPath+"/OpenDRIVE_Railroad.xsd")
	if err != nil {
		panic(err)
	}
	railroadSchema := new(RailroadSchema)
	err = xml.Unmarshal(railroadFile, railroadSchema)
	if err != nil {
		panic(err)
	}
	roadFile, err := os.ReadFile(*xsdPath+"/OpenDRIVE_Road.xsd")
	if err != nil {
		panic(err)
	}
	roadSchema := new(RoadSchema)
	err = xml.Unmarshal(roadFile, roadSchema)
	if err != nil {
		panic(err)
	}
	signalFile, err := os.ReadFile(*xsdPath+"/OpenDRIVE_Signal.xsd")
	if err != nil {
		panic(err)
	}
	signalSchema := new(SignalSchema)
	err = xml.Unmarshal(signalFile, signalSchema)
	if err != nil {
		panic(err)
	}

	fnMap := template.FuncMap{
		"mapPrimitives":                        mapPrimitives,
		"formatStructDocumentation":            formatStructDocumentation,
		"toCamel":                              strcase.ToCamel,
		"canCreateTRoadSignalsSignalReference": canCreateTRoadSignalsSignalReference,
		"removeTypePrefix":                     removeTypePrefix,
	}

	generate(fnMap, coreSchema, "core")
	generate(fnMap, junctionSchema, "junction")
	generate(fnMap, laneSchema, "lane")
	generate(fnMap, objectSchema, "object")
	generate(fnMap, railroadSchema, "railroad")
	generate(fnMap, roadSchema, "road")
	generate(fnMap, signalSchema, "signal")
}

func removeTypePrefix(t string) string {
	if len(t) < 3 {
		return t
	}
	firstLetter := t[0:1]
	suffix := t[2:]
	switch firstLetter {
	case "t_":
		return suffix
	case "T_":
		return suffix
	default:
		return t

	}
}

var funcCnt = 0

func canCreateTRoadSignalsSignalReference(camelCaseName string) bool {
	if camelCaseName != "TRoadSignalsSignalReference" {
		return true
	}
	funcCnt++
	return funcCnt < 2
}

func generate(fnMap template.FuncMap, schema any, name string) {
	tmpl := template.New(name + ".tmpl").Funcs(fnMap)
	tmpl, err := tmpl.ParseFiles("/Users/christian/repositories/go-xodr/templates/" + name + ".tmpl")
	if err != nil {
		panic(err)
	}
	signalBuf := new(bytes.Buffer)
	err = tmpl.Execute(signalBuf, schema)
	if err != nil {
		panic(err)
	}
	fmtSrc, err := format.Source(signalBuf.Bytes())
	if err != nil {
		fmtSrc = signalBuf.Bytes()
	}
	err = os.WriteFile(name+".go", fmtSrc, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
func formatStructDocumentation(doc string) string {
	// TODO
	return "TODO: Doc formatting needs to be implemented!"
}

var primitives = map[string]string{
	"xs:double":          "float64",
	"xs:int":             "int",
	"xs:integer":         "int",
	"xs:positiveInteger": "int",
	"xs:string":          "string",
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
