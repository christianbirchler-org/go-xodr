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
	// Time/date of database creation according to ISO 8601 (preference: YYYY-MM-DDThh:mm:ss)
	Date string
	// Maximum inertial x value
	East float64
	// Database name
	Name string
	// Maximum inertial y value
	North float64
	// Major revision number of OpenDRIVE format
	RevMajor int
	// Minor revision number of OpenDRIVE format; 6 for OpenDrive 1.6
	RevMinor int
	// Minimum inertial y value
	South float64
	// Vendor name
	Vendor string
	// Version of this road network
	Version string
	// Minimum inertial x value
	West               float64
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
	// The full name of the license. Informational only.
	Name string
	// Link to an URL where the full license text can be found.
	Resource string
	// The identifier of the license from the SPDX license list. Can also be an SPDX License Expression, which is also applicable to custom licenses (LicenseRef-…).
	Spdxid string
	// The full license text
	Text string
}

type DefaultRegulations struct {
	RoadRegulations   *RoadRegulations
	SignalRegulations *SignalRegulations
}

type RoadRegulations struct {
	//
	Type string
}

type SignalRegulations struct {
	//
	Subtype string
	//
	Type string
}

type Road struct {
	// Unique ID within the database. If it represents an integer number, it should comply to uint32_t and stay within the given range.
	Id string
	// ID of the junction to which the road belongs, for example connecting roads, cross paths, and roads of a junction boundary. Use -1 for none.
	Junction string
	// Total length of the reference line in the xy-plane. Change in length due to elevation is not considered
	Length int
	// Name of the road. May be chosen freely.
	Name string
	// Basic rule for using the road; RHT=right-hand traffic, LHT=left-hand traffic. When this attribute is missing, RHT is assumed.
	Rule             string
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
	// Contact point of link on the linked element
	ContactPoint string
	// To be provided when elementS is used for the connection definition. Indicates the direction on the predecessor from which the road is entered.
	ElementDir string
	// ID of the linked element
	ElementId string
	// Alternative to contactPoint for virtual junctions. Indicates a connection within the predecessor, meaning not at the start or end of the predecessor. Shall only be used for elementType 'road'
	ElementS int
	// Type of the linked element
	ElementType string
}

type Successor struct {
	// Contact point of link on the linked element
	ContactPoint string
	// To be provided when elementS is used for the connection definition. Indicates the direction on the predecessor from which the road is entered.
	ElementDir string
	// ID of the linked element
	ElementId string
	// Alternative to contactPoint for virtual junctions. Indicates a connection within the predecessor, meaning not at the start or end of the predecessor. Shall only be used for elementType 'road'
	ElementS int
	// Type of the linked element
	ElementType string
}

type RoadType struct {
	// Country code of the road, see ISO 3166-1, alpha-2 codes.
	Country string
	// s-coordinate of start position
	S int
	// Type of the road defined as enumeration
	Type  string
	Speed *Speed
}

type Speed struct {
	// Maximum allowed speed. Given as string (only 'no limit' / 'undefined') or numerical value in the respective unit (see attribute unit). If the attribute unit is not specified, m/s is used as default.
	Max int
	// Unit of the attribute max. For values, see chapter 'units'.
	Unit string
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
	// Polynom parameter a, elevation at @s (ds=0)
	A float64
	// Polynom parameter b
	B float64
	// Polynom parameter c
	C float64
	// Polynom parameter d
	D float64
	// s-coordinate of start position
	S float64
}

type LateralProfile struct {
	Superelevation      *Superelevation
	Shape               *Shape
	CrossSectionSurface *CrossSectionSurface
}

type Superelevation struct {
	// Polynom parameter a, superelevation at @s (ds=0)
	A float64
	// Polynom parameter b
	B float64
	// Polynom parameter c
	C float64
	// Polynom parameter d
	D float64
	// s-coordinate of start position
	S float64
}

type Shape struct {
	// Polynom parameter a, relative height at @t (dt=0)
	A float64
	// Polynom parameter b
	B float64
	// Polynom parameter c
	C float64
	// Polynom parameter d
	D float64
	// s-coordinate of start position
	S float64
	// t-coordinate of start position
	T float64
}

type CrossSectionSurface struct {
	TOffset       *TOffset
	SurfaceStrips *SurfaceStrips
}

type TOffset struct {
}

type Coefficients struct {
	// Polynomial parameter a. If the attribute is not specified, the value is 0.
	A float64
	// Polynomial parameter b. If the attribute is not specified, the value is 0.
	B float64
	// Polynomial parameter c. If the attribute is not specified, the value is 0.
	C float64
	// Polynomial parameter d. If the attribute is not specified, the value is 0.
	D float64
	// s-coordinate of start position
	S float64
}

type SurfaceStrips struct {
}

type Strip struct {
	// 1 for the inner left strip, -1 for the inner right strip, 2 for the outer left strip, -2 for the outer right strip
	Id int
	// Can only be defined for an outer strip.
	Mode string
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
	// If true, lane can be used also by a neighboring lane. Advisory lane has priority, for example a bike lane, that can also be used by cars. If not specified, default value is none.
	Advisory string
	// If not specified, direction is determined by the combination of <left> or <right> lane grouping and the values LHT or RHT of the @rule attribute of a road. The standard direction can be overwritten with this attribute.
	Direction string
	// If true, lane direction can be changed dynamically by the scenario during the simulation. If not specified, default boolean value is false.
	DynamicLaneDirection string
	// If true, lane type can be changed dynamically by the scenario during the simulation. Typical example is a stop lane that can be changed by VMS boards to a driving lane. If not specified, default boolean value is false.
	DynamicLaneType string
	// ID of the lane
	Id int
	// 'true' = keep lane on level, that is, do not apply superelevation; 'false' = apply superelevation to this lane (default, also used if attribute level is missing)
	Level string
	// If true, lane is under construction.
	RoadWorks string
	// Type of the lane
	Type     string
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
	// Name of the file containing the CRG data
	File string
	// Heading offset between CRG center line and reference line of the road (only allowed for mode genuine, default = 0.0).
	HOffset float64
	// Attachment mode for the surface data, see specification.
	Mode string
	// Orientation of the CRG data set relative to the parent <road> element. Only allowed for mode attached and attached0.
	Orientation string
	// Physical purpose of the data contained in the CRG file; if the attribute is missing, data will be interpreted as elevation data.
	Purpose string
	// End of the application of CRG (s-coordinate)
	SEnd float64
	// s-offset between CRG center line and reference line of the road (default = 0.0)
	SOffset float64
	// Start of the application of CRG data (s-coordinate)
	SStart float64
	// t-offset between CRG center line and reference line of the road (default = 0.0)
	TOffset float64
	// z-offset between CRG center line and reference line of the road (default = 0.0). Only allowed for purpose elevation.
	ZOffset float64
	// z-scale factor for the surface description (default = 1.0). Only allowed for purpose elevation.
	ZScale float64
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
