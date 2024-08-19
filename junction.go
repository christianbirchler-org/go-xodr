// Code generated; DO NOT EDIT.

package xodr

type EConnectionType string

type EContactPoint string

type EElementDir string

type EJunctionSegmentType string

type EJunctionType string

type EJunctionGroupType string

type ERoadSurfaceCrgMode string

type ERoadSurfaceCrgPurpose string

type JunctionGridPositionList []float64

// TODO: Doc formatting needs to be implemented!
type Junction struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type JunctionBoundary struct {
	OpenDriveElement
	Segment []*JunctionBoundarySegment
}

// TODO: Doc formatting needs to be implemented!
type JunctionBoundarySegment struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type JunctionBoundarySegmentJoint struct {
	JunctionBoundarySegment
}

// TODO: Doc formatting needs to be implemented!
type JunctionBoundarySegmentLane struct {
	JunctionBoundarySegment
}

// TODO: Doc formatting needs to be implemented!
type JunctionCommon struct {
	Junction
	Connection    []*JunctionConnectionCommon
	CrossPath     []*JunctionCrossPath
	Priority      []*JunctionPriority
	Controller    []*JunctionController
	Surface       []*RoadSurface
	PlanView      []*RoadPlanView
	Objects       []*RoadObjects
	Boundary      []*JunctionBoundary
	ElevationGrid []*JunctionElevationGrid
}

// TODO: Doc formatting needs to be implemented!
type JunctionConnection struct {
	OpenDriveElement
	LaneLink []*JunctionConnectionLaneLink
}

// TODO: Doc formatting needs to be implemented!
type JunctionConnectionCommon struct {
	JunctionConnection
}

// TODO: Doc formatting needs to be implemented!
type JunctionConnectionDirect struct {
	JunctionConnection
}

// TODO: Doc formatting needs to be implemented!
type JunctionConnectionLaneLink struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type JunctionConnectionVirtual struct {
	JunctionConnection
	Predecessor []*JunctionPredecessorSuccessor
	Successor   []*JunctionPredecessorSuccessor
}

// TODO: Doc formatting needs to be implemented!
type JunctionConnectionVirtualDefault struct {
	JunctionConnection
}

// TODO: Doc formatting needs to be implemented!
type JunctionController struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type JunctionCrossing struct {
	Junction
	RoadSection []*JunctionRoadSection
	Priority    []*JunctionPriority
	Controller  []*JunctionController
	Surface     []*RoadSurface
	PlanView    []*RoadPlanView
	Objects     []*RoadObjects
}

// TODO: Doc formatting needs to be implemented!
type JunctionCrossPath struct {
	CrossingRoad  string
	Id            string
	RoadAtEnd     string
	RoadAtStart   string
	StartLaneLink JunctionCrossPathLaneLink
	EndLaneLink   JunctionCrossPathLaneLink
}

// TODO: Doc formatting needs to be implemented!
type JunctionCrossPathLaneLink struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type JunctionDirect struct {
	Junction
	Connection []*JunctionConnectionDirect
	Priority   []*JunctionPriority
	Controller []*JunctionController
	Surface    []*RoadSurface
	PlanView   []*RoadPlanView
	Objects    []*RoadObjects
}

// TODO: Doc formatting needs to be implemented!
type JunctionElevationGrid struct {
	OpenDriveElement
	Elevation []*JunctionElevationGridElevation
}

// TODO: Doc formatting needs to be implemented!
type JunctionElevationGridElevation struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type JunctionPredecessorSuccessor struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type JunctionPriority struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type JunctionRoadSection struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type JunctionVirtual struct {
	Junction
	Connection []*JunctionConnection
	CrossPath  []*JunctionCrossPath
	Priority   []*JunctionPriority
	Controller []*JunctionController
	Surface    []*RoadSurface
	PlanView   []*RoadPlanView
	Objects    []*RoadObjects
}

// TODO: Doc formatting needs to be implemented!
type JunctionGroup struct {
	OpenDriveElement
	JunctionReference []*JunctionGroupJunctionReference
}

// TODO: Doc formatting needs to be implemented!
type JunctionGroupJunctionReference struct {
	OpenDriveElement
}
