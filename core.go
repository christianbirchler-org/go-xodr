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
type TDataQuality struct {
}

// TODO: Doc formatting needs to be implemented!
type TDataQualityError struct {
	XyAbsolute float64
	XyRelative float64
	ZAbsolute  float64
	ZRelative  float64
}

// TODO: Doc formatting needs to be implemented!
type TDataQualityRawData struct {
	Date                  string
	PostProcessing        EDataQualityRawDataPostProcessing
	PostProcessingComment string
	Source                EDataQualityRawDataSource
	SourceComment         string
}

// TODO: Doc formatting needs to be implemented!
type THeader struct {
	OpenDriveElement
	GeoReference       THeaderGeoReference
	Offset             THeaderOffset
	License            TLicense
	DefaultRegulations THeaderDefaultRegulations
}

// TODO: Doc formatting needs to be implemented!
type THeaderDefaultRegulations struct {
	OpenDriveElement
	RoadRegulations   THeaderRoadRegulation
	SignalRegulations THeaderSignalRegulation
}

// TODO: Doc formatting needs to be implemented!
type THeaderGeoReference struct {
}

// TODO: Doc formatting needs to be implemented!
type THeaderOffset struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type THeaderRoadRegulation struct {
	OpenDriveElement
	Semantics TSignalsSemantics
}

// TODO: Doc formatting needs to be implemented!
type THeaderSignalRegulation struct {
	OpenDriveElement
	Semantics TSignalsSemantics
}

// TODO: Doc formatting needs to be implemented!
type TInclude struct {
	File string
}

// TODO: Doc formatting needs to be implemented!
type TLicense struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type TUserData struct {
	Code  string
	Value string
}
