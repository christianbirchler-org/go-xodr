// Code generated; DO NOT EDIT.

package xodr

type ECountryCode interface{}

func ECountryCodeString(u ECountryCode) string {
	switch u.(type) {
	case ECountryCodeIso3166Alpha2:
		return u.(string)
	case ECountryCodeIso3166Alpha3Deprecated:
		return u.(string)
	case ECountryCodeDeprecated:
		return u.(string)
	default:
		return ""
	}
}

type ECountryCodeDeprecated string

type ECountryCodeIso3166Alpha2 string

type ECountryCodeIso3166Alpha3Deprecated string

type EDirection string

type EMaxSpeedString string

type EParamPoly3PRange string

type ERoadLinkElementType string

type ERoadType string

type EStripMode string

type ETrafficRule string

type MaxSpeed interface{}

func MaxSpeedString(u MaxSpeed) string {
	switch u.(type) {
	case GrEqZero:
		return u.(string)
	case EMaxSpeedString:
		return u.(string)
	default:
		return ""
	}
}

// Roads are the core elements for any road network in ASAM OpenDRIVE. Each ro
// ad runs along one road reference line. A road shall have at least the cente
// r lane. Vehicles may drive in both directions of the road reference line. T
// he standard driving direction is defined by the value which is assigned to
// the @rule attribute (RHT=right-hand traffic, LHT=left-han traffic). ASAM Op
// enDRIVE roads may be roads in the real road network or artificial road netw
// ork created for application use. Each road is described by one or more <roa
// d> elements. One <road> element may cover a long stretch of a road, shorter
// stretches between junctions, or even several roads. A new <road> element sh
// ould only start if the properties of the road cannot be described within th
// e previous <road> element or if a junction is required.d
type Road struct {
	OpenDriveElement
	Link             []*RoadLink
	Type             []*RoadType
	PlanView         []*RoadPlanView
	ElevationProfile []*RoadElevationProfile
	LateralProfile   []*RoadLateralProfile
	Lanes            []*RoadLanes
	Objects          []*RoadObjects
	Signals          []*RoadSignals
	Surface          []*RoadSurface
	Railroad         []*RoadRailroad
	Id               string
	Junction         string
	Length           GrZero
	Name             string
	Rule             ETrafficRule
}

// Road elevation specifies the elevation along the road reference line, that
// is in s-direction.
type RoadElevationProfile struct {
	OpenDriveElement
	Elevation []*RoadElevationProfileElevation
}

// Defines an elevation element at a given position on the road reference line
// . Elements shall be defined in ascending order along the reference line. Th
// e s length does not change with the elevation.
type RoadElevationProfileElevation struct {
	OpenDriveElement
	A float64
	B float64
	C float64
	D float64
	S GrEqZero
}

// Contains a series of lateral elevation elements that define the characteris
// tics of the road surfaces banking along the road reference line. The latera
// l profile is defined relative to the elevation of the road reference line.
type RoadLateralProfile struct {
	OpenDriveElement
	Superelevation      []*RoadLateralProfileSuperelevation
	Shape               []*RoadLateralProfileShape
	CrossSectionSurface []*RoadLateralProfileCrossSectionSurface
}

// A cross section surface defines the lateral profile by means of constant, l
// inear, quadratic, and cubic polynomials in t-direction. A cross section sur
// face is valid for the full length of the road.
type RoadLateralProfileCrossSectionSurface struct {
	OpenDriveElement
	TOffset       []*RoadLateralProfileCrossSectionSurfaceTOffset
	SurfaceStrips []*RoadLateralProfileCrossSectionSurfaceSurfaceStrip
}

// Defines the coefficients of a cubic polynomial in s-direction. The first <c
// oefficients> element shall start at the beginning of the road reference lin
// e with @s="0".
type RoadLateralProfileCrossSectionSurfaceCoefficients struct {
	OpenDriveElement
	A float64
	B float64
	C float64
	D float64
	S GrEqZero
}

// A strip defines the lateral profile in t- and s-direction.
type RoadLateralProfileCrossSectionSurfaceStrip struct {
	OpenDriveElement
	Width     []*RoadLateralProfileCrossSectionSurfaceStripWidth
	Constant  []*RoadLateralProfileCrossSectionSurfaceStripConstant
	Linear    []*RoadLateralProfileCrossSectionSurfaceStripLinear
	Quadratic []*RoadLateralProfileCrossSectionSurfaceStripQuadratic
	Cubic     []*RoadLateralProfileCrossSectionSurfaceStripCubic
	Id        int
	Mode      EStripMode
}

// Defines in t a constant height of the surface.
type RoadLateralProfileCrossSectionSurfaceStripConstant struct {
	OpenDriveElement
	Coefficients []*RoadLateralProfileCrossSectionSurfaceCoefficients
}

// Defines in t a cubic height of the surface.
type RoadLateralProfileCrossSectionSurfaceStripCubic struct {
	OpenDriveElement
	Coefficients []*RoadLateralProfileCrossSectionSurfaceCoefficients
}

