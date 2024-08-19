// Code generated; DO NOT EDIT.

package xodr

type EDataQualityRawDataPostProcessing string

type EDataQualityRawDataSource string

type EUnit interface{}

func EUnitString(u EUnit) string {
	switch u.(type) {
	case EUnitDistance:
		return u.(string)
	case EUnitSpeed:
		return u.(string)
	case EUnitMass:
		return u.(string)
	case EUnitSlope:
		return u.(string)
	default:
		return ""
	}
}

type EUnitDistance string

type EUnitMass string

type EUnitSlope string

type EUnitSpeed string

type GrEqZero float64

type GrEqZeroOrContactPoint interface{}

func GrEqZeroOrContactPointString(u GrEqZeroOrContactPoint) string {
	switch u.(type) {
	case GrZero:
		return u.(string)
	case EContactPoint:
		return u.(string)
	default:
		return ""
	}
}

type GrZero float64

type YesNo string

type ZeroOne float64

// TODO: Doc formatting needs to be implemented!
type OpenDriveElement struct {
}

// TODO: Doc formatting needs to be implemented!
type DataQuality struct {
}

// TODO: Doc formatting needs to be implemented!
type DataQualityError struct {
}

// TODO: Doc formatting needs to be implemented!
type DataQualityRawData struct {
}

// TODO: Doc formatting needs to be implemented!
type Header struct {
	OpenDriveElement
	GeoReference       []*HeaderGeoReference
	Offset             []*HeaderOffset
	License            []*License
	DefaultRegulations []*HeaderDefaultRegulations
	Date               string
	East               float64
	Name               string
	North              float64
	RevMajor           int
	RevMinor           int
	South              float64
	Vendor             string
	Version            string
	West               float64
}

// TODO: Doc formatting needs to be implemented!
type HeaderDefaultRegulations struct {
	OpenDriveElement
	RoadRegulations   []*HeaderRoadRegulation
	SignalRegulations []*HeaderSignalRegulation
}

// TODO: Doc formatting needs to be implemented!
type HeaderGeoReference struct {
}

// TODO: Doc formatting needs to be implemented!
type HeaderOffset struct {
	OpenDriveElement
	Hdg float64
	X   float64
	Y   float64
	Z   float64
}

// TODO: Doc formatting needs to be implemented!
type HeaderRoadRegulation struct {
	OpenDriveElement
	Semantics []*SignalsSemantics
	Type      ERoadType
}

// TODO: Doc formatting needs to be implemented!
type HeaderSignalRegulation struct {
	OpenDriveElement
	Semantics []*SignalsSemantics
	Subtype   string
	Type      string
}

// TODO: Doc formatting needs to be implemented!
type Include struct {
}

// TODO: Doc formatting needs to be implemented!
type License struct {
	OpenDriveElement
	Name     string
	Resource string
	Spdxid   string
	Text     string
}

// TODO: Doc formatting needs to be implemented!
type UserData struct {
}
