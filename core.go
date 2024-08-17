// Code generated; DO NOT EDIT.

package xodr

type EDataQualityRawDataPostProcessing struct {
}

type EDataQualityRawDataSource struct {
}

type EUnit struct {
}

type EUnitDistance struct {
}

type EUnitMass struct {
}

type EUnitSlope struct {
}

type EUnitSpeed struct {
}

type TGrEqZero struct {
}

type TGrEqZeroOrContactPoint struct {
}

type TGrZero struct {
}

type TYesNo struct {
}

type TZeroOne struct {
}

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
	GeoReference       HeaderGeoReference
	Offset             HeaderOffset
	License            License
	DefaultRegulations HeaderDefaultRegulations
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
	RoadRegulations   HeaderRoadRegulation
	SignalRegulations HeaderSignalRegulation
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
	Semantics SignalsSemantics
	Type      ERoadType
}

// TODO: Doc formatting needs to be implemented!
type HeaderSignalRegulation struct {
	OpenDriveElement
	Semantics SignalsSemantics
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
