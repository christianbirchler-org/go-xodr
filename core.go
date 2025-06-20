// Code generated; DO NOT EDIT.

package xodr

// Root element containing all information about the ASAM OpenDRIVE file
type OpenDrive struct {
	Header        *Header                `xml:"header"`
	Road          []*Road                `xml:"road"`
	Controller    []*Controller          `xml:"controller"`
	Junction      []*Junction            `xml:"junction"`
	JunctionGroup []*JunctionGroup       `xml:"junctionGroup"`
	Station       []*Station             `xml:"station"`
	VmsGroup      []*SignalGroupVmsGroup `xml:"vmsGroup"`
}

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

type OpenDriveElement struct {
}

// Describes the quality and accuracy of measurement data that is integrated i
// nto the ASAM OpenDRIVE file.
type DataQuality struct {
}

// Describes the error range, given in [m], of measurement data that is integr
// ated into the ASAM OpenDRIVE file.
type DataQualityError struct {
}

// Describes some basic metadata containing information about the raw data.
type DataQualityRawData struct {
}

// Contains general information about the ASAM OpenDRIVE file
type Header struct {
	OpenDriveElement
	GeoReference       *HeaderGeoReference       `xml:"geoReference"`
	Offset             *HeaderOffset             `xml:"offset"`
	License            *License                  `xml:"license"`
	DefaultRegulations *HeaderDefaultRegulations `xml:"defaultRegulations"`
	Date               string                    `xml:"date,attr"`
	East               float64                   `xml:"east,attr"`
	Name               string                    `xml:"name,attr"`
	North              float64                   `xml:"north,attr"`
	RevMajor           int                       `xml:"revMajor,attr"`
	RevMinor           int                       `xml:"revMinor,attr"`
	South              float64                   `xml:"south,attr"`
	Vendor             string                    `xml:"vendor,attr"`
	Version            string                    `xml:"version,attr"`
	West               float64                   `xml:"west,attr"`
}

// Defines the default regulations. In each country there are different speed
// limits to a rural road. For example a rural road has a speed limit of 100km
// /h in Germany and 80km/h in the Netherlands. In some countries, one is allo
// wed to turn right at a red traffic light; in others, one is not. Instead of
// writing this for each road or each signal, the default regulations can be s
// pecified once in the header for the entire {THIS_STANDARD} file. The defaul
// t driving regulations can be overwritten with road, lane, or signal definit
// ions.
type HeaderDefaultRegulations struct {
	OpenDriveElement
	RoadRegulations   []*HeaderRoadRegulation   `xml:"roadRegulations"`
	SignalRegulations []*HeaderSignalRegulation `xml:"signalRegulations"`
}

// Spatial reference systems are standardized by the European Petroleum Survey
// Group Geodesy (EPSG) and are defined by parameters describing the geodetic
// datum. A geodetic datum is a coordinate reference system for a collection o
// f positions that are relative to an ellipsoid model of the earth. A geodeti
// c datum is described by a projection string according to PROJ, that is, a f
// ormat for the exchange of data between two coordinate systems. This data sh
// all be marked as CDATA, because it may contain characters that interfere wi
// th the XML syntax of an elementâ  s attribute.
type HeaderGeoReference struct {
}

// To avoid large coordinates, an offset of the whole dataset may be applied u
// sing the <offset> element. It enables inertial relocation and re-orientatio
// n of datasets. The dataset is first translated by @x, @y, and @z. Afterward
// s, it is rotated by @hdg around the new origin. Rotation around the z-axis
// should be avoided.
type HeaderOffset struct {
	OpenDriveElement
	Hdg float64 `xml:"hdg,attr"`
	X   float64 `xml:"x,attr"`
	Y   float64 `xml:"y,attr"`
	Z   float64 `xml:"z,attr"`
}

// Defines the default regulations for different road types.
type HeaderRoadRegulation struct {
	OpenDriveElement
	Semantics *SignalsSemantics `xml:"semantics"`
	Type      ERoadType         `xml:"type,attr"`
}

// Defines the default regulations for signs in different countries, for examp
// le, if it is allowed to turn right when a red traffic light appears.
type HeaderSignalRegulation struct {
	OpenDriveElement
	Semantics *SignalsSemantics `xml:"semantics"`
	Subtype   string            `xml:"subtype,attr"`
	Type      string            `xml:"type,attr"`
}

// Provides information about additional files that should be included while p
// rocessing the ASAM OpenDRIVE file.
type Include struct {
}

// Licensing information about the OpenDRIVE file.
type License struct {
	OpenDriveElement
	Name     string `xml:"name,attr"`
	Resource string `xml:"resource,attr"`
	Spdxid   string `xml:"spdxid,attr"`
	Text     string `xml:"text,attr"`
}

// Describes any additional information or data that is needed by an applicati
// on for a specific reason.
type UserData struct {
}
