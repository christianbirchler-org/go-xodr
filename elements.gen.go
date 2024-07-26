// Code generated; DO NOT EDIT.

package xodr

type Element interface {
}

type OpenDRIVE struct {
	Header *Header

	Road *Road

	Controller *Controller

	JunctiontypeCrossing *JunctiontypeCrossing

	JunctiontypeDefault *JunctiontypeDefault

	JunctiontypeDirect *JunctiontypeDirect

	JunctiontypeVirtual *JunctiontypeVirtual

	JunctionGroup *JunctionGroup

	Station *Station
}

type Header struct {
	GeoReference *GeoReference

	Offset *Offset

	License *License

	DefaultRegulations *DefaultRegulations
}

type GeoReference struct {
}

type Offset struct {
}

type License struct {
}

type DefaultRegulations struct {
	RoadRegulations *RoadRegulations

	SignalRegulations *SignalRegulations
}

type RoadRegulations struct {
}

type SignalRegulations struct {
}

type Road struct {
	Link *Link

	ElevationProfile *ElevationProfile

	LateralProfile *LateralProfile

	Lanes *Lanes

	Objects *Objects

	Signals *Signals

	Surface *Surface

	Railroad *Railroad
}

type Link struct {
	Predecessor *Predecessor

	Successor *Successor
}

type Predecessor struct {
}

type Successor struct {
}

type Type struct {
	Speed *Speed
}

type Speed struct {
}

type PlanView struct {
}

type Geometry struct {
}

type Line struct {
}

type Spiral struct {
}

type Arc struct {
}

type Poly3 struct {
}

type ParamPoly3 struct {
}

type ElevationProfile struct {
}

type Elevation struct {
}

type LateralProfile struct {
}

type Superelevation struct {
}

type Shape struct {
}

type CrossSectionSurface struct {
}

type TOffset struct {
}

type Coefficients struct {
}

type SurfaceStrips struct {
}

type Strip struct {
}

type Width struct {
}

type Constant struct {
}

type Linear struct {
}

type Quadratic struct {
}

type Cubic struct {
}

type Lanes struct {
}

type LaneOffset struct {
}

type LaneSection struct {
}

type Left struct {
}

type Lane struct {
}

type Border struct {
}

type RoadMark struct {
}

type Sway struct {
}

type Explicit struct {
}

type Material struct {
}

type Access struct {
}

type Restriction struct {
}

type Height struct {
}

type Rule struct {
}

type Center struct {
}

type Right struct {
}

type Objects struct {
}

type Object struct {
}

type Repeat struct {
}

type Outline struct {
}

type CornerRoad struct {
}

type CornerLocal struct {
}

type Outlines struct {
}

type Validity struct {
}

type ParkingSpace struct {
}

type Markings struct {
}

type Marking struct {
}

type CornerReference struct {
}

type Borders struct {
}

type Surface struct {
}

type Skeleton struct {
}

type Polyline struct {
}

type VertexRoad struct {
}

type VertexLocal struct {
}

type ObjectReference struct {
}

type Tunnel struct {
}

type Bridge struct {
}

type Signals struct {
}

type Signal struct {
}

type PositionInertial struct {
}

type PositionRoad struct {
}

type Dependency struct {
}

type Reference struct {
}

type SignalReference struct {
}

type CRG struct {
}

type Railroad struct {
}

type Switch struct {
}

type MainTrack struct {
}

type SideTrack struct {
}

type Partner struct {
}

type Controller struct {
	Control *Control
}

type Control struct {
}

type JunctiontypeCrossing struct {
	RoadSection *RoadSection

	Priority *Priority

	Controller *Controller

	Surface *Surface

	PlanView *PlanView
}

type RoadSection struct {
}

type Priority struct {
}

type JunctiontypeDefault struct {
	Connection *Connection

	CrossPath *CrossPath

	Priority *Priority

	Controller *Controller

	Surface *Surface

	PlanView *PlanView

	Boundary *Boundary

	ElevationGrid *ElevationGrid
}

type Connection struct {
}

type LaneLink struct {
}

type CrossPath struct {
}

type StartLaneLink struct {
}

type EndLaneLink struct {
}

type Boundary struct {
}

type SegmenttypeJoint struct {
}

type SegmenttypeLane struct {
}

type ElevationGrid struct {
}

type JunctiontypeDirect struct {
	Connection *Connection

	Priority *Priority

	Controller *Controller

	Surface *Surface

	PlanView *PlanView
}

type JunctiontypeVirtual struct {
	ConnectiontypeDefault *ConnectiontypeDefault

	ConnectiontypeVirtual *ConnectiontypeVirtual

	CrossPath *CrossPath

	Priority *Priority

	Controller *Controller

	Surface *Surface

	PlanView *PlanView
}

type ConnectiontypeDefault struct {
}

type ConnectiontypeVirtual struct {
}

type JunctionGroup struct {
	JunctionReference *JunctionReference
}

type JunctionReference struct {
}

type Station struct {
	Platform *Platform
}

type Platform struct {
	Segment *Segment
}

type Segment struct {
}
