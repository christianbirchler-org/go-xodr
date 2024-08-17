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
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type TJunctionBoundary struct {
	OpenDriveElement

	Segment TJunctionBoundarySegment
}

// TODO: Doc formatting needs to be implemented!
type TJunctionBoundarySegment struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type TJunctionBoundarySegmentJoint struct {
	TJunctionBoundarySegment
}

// TODO: Doc formatting needs to be implemented!
type TJunctionBoundarySegmentLane struct {
	TJunctionBoundarySegment
}

// TODO: Doc formatting needs to be implemented!
type TJunctionCommon struct {
	TJunction

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
	OpenDriveElement

	LaneLink TJunctionConnectionLaneLink
}

// TODO: Doc formatting needs to be implemented!
type TJunctionConnectionCommon struct {
	TJunctionConnection
}

// TODO: Doc formatting needs to be implemented!
type TJunctionConnectionDirect struct {
	TJunctionConnection
}

// TODO: Doc formatting needs to be implemented!
type TJunctionConnectionLaneLink struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type TJunctionConnectionVirtual struct {
	TJunctionConnection

	Predecessor TJunctionPredecessorSuccessor

	Successor TJunctionPredecessorSuccessor
}

// TODO: Doc formatting needs to be implemented!
type TJunctionConnectionVirtualDefault struct {
	TJunctionConnection
}

// TODO: Doc formatting needs to be implemented!
type TJunctionController struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type TJunctionCrossing struct {
	TJunction

	RoadSection TJunctionRoadSection

	Priority TJunctionPriority

	Controller TJunctionController

	Surface TRoadSurface

	PlanView TRoadPlanView

	Objects TRoadObjects
}

// TODO: Doc formatting needs to be implemented!
type TJunctionCrossPath struct {
	CrossingRoad  string
	Id            string
	RoadAtEnd     string
	RoadAtStart   string
	StartLaneLink TJunctionCrossPathLaneLink
	EndLaneLink   TJunctionCrossPathLaneLink
}

// TODO: Doc formatting needs to be implemented!
type TJunctionCrossPathLaneLink struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type TJunctionDirect struct {
	TJunction

	Connection TJunctionConnectionDirect

	Priority TJunctionPriority

	Controller TJunctionController

	Surface TRoadSurface

	PlanView TRoadPlanView

	Objects TRoadObjects
}

// TODO: Doc formatting needs to be implemented!
type TJunctionElevationGrid struct {
	OpenDriveElement

	Elevation TJunctionElevationGridElevation
}

// TODO: Doc formatting needs to be implemented!
type TJunctionElevationGridElevation struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type TJunctionPredecessorSuccessor struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type TJunctionPriority struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type TJunctionRoadSection struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type TJunctionVirtual struct {
	TJunction

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
	OpenDriveElement

	JunctionReference TJunctionGroupJunctionReference
}

// TODO: Doc formatting needs to be implemented!
type TJunctionGroupJunctionReference struct {
	OpenDriveElement
}