// Defines in t a linear height of the surface.
type RoadLateralProfileCrossSectionSurfaceStripLinear struct {
	OpenDriveElement
	Coefficients []*RoadLateralProfileCrossSectionSurfaceCoefficients
}

// Defines in t a quadratic height of the surface.
type RoadLateralProfileCrossSectionSurfaceStripQuadratic struct {
	OpenDriveElement
	Coefficients []*RoadLateralProfileCrossSectionSurfaceCoefficients
}

// Defines the width of the inner strip.
type RoadLateralProfileCrossSectionSurfaceStripWidth struct {
	OpenDriveElement
	Coefficients []*RoadLateralProfileCrossSectionSurfaceCoefficients
}

// Surface strips contains the strips.
type RoadLateralProfileCrossSectionSurfaceSurfaceStrip struct {
	OpenDriveElement
	Strip []*RoadLateralProfileCrossSectionSurfaceStrip
}

// A t offset shifts all strips relative to the road reference line in t-direc
// tion.
type RoadLateralProfileCrossSectionSurfaceTOffset struct {
	OpenDriveElement
	Coefficients []*RoadLateralProfileCrossSectionSurfaceCoefficients
}

// Defined as the road sectionâ  s surface relative to the reference plane. T
// here may be several shape definitions at one s-position that have different
// t-values, thereby describing the curvy shape of the road.
type RoadLateralProfileShape struct {
	OpenDriveElement
	A float64
	B float64
	C float64
	D float64
	S GrEqZero
	T float64
}

// Superelevation specifies the transverse slope along the road reference line
// . Superelevation is constant in each cross section and can vary in road ref
// erence line direction. Elements must be defined in ascending order along th
// e reference line. The parameters of an element are valid until the next ele
// ment starts or the road reference line ends. Per default, the superelevatio
// n of a road is zero.
type RoadLateralProfileSuperelevation struct {
	OpenDriveElement
	A float64
	B float64
	C float64
	D float64
	S GrEqZero
}

// Follows the road header if the road is linked to a successor or a predecess
// or. Isolated roads may omit this element.
type RoadLink struct {
	OpenDriveElement
	Predecessor []*RoadLinkPredecessorSuccessor
	Successor   []*RoadLinkPredecessorSuccessor
}

// Successors and predecessors can be junctions or roads. For each, different
// attribute sets shall be used.
type RoadLinkPredecessorSuccessor struct {
	OpenDriveElement
	ContactPoint EContactPoint
	ElementDir   EElementDir
	ElementId    string
	ElementS     GrEqZero
	ElementType  ERoadLinkElementType
}

// Contains geometry elements that define the layout of the road reference lin
// e in the x/y-plane (plan view).
type RoadPlanView struct {
	OpenDriveElement
	Geometry []*RoadPlanViewGeometry
}

//
type RoadPlanViewGeometry struct {
	OpenDriveElement
	Hdg    float64
	Length GrZero
	S      GrEqZero
	X      float64
	Y      float64
}

// Arcs describe road reference lines with constant curvature.
type RoadPlanViewGeometryArc struct {
	OpenDriveElement
	Curvature float64
}

// A straight line is the simplest geometry element. It contains no further at
// tributes.
type RoadPlanViewGeometryLine struct {
	OpenDriveElement
}

// A parametric cubic curve describing the road reference line.
type RoadPlanViewGeometryParamPoly3 struct {
	OpenDriveElement
	AU     float64
	AV     float64
	BU     float64
	BV     float64
	CU     float64
	CV     float64
	DU     float64
	DV     float64
	PRange EParamPoly3PRange
}

// A cubic polynom describing the road reference line.
type RoadPlanViewGeometryPoly3 struct {
	OpenDriveElement
	A float64
	B float64
	C float64
	D float64
}

// Spirals or more specifically Euler spirals also known as clothoids. They de
// scribe road reference lines with constantly changing curvatures.
type RoadPlanViewGeometrySpiral struct {
	OpenDriveElement
	CurvEnd   float64
	CurvStart float64
}

// Contains a series of elements describing a surface.
type RoadSurface struct {
	OpenDriveElement
	Crg []*RoadSurfaceCrg
}

// Links road surface data defined according to ASAM OpenCRG format.
type RoadSurfaceCrg struct {
	OpenDriveElement
	File        string
	HOffset     float64
	Mode        ERoadSurfaceCrgMode
	Orientation EDirection
	Purpose     ERoadSurfaceCrgPurpose
	SEnd        GrEqZero
	SOffset     float64
	SStart      GrEqZero
	TOffset     float64
	ZOffset     float64
	ZScale      float64
}

// A road type element is valid for the entire cross section of a road. It is
// valid until a new road type element is provided or until the road ends.
type RoadType struct {
	OpenDriveElement
	Speed   []*RoadTypeSpeed
	Country ECountryCode
	S       GrEqZero
	Type    ERoadType
}

// Defines the default maximum speed allowed in conjunction with the specified
// road type.
type RoadTypeSpeed struct {
	OpenDriveElement
	Max  MaxSpeed
	Unit EUnitSpeed
}
