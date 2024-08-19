package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"go/format"
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"
	"unicode"

	"github.com/iancoleman/strcase"
)

type SchemaDetail []any

var (
	// flags
	templPath *string
	xsdPath   *string

	funcCnt    = 0
	primitives = map[string]string{
		"xs:double":             "float64",
		"Xsdouble":              "float64",
		"GrEqZero":              "float64",
		"xs:int":                "int",
		"xs:integer":            "int",
		"xs:nonNegativeInteger": "int",
		"xs:positiveInteger":    "int",
		"XspositiveInteger":     "int",
		"xs:string":             "string",
		"Bool":                  "bool",
	}

	schemas = map[string][]any{
		"core":     SchemaDetail{"OpenDRIVE_Core.xsd", new(CoreSchema)},
		"junction": SchemaDetail{"OpenDRIVE_Junction.xsd", new(JunctionSchema)},
		"lane":     SchemaDetail{"OpenDRIVE_Lane.xsd", new(LaneSchema)},
		"object":   SchemaDetail{"OpenDRIVE_Object.xsd", new(ObjectSchema)},
		"railroad": SchemaDetail{"OpenDRIVE_Railroad.xsd", new(RailroadSchema)},
		"road":     SchemaDetail{"OpenDRIVE_Road.xsd", new(RoadSchema)},
		"signal":   SchemaDetail{"OpenDRIVE_Signal.xsd", new(SignalSchema)},
	}

	fnMap = template.FuncMap{
		"mapPrimitives":                        mapPrimitives,
		"formatStructDocumentation":            formatStructDocumentation,
		"toCamel":                              strcase.ToCamel,
		"canCreateTRoadSignalsSignalReference": canCreateTRoadSignalsSignalReference,
		"removeTypePrefix":                     removeTypePrefix,
		"distinctSignalsReferenceToCamel":      distinctSignalsReference,
		"goType":                               goType,
		"goUnion":                              goUnion,
	}
)

func main() {
	xsdPath = flag.String("xsd", ".", "path to xsd_schema directory")
	templPath = flag.String("templ", ".", "path to templates directory")
	flag.Parse()

	loadSchemaFiles()

	for name, schemaDetail := range schemas {
		fmt.Printf("generate %v\n", name)
		generate(schemaDetail[1], name)
	}
}

func goUnion(unionTypeInterface string, rawXodrTypes string) string {
	rawTypes := strings.Split(rawXodrTypes, " ")

	out := "func " + unionTypeInterface + "String(u " + unionTypeInterface + ") string {\n"
	out += "switch u.(type) {\n"

	for _, rawType := range rawTypes {
		gt := goType(rawType)
		out += "case " + gt + ": return u.(string)\n"
	}
	out += `default: return ""
}
}`

	return out
}

func goType(rawXodrType string) string {
	p, ok := primitives[rawXodrType]
	if ok {
		return p
	}
	t := removeTypePrefix(rawXodrType)

	t = distinctSignalsReference(t)

	t = strcase.ToCamel(t)
	return t
}

func loadSchemaFiles() {
	for name, schemaDetails := range schemas {
		fmt.Printf("load %v schema\n", name)

		file, err := os.ReadFile(*xsdPath + "/" + schemaDetails[0].(string))
		if err != nil {
			panic(err)
		}
		err = xml.Unmarshal(file, schemaDetails[1])
		if err != nil {
			panic(err)
		}

	}

}

func distinctSignalsReference(t string) string {
	switch t {
	case "road_signals_signal_reference":
		return "road_signals_signal_reference"
	case "road_signals_signalReference":
		return "road_signals_SpatialSignalReference"
	case "t_road_signals_signal_reference":
		return "road_signals_signal_reference"
	case "t_road_signals_signalReference":
		return "road_signals_SpatialSignalReference"
	default:
		return t

	}
}

func removeTypePrefix(t string) string {
	re, err := regexp.Compile("^t_")
	if err != nil {
		panic(err)
	}
	return string(re.ReplaceAll([]byte(t), []byte("")))
}

func canCreateTRoadSignalsSignalReference(camelCaseName string) bool {
	if camelCaseName != "TRoadSignalsSignalReference" {
		return true
	}
	funcCnt++
	return funcCnt < 2
}

func generate(schema any, name string) {
	tmpl := template.New(name + ".tmpl").Funcs(fnMap)
	tmpl, err := tmpl.ParseFiles(*templPath + name + ".tmpl")
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
