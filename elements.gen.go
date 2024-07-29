// Code generated; DO NOT EDIT.

package xodr

type Element interface {
}

type OpenDRIVE struct {
	Header               *Header
	Road                 *Road
	Controller           *Controller
	JunctiontypeCrossing *JunctiontypeCrossing
	JunctiontypeDefault  *JunctiontypeDefault
	JunctiontypeDirect   *JunctiontypeDirect
	JunctiontypeVirtual  *JunctiontypeVirtual
	JunctionGroup        *JunctionGroup
	Station              *Station
}

type Header struct {
	GeoReference       *GeoReference
	Offset             *Offset
	License            *License
	DefaultRegulations *DefaultRegulations
}

type GeoReference struct {
}

type Offset struct {
}

type License struct {
}

type DefaultRegulations struct {
	RoadRegulations   *RoadRegulations
	SignalRegulations *SignalRegulations
}

type RoadRegulations struct {
}

type SignalRegulations struct {
}

type Road struct {
	Link             *Link
	ElevationProfile *ElevationProfile
	LateralProfile   *LateralProfile
	Lanes            *Lanes
	Objects          *Objects
	Signals          *Signals
	Surface          *Surface
	Railroad         *Railroad
}

type Link struct {
	Predecessor *Predecessor
	Successor   *Successor
}

type Predecessor struct {
}

type Successor struct {
}

type RoadType struct {
	Speed *Speed
}

type Speed struct {
}

type PlanView struct {
	Geometry *Geometry
}

type Geometry struct {
	Line       *Line
	Spiral     *Spiral
	Arc        *Arc
	Poly3      *Poly3
	ParamPoly3 *ParamPoly3
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
	Elevation *Elevation
}

type Elevation struct {
}

type LateralProfile struct {
	Superelevation      *Superelevation
	Shape               *Shape
	CrossSectionSurface *CrossSectionSurface
}

type Superelevation struct {
}

type Shape struct {
}

type CrossSectionSurface struct {
	TOffset       *TOffset
	SurfaceStrips *SurfaceStrips
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
	LaneOffset  *LaneOffset
	LaneSection *LaneSection
}

type LaneOffset struct {
}

type LaneSection struct {
	Left   *Left
	Center *Center
	Right  *Right
}

type Left struct {
	Lane *Lane
}

type Lane struct {
	Width    *Width
	Border   *Border
	Link     *Link
	RoadMark *RoadMark
	Material *Material
	Speed    *Speed
	Access   *Access
	Height   *Height
	Rule     *Rule
}

type Border struct {
	CornerReference *CornerReference
}

type RoadMark struct {
	Sway         *Sway
	RoadMarkType *RoadMarkType
	Explicit     *Explicit
}

type Sway struct {
}

type RoadMarkType struct {
	Line *Line
}

type Explicit struct {
	Line *Line
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
	Object          *Object
	ObjectReference *ObjectReference
	Tunnel          *Tunnel
	Bridge          *Bridge
}

type Object struct {
	Repeat       *Repeat
	Outline      *Outline
	Outlines     *Outlines
	Material     *Material
	Validity     *Validity
	ParkingSpace *ParkingSpace
	Markings     *Markings
	Borders      *Borders
	Surface      *Surface
	Skeleton     *Skeleton
}

type Repeat struct {
}

type Outline struct {
	CornerRoad  *CornerRoad
	CornerLocal *CornerLocal
}

type CornerRoad struct {
}

type CornerLocal struct {
}

type Outlines struct {
	Outline *Outline
}

type Validity struct {
}

type ParkingSpace struct {
}

type Markings struct {
	Marking *Marking
}

type Marking struct {
	CornerReference *CornerReference
}

type CornerReference struct {
}

type Borders struct {
	Border *Border
}

type Surface struct {
	CRG *CRG
}

type Skeleton struct {
	Polyline *Polyline
}

type Polyline struct {
	VertexRoad  *VertexRoad
	VertexLocal *VertexLocal
}

type VertexRoad struct {
}

type VertexLocal struct {
}

type ObjectReference struct {
	Validity *Validity
}

type Tunnel struct {
	Validity *Validity
}

type Bridge struct {
	Validity *Validity
}

type Signals struct {
	Signal          *Signal
	SignalReference *SignalReference
}

type Signal struct {
	PositionInertial *PositionInertial
	PositionRoad     *PositionRoad
	Validity         *Validity
	Dependency       *Dependency
	Reference        *Reference
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
	Validity *Validity
}

type CRG struct {
}

type Railroad struct {
	Switch *Switch
}

type Switch struct {
	MainTrack *MainTrack
	SideTrack *SideTrack
	Partner   *Partner
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
	Priority    *Priority
	Controller  *Controller
	Surface     *Surface
	PlanView    *PlanView
}

type RoadSection struct {
}

type Priority struct {
}

type JunctiontypeDefault struct {
	Connection    *Connection
	CrossPath     *CrossPath
	Priority      *Priority
	Controller    *Controller
	Surface       *Surface
	PlanView      *PlanView
	Boundary      *Boundary
	ElevationGrid *ElevationGrid
}

type Connection struct {
	LaneLink *LaneLink
}

type LaneLink struct {
}

type CrossPath struct {
	StartLaneLink *StartLaneLink
	EndLaneLink   *EndLaneLink
}

type StartLaneLink struct {
}

type EndLaneLink struct {
}

type Boundary struct {
	SegmenttypeJoint *SegmenttypeJoint
	SegmenttypeLane  *SegmenttypeLane
}

type SegmenttypeJoint struct {
}

type SegmenttypeLane struct {
}

type ElevationGrid struct {
	Elevation *Elevation
}

type JunctiontypeDirect struct {
	Connection *Connection
	Priority   *Priority
	Controller *Controller
	Surface    *Surface
	PlanView   *PlanView
}

type JunctiontypeVirtual struct {
	ConnectiontypeDefault *ConnectiontypeDefault
	ConnectiontypeVirtual *ConnectiontypeVirtual
	CrossPath             *CrossPath
	Priority              *Priority
	Controller            *Controller
	Surface               *Surface
	PlanView              *PlanView
}

type ConnectiontypeDefault struct {
	LaneLink *LaneLink
}

type ConnectiontypeVirtual struct {
	Predecessor *Predecessor
	Successor   *Successor
	LaneLink    *LaneLink
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
