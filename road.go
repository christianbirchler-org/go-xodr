// Code generated; DO NOT EDIT.

package xodr

type ECountryCode struct {
}

type ECountryCodeDeprecated struct {
}

type ECountryCodeIso3166Alpha2 struct {
}

type ECountryCodeIso3166Alpha3Deprecated struct {
}

type EDirection struct {
}

type EMaxSpeedString struct {
}

type EParamPoly3PRange struct {
}

type ERoadLinkElementType struct {
}

type ERoadType struct {
}

type EStripMode struct {
}

type ETrafficRule struct {
}

type MaxSpeed struct {
}

// TODO: Doc formatting needs to be implemented!
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

// TODO: Doc formatting needs to be implemented!
type RoadElevationProfile struct {
	OpenDriveElement
	Elevation []*RoadElevationProfileElevation
}

// TODO: Doc formatting needs to be implemented!
type RoadElevationProfileElevation struct {
	OpenDriveElement
	A float64
	B float64
	C float64
	D float64
	S GrEqZero
}

// TODO: Doc formatting needs to be implemented!
type RoadLateralProfile struct {
	OpenDriveElement
	Superelevation      []*RoadLateralProfileSuperelevation
	Shape               []*RoadLateralProfileShape
	CrossSectionSurface []*RoadLateralProfileCrossSectionSurface
}

// TODO: Doc formatting needs to be implemented!
type RoadLateralProfileCrossSectionSurface struct {
	OpenDriveElement
	TOffset       []*RoadLateralProfileCrossSectionSurfaceTOffset
	SurfaceStrips []*RoadLateralProfileCrossSectionSurfaceSurfaceStrip
}

// TODO: Doc formatting needs to be implemented!
type RoadLateralProfileCrossSectionSurfaceCoefficients struct {
	OpenDriveElement
	A float64
	B float64
	C float64
	D float64
	S GrEqZero
}

// TODO: Doc formatting needs to be implemented!
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

// TODO: Doc formatting needs to be implemented!
type RoadLateralProfileCrossSectionSurfaceStripConstant struct {
	OpenDriveElement
	Coefficients []*RoadLateralProfileCrossSectionSurfaceCoefficients
}

// TODO: Doc formatting needs to be implemented!
type RoadLateralProfileCrossSectionSurfaceStripCubic struct {
	OpenDriveElement
	Coefficients []*RoadLateralProfileCrossSectionSurfaceCoefficients
}

// TODO: Doc formatting needs to be implemented!
type RoadLateralProfileCrossSectionSurfaceStripLinear struct {
	OpenDriveElement
	Coefficients []*RoadLateralProfileCrossSectionSurfaceCoefficients
}

// TODO: Doc formatting needs to be implemented!
type RoadLateralProfileCrossSectionSurfaceStripQuadratic struct {
	OpenDriveElement
	Coefficients []*RoadLateralProfileCrossSectionSurfaceCoefficients
}

// TODO: Doc formatting needs to be implemented!
type RoadLateralProfileCrossSectionSurfaceStripWidth struct {
	OpenDriveElement
	Coefficients []*RoadLateralProfileCrossSectionSurfaceCoefficients
}

// TODO: Doc formatting needs to be implemented!
type RoadLateralProfileCrossSectionSurfaceSurfaceStrip struct {
	OpenDriveElement
	Strip []*RoadLateralProfileCrossSectionSurfaceStrip
}

// TODO: Doc formatting needs to be implemented!
type RoadLateralProfileCrossSectionSurfaceTOffset struct {
	OpenDriveElement
	Coefficients []*RoadLateralProfileCrossSectionSurfaceCoefficients
}

// TODO: Doc formatting needs to be implemented!
type RoadLateralProfileShape struct {
	OpenDriveElement
	A float64
	B float64
	C float64
	D float64
	S GrEqZero
	T float64
}

// TODO: Doc formatting needs to be implemented!
type RoadLateralProfileSuperelevation struct {
	OpenDriveElement
	A float64
	B float64
	C float64
	D float64
	S GrEqZero
}

// TODO: Doc formatting needs to be implemented!
type RoadLink struct {
	OpenDriveElement
	Predecessor []*RoadLinkPredecessorSuccessor
	Successor   []*RoadLinkPredecessorSuccessor
}

// TODO: Doc formatting needs to be implemented!
type RoadLinkPredecessorSuccessor struct {
	OpenDriveElement
	ContactPoint EContactPoint
	ElementDir   EElementDir
	ElementId    string
	ElementS     GrEqZero
	ElementType  ERoadLinkElementType
}

// TODO: Doc formatting needs to be implemented!
type RoadPlanView struct {
	OpenDriveElement
	Geometry []*RoadPlanViewGeometry
}

// TODO: Doc formatting needs to be implemented!
type RoadPlanViewGeometry struct {
	OpenDriveElement
	Hdg    float64
	Length GrZero
	S      GrEqZero
	X      float64
	Y      float64
}

// TODO: Doc formatting needs to be implemented!
type RoadPlanViewGeometryArc struct {
	OpenDriveElement
	Curvature float64
}

// TODO: Doc formatting needs to be implemented!
type RoadPlanViewGeometryLine struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
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

// TODO: Doc formatting needs to be implemented!
type RoadPlanViewGeometryPoly3 struct {
	OpenDriveElement
	A float64
	B float64
	C float64
	D float64
}

// TODO: Doc formatting needs to be implemented!
type RoadPlanViewGeometrySpiral struct {
	OpenDriveElement
	CurvEnd   float64
	CurvStart float64
}

// TODO: Doc formatting needs to be implemented!
type RoadSurface struct {
	OpenDriveElement
	Crg []*RoadSurfaceCrg
}

// TODO: Doc formatting needs to be implemented!
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

// TODO: Doc formatting needs to be implemented!
type RoadType struct {
	OpenDriveElement
	Speed   []*RoadTypeSpeed
	Country ECountryCode
	S       GrEqZero
	Type    ERoadType
}

// TODO: Doc formatting needs to be implemented!
type RoadTypeSpeed struct {
	OpenDriveElement
	Max  MaxSpeed
	Unit EUnitSpeed
}
