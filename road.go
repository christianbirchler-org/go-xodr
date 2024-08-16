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

type TMaxSpeed struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoad struct {
	Link TRoadLink

	Type TRoadType

	PlanView TRoadPlanView

	ElevationProfile TRoadElevationProfile

	LateralProfile TRoadLateralProfile

	Lanes TRoadLanes

	Objects TRoadObjects

	Signals TRoadSignals

	Surface TRoadSurface

	Railroad TRoadRailroad
}

// TODO: Doc formatting needs to be implemented!
type TRoadElevationProfile struct {
	Elevation TRoadElevationProfileElevation
}

// TODO: Doc formatting needs to be implemented!
type TRoadElevationProfileElevation struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadLateralProfile struct {
	Superelevation TRoadLateralProfileSuperelevation

	Shape TRoadLateralProfileShape

	CrossSectionSurface TRoadLateralProfileCrossSectionSurface
}

// TODO: Doc formatting needs to be implemented!
type TRoadLateralProfileCrossSectionSurface struct {
	TOffset TRoadLateralProfileCrossSectionSurfaceTOffset

	SurfaceStrips TRoadLateralProfileCrossSectionSurfaceSurfaceStrip
}

// TODO: Doc formatting needs to be implemented!
type TRoadLateralProfileCrossSectionSurfaceCoefficients struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadLateralProfileCrossSectionSurfaceStrip struct {
	Width TRoadLateralProfileCrossSectionSurfaceStripWidth

	Constant TRoadLateralProfileCrossSectionSurfaceStripConstant

	Linear TRoadLateralProfileCrossSectionSurfaceStripLinear

	Quadratic TRoadLateralProfileCrossSectionSurfaceStripQuadratic

	Cubic TRoadLateralProfileCrossSectionSurfaceStripCubic
}

// TODO: Doc formatting needs to be implemented!
type TRoadLateralProfileCrossSectionSurfaceStripConstant struct {
	Coefficients TRoadLateralProfileCrossSectionSurfaceCoefficients
}

// TODO: Doc formatting needs to be implemented!
type TRoadLateralProfileCrossSectionSurfaceStripCubic struct {
	Coefficients TRoadLateralProfileCrossSectionSurfaceCoefficients
}

// TODO: Doc formatting needs to be implemented!
type TRoadLateralProfileCrossSectionSurfaceStripLinear struct {
	Coefficients TRoadLateralProfileCrossSectionSurfaceCoefficients
}

// TODO: Doc formatting needs to be implemented!
type TRoadLateralProfileCrossSectionSurfaceStripQuadratic struct {
	Coefficients TRoadLateralProfileCrossSectionSurfaceCoefficients
}

// TODO: Doc formatting needs to be implemented!
type TRoadLateralProfileCrossSectionSurfaceStripWidth struct {
	Coefficients TRoadLateralProfileCrossSectionSurfaceCoefficients
}

// TODO: Doc formatting needs to be implemented!
type TRoadLateralProfileCrossSectionSurfaceSurfaceStrip struct {
	Strip TRoadLateralProfileCrossSectionSurfaceStrip
}

// TODO: Doc formatting needs to be implemented!
type TRoadLateralProfileCrossSectionSurfaceTOffset struct {
	Coefficients TRoadLateralProfileCrossSectionSurfaceCoefficients
}

// TODO: Doc formatting needs to be implemented!
type TRoadLateralProfileShape struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadLateralProfileSuperelevation struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadLink struct {
	Predecessor TRoadLinkPredecessorSuccessor

	Successor TRoadLinkPredecessorSuccessor
}

// TODO: Doc formatting needs to be implemented!
type TRoadLinkPredecessorSuccessor struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadPlanView struct {
	Geometry TRoadPlanViewGeometry
}

// TODO: Doc formatting needs to be implemented!
type TRoadPlanViewGeometry struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadPlanViewGeometryArc struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadPlanViewGeometryLine struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadPlanViewGeometryParamPoly3 struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadPlanViewGeometryPoly3 struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadPlanViewGeometrySpiral struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadSurface struct {
	Crg TRoadSurfaceCrg
}

// TODO: Doc formatting needs to be implemented!
type TRoadSurfaceCrg struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadType struct {
	Speed TRoadTypeSpeed
}

// TODO: Doc formatting needs to be implemented!
type TRoadTypeSpeed struct {
}
