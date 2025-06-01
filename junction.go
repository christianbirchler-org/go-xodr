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

// Junctions model intersections between roads.
type Junction struct {
	OpenDriveElement
}

// Junction boundaries enclose the area intended for traffic. This also includ
// es the sidewalks for pedestrians.
type JunctionBoundary struct {
	OpenDriveElement
	Segment []*JunctionBoundarySegment
}

// Segments run counter clockwise around the junction and form a closed juncti
// on boundary.
type JunctionBoundarySegment struct {
	OpenDriveElement
}

// A segment element with @type="joint" is perpendicular to the start or end o
// f the given road.
type JunctionBoundarySegmentJoint struct {
	JunctionBoundarySegment
}

// A segment element with @type="lane" goes along @boundaryLane for the given
// s range. It is the outmost edge of the lane  relative to the center of the
// junction.
type JunctionBoundarySegmentLane struct {
	JunctionBoundarySegment
}

// Common junctions are the default type of junction in ASAM OpenDRIVE and spe
// cify areas where drivable lanes may overlap and traffic may cross.
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

// Provides information about a single connection within a junction.
type JunctionConnection struct {
	OpenDriveElement
	LaneLink []*JunctionConnectionLaneLink
}

// Provides information about a single connection within a common junction.
type JunctionConnectionCommon struct {
	JunctionConnection
}

// Provides information about a single connection within a direct junction.
type JunctionConnectionDirect struct {
	JunctionConnection
}

// Provides information about the lanes that are linked between an incoming ro
// ad and a connecting road. It is strongly recommended to provide this elemen
// t. It is deprecated to omit the <laneLink> element.
type JunctionConnectionLaneLink struct {
	OpenDriveElement
}

// Virtual connections indicate possible connections between two roads or one
// or more lanes of two roads. Virtual connections do not specify connecting r
// oads.
type JunctionConnectionVirtual struct {
	JunctionConnection
	Predecessor []*JunctionPredecessorSuccessor
	Successor   []*JunctionPredecessorSuccessor
}

// Provides information about a single connection within a virtual junction.
type JunctionConnectionVirtualDefault struct {
	JunctionConnection
}

// Lists the controllers that should be grouped in a sychronization group (lim
// ited to that particular junction).
type JunctionController struct {
	OpenDriveElement
}

// Crossings are junctions without connecting roads. They define sections wher
// e crossing traffic can appear. Traffic does not change roads at crossings,
// for example, at railway crossings.
type JunctionCrossing struct {
	Junction
	RoadSection []*JunctionRoadSection
	Priority    []*JunctionPriority
	Controller  []*JunctionController
	Surface     []*RoadSurface
	PlanView    []*RoadPlanView
	Objects     []*RoadObjects
}

// Cross paths are intended for pedestrian crossings and are junctions element
// s where traffic of a lane can cross other lanes and continue on a different
// lane of the same or a different road. The cross path itself is a separate r
// oad.
type JunctionCrossPath struct {
	CrossingRoad  string
	Id            string
	RoadAtEnd     string
	RoadAtStart   string
	StartLaneLink JunctionCrossPathLaneLink
	EndLaneLink   JunctionCrossPathLaneLink
}

// Define the links between the lanes of the <crossPath> to the lanes of other
// roads.
type JunctionCrossPathLaneLink struct {
	OpenDriveElement
}

// Direct junctions are intended to model entries and exits where drivable lan
// es may overlap to split or merge, but traffic does not cross.
type JunctionDirect struct {
	Junction
	Connection []*JunctionConnectionDirect
	Priority   []*JunctionPriority
	Controller []*JunctionController
	Surface    []*RoadSurface
	PlanView   []*RoadPlanView
	Objects    []*RoadObjects
}

// An elevation grid is a coarse square grid with z-values at evenly spaced po
// ints. Elevation grids do not replace OpenCRG.
type JunctionElevationGrid struct {
	OpenDriveElement
	Elevation []*JunctionElevationGridElevation
}

// Defines the z-values at the regular grid points along the junction referenc
// e line.
type JunctionElevationGridElevation struct {
	OpenDriveElement
}

// Provides detailed information about the predecessor / successor road of a v
// irtual connection. Currently, only the @elementType â  roadâ   is allowed
// .
type JunctionPredecessorSuccessor struct {
	OpenDriveElement
}

// The junction priority record provides information about the priority of one
// road over another road that are part of this junction. It is only required
// if priorities cannot be derived from signs or signals in a junction or on t
// racks leading to a junction.
type JunctionPriority struct {
	OpenDriveElement
}

// Define the s range of the crossing roads with possible crossing traffic.
type JunctionRoadSection struct {
	OpenDriveElement
}

// Virtual junctions manage connections within an uninterrupted road, for exam
// ple, entries and exits to parking lots, and pedestrian crossings.
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

// Junction groups indicate for routing that the grouped junctions belong to t
// he same node and are commonly seen as one big junction, for example roundab
// outs or highway interchanges. The <junctionGroup> element is split into a h
// eader element and a series of member elements.
type JunctionGroup struct {
	OpenDriveElement
	JunctionReference []*JunctionGroupJunctionReference
}

// References to existing <junction> elements.
type JunctionGroupJunctionReference struct {
	OpenDriveElement
}
