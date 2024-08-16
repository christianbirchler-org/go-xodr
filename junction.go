// Code generated; DO NOT EDIT.

package xodr

type EConnectionType struct {
}

type EContactPoint struct {
}

type EElementDir struct {
}

type EJunctionSegmentType struct {
}

type EJunctionType struct {
}

type EJunctionGroupType struct {
}

type ERoadSurfaceCrgMode struct {
}

type ERoadSurfaceCrgPurpose struct {
}

type TJunctionGridPositionList struct {
}

// TODO: Doc formatting needs to be implemented!
type TJunction struct {
}

// TODO: Doc formatting needs to be implemented!
type TJunctionBoundary struct {
	Segment TJunctionBoundarySegment
}

// TODO: Doc formatting needs to be implemented!
type TJunctionBoundarySegment struct {
}

// TODO: Doc formatting needs to be implemented!
type TJunctionBoundarySegmentJoint struct {
}

// TODO: Doc formatting needs to be implemented!
type TJunctionBoundarySegmentLane struct {
}

// TODO: Doc formatting needs to be implemented!
type TJunctionCommon struct {
	Connection TJunctionConnectionCommon

	CrossPath TJunctionCrossPath

	Priority TJunctionPriority

	Controller TJunctionController

	Surface TRoadSurface

	PlanView TRoadPlanView

	Objects TRoadObjects

	Boundary TJunctionBoundary

	ElevationGrid TJunctionElevationGrid
}

// TODO: Doc formatting needs to be implemented!
type TJunctionConnection struct {
	LaneLink TJunctionConnectionLaneLink
}

// TODO: Doc formatting needs to be implemented!
type TJunctionConnectionCommon struct {
}

// TODO: Doc formatting needs to be implemented!
type TJunctionConnectionDirect struct {
}

// TODO: Doc formatting needs to be implemented!
type TJunctionConnectionLaneLink struct {
}

// TODO: Doc formatting needs to be implemented!
type TJunctionConnectionVirtual struct {
	Predecessor TJunctionPredecessorSuccessor

	Successor TJunctionPredecessorSuccessor
}

// TODO: Doc formatting needs to be implemented!
type TJunctionConnectionVirtualDefault struct {
}

// TODO: Doc formatting needs to be implemented!
type TJunctionController struct {
}

// TODO: Doc formatting needs to be implemented!
type TJunctionCrossing struct {
	RoadSection TJunctionRoadSection

	Priority TJunctionPriority

	Controller TJunctionController

	Surface TRoadSurface

	PlanView TRoadPlanView

	Objects TRoadObjects
}

// TODO: Doc formatting needs to be implemented!
type TJunctionCrossPath struct {
	CrossingRoad string
	Id           string
	RoadAtEnd    string
	RoadAtStart  string
}

// TODO: Doc formatting needs to be implemented!
type TJunctionCrossPathLaneLink struct {
}

// TODO: Doc formatting needs to be implemented!
type TJunctionDirect struct {
	Connection TJunctionConnectionDirect

	Priority TJunctionPriority

	Controller TJunctionController

	Surface TRoadSurface

	PlanView TRoadPlanView

	Objects TRoadObjects
}

// TODO: Doc formatting needs to be implemented!
type TJunctionElevationGrid struct {
	Elevation TJunctionElevationGridElevation
}

// TODO: Doc formatting needs to be implemented!
type TJunctionElevationGridElevation struct {
}

// TODO: Doc formatting needs to be implemented!
type TJunctionPredecessorSuccessor struct {
}

// TODO: Doc formatting needs to be implemented!
type TJunctionPriority struct {
}

// TODO: Doc formatting needs to be implemented!
type TJunctionRoadSection struct {
}

// TODO: Doc formatting needs to be implemented!
type TJunctionVirtual struct {
	Connection TJunctionConnection

	CrossPath TJunctionCrossPath

	Priority TJunctionPriority

	Controller TJunctionController

	Surface TRoadSurface

	PlanView TRoadPlanView

	Objects TRoadObjects
}

// TODO: Doc formatting needs to be implemented!
type TJunctionGroup struct {
	JunctionReference TJunctionGroupJunctionReference
}

// TODO: Doc formatting needs to be implemented!
type TJunctionGroupJunctionReference struct {
}
