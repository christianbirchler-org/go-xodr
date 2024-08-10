// Code generated; DO NOT EDIT.

package xodr

type Element interface {
}

type OpenDrive struct {
	Header             *Header             `xml:"header"`
	Road               *Road               `xml:"road"`
	JunctionController *JunctionController `xml:"controller"`
	JunctionCrossing   *JunctionCrossing   `xml:"junction type='crossing'"`
	JunctionDefault    *JunctionDefault    `xml:"junction type='default'"`
	JunctionDirect     *JunctionDirect     `xml:"junction type='direct'"`
	JunctionVirtual    *JunctionVirtual    `xml:"junction type='virtual'"`
	JunctionGroup      *JunctionGroup      `xml:"junctionGroup"`
	Station            *Station            `xml:""`
}

type Header struct {
	// Time/date of database creation according to ISO 8601 (preference: YYYY-MM-DDThh:mm:ss)
	Date string `xml:"date"`
	// Maximum inertial x value
	East float64 `xml:"east"`
	// Database name
	Name string `xml:"name"`
	// Maximum inertial y value
	North float64 `xml:"north"`
	// Major revision number of OpenDRIVE format
	RevMajor int `xml:"revMajor"`
	// Minor revision number of OpenDRIVE format; 6 for OpenDrive 1.6
	RevMinor int `xml:"revMinor"`
	// Minimum inertial y value
	South float64 `xml:"south"`
	// Vendor name
	Vendor string `xml:"vendor"`
	// Version of this road network
	Version string `xml:"version"`
	// Minimum inertial x value
	West               float64             `xml:"west"`
	GeoReference       *GeoReference       `xml:"geoReference"`
	Offset             *Offset             `xml:"offset"`
	License            *License            `xml:"license"`
	DefaultRegulations *DefaultRegulations `xml:"defaultRegulations"`
}

type GeoReference struct {
}

type Offset struct {
}

type License struct {
	// The full name of the license. Informational only.
	Name string `xml:"name"`
	// Link to an URL where the full license text can be found.
	Resource string `xml:"resource"`
	// The identifier of the license from the SPDX license list. Can also be an SPDX License Expression, which is also applicable to custom licenses (LicenseRef-…).
	Spdxid string `xml:"spdxid"`
	// The full license text
	Text string `xml:"text"`
}

type DefaultRegulations struct {
	RoadRegulations   *RoadRegulations   `xml:"roadRegulations"`
	SignalRegulations *SignalRegulations `xml:"signalRegulations"`
}

type RoadRegulations struct {
	//
	Type string `xml:"type"`
}

type SignalRegulations struct {
	//
	Subtype string `xml:"subtype"`
	//
	Type string `xml:"type"`
}

type Road struct {
	// Unique ID within the database. If it represents an integer number, it should comply to uint32_t and stay within the given range.
	Id string `xml:"id"`
	// ID of the junction to which the road belongs, for example connecting roads, cross paths, and roads of a junction boundary. Use -1 for none.
	Junction string `xml:"junction"`
	// Total length of the reference line in the xy-plane. Change in length due to elevation is not considered
	Length int `xml:"length"`
	// Name of the road. May be chosen freely.
	Name string `xml:"name"`
	// Basic rule for using the road; RHT=right-hand traffic, LHT=left-hand traffic. When this attribute is missing, RHT is assumed.
	Rule             string            `xml:"rule"`
	Link             *Link             `xml:"link"`
	ElevationProfile *ElevationProfile `xml:""`
	LateralProfile   *LateralProfile   `xml:""`
	Lanes            *Lanes            `xml:""`
	RoadObjects      *RoadObjects      `xml:"objects"`
	Signals          *Signals          `xml:""`
	RoadSurface      *RoadSurface      `xml:"surface"`
	Railroad         *Railroad         `xml:""`
}

type Link struct {
	Predecessor *Predecessor `xml:"predecessor"`
	Successor   *Successor   `xml:"successor"`
}

type Predecessor struct {
	// Contact point of link on the linked element
	ContactPoint string `xml:"contactPoint"`
	// To be provided when elementS is used for the connection definition. Indicates the direction on the predecessor from which the road is entered.
	ElementDir string `xml:"elementDir"`
	// ID of the linked element
	ElementId string `xml:"elementId"`
	// Alternative to contactPoint for virtual junctions. Indicates a connection within the predecessor, meaning not at the start or end of the predecessor. Shall only be used for elementType 'road'
	ElementS int `xml:"elementS"`
	// Type of the linked element
	ElementType string `xml:"elementType"`
}

type Successor struct {
	// Contact point of link on the linked element
	ContactPoint string `xml:"contactPoint"`
	// To be provided when elementS is used for the connection definition. Indicates the direction on the predecessor from which the road is entered.
	ElementDir string `xml:"elementDir"`
	// ID of the linked element
	ElementId string `xml:"elementId"`
	// Alternative to contactPoint for virtual junctions. Indicates a connection within the predecessor, meaning not at the start or end of the predecessor. Shall only be used for elementType 'road'
	ElementS int `xml:"elementS"`
	// Type of the linked element
	ElementType string `xml:"elementType"`
}

type RoadType struct {
	// Country code of the road, see ISO 3166-1, alpha-2 codes.
	Country string `xml:"country"`
	// s-coordinate of start position
	S int `xml:"s"`
	// Type of the road defined as enumeration
	Type  string `xml:"type"`
	Speed *Speed `xml:"speed"`
}

type LaneSpeed struct {
	// Maximum allowed speed. If the attribute unit is not specified, m/s is used as default.
	Max float64 `xml:"max"`
	// s-coordinate of start position, relative to the position of the preceding <laneSection> element
	SOffset float64 `xml:"sOffset"`
	// Unit of the attribute max
	Unit string `xml:"unit"`
}

type Speed struct {
	// Maximum allowed speed. Given as string (only 'no limit' / 'undefined') or numerical value in the respective unit (see attribute unit). If the attribute unit is not specified, m/s is used as default.
	Max int `xml:"max"`
	// Unit of the attribute max. For values, see chapter 'units'.
	Unit string `xml:"unit"`
}

type RoadPlanView struct {
	RoadPlanViewGeometry *RoadPlanViewGeometry `xml:"geometry"`
}

type RoadPlanViewGeometry struct {
	// Start orientation (inertial heading)
	Hdg float64 `xml:"hdg"`
	// Length of the element’s reference line
	Length float64 `xml:"length"`
	// s-coordinate of start position
	S float64 `xml:"s"`
	// Start position (x inertial)
	X float64 `xml:"x"`
	// Start position (y inertial)
	Y          float64     `xml:"y"`
	Line       *Line       `xml:""`
	Spiral     *Spiral     `xml:""`
	Arc        *Arc        `xml:""`
	Poly3      *Poly3      `xml:""`
	ParamPoly3 *ParamPoly3 `xml:""`
}

type Line struct {
	// Line color. If given, this attribute supersedes the definition in the <roadMark> element.
	Color string `xml:"color"`
	// Length of the visible part
	Length float64 `xml:"length"`
	// Rule that must be observed when passing the line from inside, for example, from the lane with the lower absolute ID to the lane with the higher absolute ID
	Rule string `xml:"rule"`
	// Initial longitudinal offset of the line definition from the start of the road mark definition
	SOffset float64 `xml:"sOffset"`
	// Length of the gap between the visible parts
	Space float64 `xml:"space"`
	// Lateral offset from the lane border. If <sway> element is present, the lateral offset follows the sway.
	TOffset float64 `xml:"tOffset"`
	// Line width
	Width float64 `xml:"width"`
}

type ExplicitLine struct {
	// Length of the visible line
	Length float64 `xml:"length"`
	// Rule that must be observed when passing the line from inside, that is, from the lane with the lower absolute ID to the lane with the higher absolute ID
	Rule string `xml:"rule"`
	// Offset of start position of the <line> element, relative to the @sOffset given in the <roadMark> element
	SOffset float64 `xml:"sOffset"`
	// Lateral offset from the lane border. If <sway> element is present, the lateral offset follows the sway.
	TOffset float64 `xml:"tOffset"`
	// Line width. This attribute supersedes the definition in the <roadMark> element.
	Width float64 `xml:"width"`
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
	Elevation *Elevation `xml:"elevation"`
}

type JunctionElevationGridElevation struct {
	// List of defined z-values.
	Center string `xml:"center"`
	// List of defined z-values from inside to outside.
	Left string `xml:"left"`
	// List of defined z-values from inside to outside.
	Right string `xml:"right"`
}

type Elevation struct {
	// Polynom parameter a, elevation at @s (ds=0)
	A float64 `xml:"a"`
	// Polynom parameter b
	B float64 `xml:"b"`
	// Polynom parameter c
	C float64 `xml:"c"`
	// Polynom parameter d
	D float64 `xml:"d"`
	// s-coordinate of start position
	S float64 `xml:"s"`
}

type LateralProfile struct {
	Superelevation      *Superelevation      `xml:""`
	Shape               *Shape               `xml:""`
	CrossSectionSurface *CrossSectionSurface `xml:""`
}

type Superelevation struct {
	// Polynom parameter a, superelevation at @s (ds=0)
	A float64 `xml:"a"`
	// Polynom parameter b
	B float64 `xml:"b"`
	// Polynom parameter c
	C float64 `xml:"c"`
	// Polynom parameter d
	D float64 `xml:"d"`
	// s-coordinate of start position
	S float64 `xml:"s"`
}

type Shape struct {
	// Polynom parameter a, relative height at @t (dt=0)
	A float64 `xml:"a"`
	// Polynom parameter b
	B float64 `xml:"b"`
	// Polynom parameter c
	C float64 `xml:"c"`
	// Polynom parameter d
	D float64 `xml:"d"`
	// s-coordinate of start position
	S float64 `xml:"s"`
	// t-coordinate of start position
	T float64 `xml:"t"`
}

type CrossSectionSurface struct {
	TOffset       *TOffset       `xml:""`
	SurfaceStrips *SurfaceStrips `xml:""`
}

type TOffset struct {
}

type Coefficients struct {
	// Polynomial parameter a. If the attribute is not specified, the value is 0.
	A float64 `xml:"a"`
	// Polynomial parameter b. If the attribute is not specified, the value is 0.
	B float64 `xml:"b"`
	// Polynomial parameter c. If the attribute is not specified, the value is 0.
	C float64 `xml:"c"`
	// Polynomial parameter d. If the attribute is not specified, the value is 0.
	D float64 `xml:"d"`
	// s-coordinate of start position
	S float64 `xml:"s"`
}

type SurfaceStrips struct {
}

type Strip struct {
	// 1 for the inner left strip, -1 for the inner right strip, 2 for the outer left strip, -2 for the outer right strip
	Id int `xml:"id"`
	// Can only be defined for an outer strip.
	Mode string `xml:"mode"`
}

type Width struct {
	// Polynom parameter a, width at @s (ds=0)
	A float64 `xml:"a"`
	// Polynom parameter b
	B float64 `xml:"b"`
	// Polynom parameter c
	C float64 `xml:"c"`
	// Polynom parameter d
	D float64 `xml:"d"`
	// s-coordinate of start position of the <width> element, relative to the position of the preceding <laneSection> element
	SOffset float64 `xml:"sOffset"`
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
	LaneOffset  *LaneOffset  `xml:""`
	LaneSection *LaneSection `xml:""`
}

type LaneOffset struct {
	// Polynom parameter a, offset at @s (ds=0)
	A float64 `xml:"a"`
	// Polynom parameter b
	B float64 `xml:"b"`
	// Polynom parameter c
	C float64 `xml:"c"`
	// Polynom parameter d
	D float64 `xml:"d"`
	// s-coordinate of start position
	S float64 `xml:"s"`
}

type LaneSection struct {
	// s-coordinate of start position
	S float64 `xml:"s"`
	// Lane section element is valid for one side only (left, center, or right), depending on the child elements.
	SingleSide string  `xml:"singleSide"`
	Left       *Left   `xml:""`
	Center     *Center `xml:""`
	Right      *Right  `xml:""`
}

type Left struct {
	Lane *Lane `xml:""`
}

type Lane struct {
	// If true, lane can be used also by a neighboring lane. Advisory lane has priority, for example a bike lane, that can also be used by cars. If not specified, default value is none.
	Advisory string `xml:"advisory"`
	// If not specified, direction is determined by the combination of <left> or <right> lane grouping and the values LHT or RHT of the @rule attribute of a road. The standard direction can be overwritten with this attribute.
	Direction string `xml:"direction"`
	// If true, lane direction can be changed dynamically by the scenario during the simulation. If not specified, default boolean value is false.
	DynamicLaneDirection string `xml:"dynamicLaneDirection"`
	// If true, lane type can be changed dynamically by the scenario during the simulation. Typical example is a stop lane that can be changed by VMS boards to a driving lane. If not specified, default boolean value is false.
	DynamicLaneType string `xml:"dynamicLaneType"`
	// ID of the lane
	Id int `xml:"id"`
	// 'true' = keep lane on level, that is, do not apply superelevation; 'false' = apply superelevation to this lane (default, also used if attribute level is missing)
	Level string `xml:"level"`
	// If true, lane is under construction.
	RoadWorks string `xml:"roadWorks"`
	// Type of the lane
	Type     string    `xml:"type"`
	Width    *Width    `xml:""`
	Border   *Border   `xml:""`
	Link     *Link     `xml:"link"`
	RoadMark *RoadMark `xml:""`
	Material *Material `xml:""`
	Speed    *Speed    `xml:"speed"`
	Access   *Access   `xml:""`
	Height   *Height   `xml:""`
	Rule     *Rule     `xml:""`
}

type Border struct {
	// Polynom parameter a, width at @s (ds=0)
	A float64 `xml:"a"`
	// Polynom parameter b
	B float64 `xml:"b"`
	// Polynom parameter c
	C float64 `xml:"c"`
	// Polynom parameter d
	D float64 `xml:"d"`
	// s-coordinate of start position of the <border> element , relative to the position of the preceding <laneSection> element
	SOffset         float64          `xml:"sOffset"`
	CornerReference *CornerReference `xml:""`
}

type RoadMark struct {
	// Color of the road mark
	Color string `xml:"color"`
	// Height of road mark above the road, i.e. thickness of the road mark
	Height float64 `xml:"height"`
	// Allows a lane change in the indicated direction, taking into account that lanes are numbered in ascending order from right to left. If the attribute is missing, “both” is used as default.
	LaneChange string `xml:"laneChange"`
	// Material of the road mark. Identifiers to be defined by the user, use 'standard' as default value.
	Material string `xml:"material"`
	// s-coordinate of start position of the <roadMark> element, relative to the position of the preceding <laneSection> element
	SOffset float64 `xml:"sOffset"`
	// Type of the road mark
	Type string `xml:"type"`
	// Weight of the road mark. This attribute is optional if detailed definition is given below.
	Weight string `xml:"weight"`
	// Width of the road mark. This attribute is optional if detailed definition is given by <line> element.
	Width        float64       `xml:"width"`
	Sway         *Sway         `xml:""`
	RoadMarkType *RoadMarkType `xml:""`
	Explicit     *Explicit     `xml:""`
}

type Sway struct {
	// Polynom parameter a, sway value at @s (ds=0)
	A float64 `xml:"a"`
	// Polynom parameter b
	B float64 `xml:"b"`
	// Polynom parameter c
	C float64 `xml:"c"`
	// Polynom parameter d
	D float64 `xml:"d"`
	// s-coordinate of start position of the <sway> element, relative to the @sOffset given in the <roadMark> element
	Ds float64 `xml:"ds"`
}

type RoadMarkType struct {
	// Name of the road mark type. May be chosen freely.
	Name string `xml:"name"`
	// Accumulated width of the road mark. In case of several <line> elements this @width is the sum of all @width of <line> elements and spaces in between, necessary to form the road mark. This attribute supersedes the definition in the <roadMark> element.
	Width float64 `xml:"width"`
	Line  *Line   `xml:""`
}

type Explicit struct {
	Line *Line `xml:""`
}

type Material struct {
	// Friction coefficient
	Friction float64 `xml:"friction"`
	// Roughness, for example, for sound and motion systems
	Roughness float64 `xml:"roughness"`
	// s-coordinate of start position, relative to the position of the preceding <laneSection> element
	SOffset float64 `xml:"sOffset"`
	// Surface material code, depending on application
	Surface string `xml:"surface"`
}

type Access struct {
	// Identifier of the participant to whom the restriction applies
	Restriction string `xml:"restriction"`
	// Specifies whether the participant given in the attribute @restriction is allowed or denied access to the given lane
	Rule string `xml:"rule"`
	// s-coordinate of start position, relative to the position of the preceding <laneSection> element
	SOffset float64 `xml:"sOffset"`
}

type Restriction struct {
	// Identifier of the participant to whom the restriction applies
	Type string `xml:"type"`
}

type Height struct {
	// Inner offset from road level
	Inner float64 `xml:"inner"`
	// Outer offset from road level
	Outer float64 `xml:"outer"`
	// s-coordinate of start position, relative to the position of the preceding <laneSection> element
	SOffset float64 `xml:"sOffset"`
}

type Rule struct {
	// s-coordinate of start position, relative to the position of the preceding <laneSection> element
	SOffset float64 `xml:"sOffset"`
	// Free text; currently recommended values are 'no stopping at any time', 'disabled parking', 'car pool'
	Value string `xml:"value"`
}

type Center struct {
}

type Right struct {
}

type RoadObjects struct {
	RoadObjectsObject *RoadObjectsObject `xml:"object"`
	ObjectReference   *ObjectReference   `xml:""`
	Tunnel            *Tunnel            `xml:""`
	Bridge            *Bridge            `xml:""`
}

type RoadObjectsObject struct {
	// Indicates whether the object is dynamic or static, default value is 'no' (static). Dynamic object cannot change its position.
	Dynamic string `xml:"dynamic"`
	// Heading angle of the object relative to road direction
	Hdg float64 `xml:"hdg"`
	// Height of the object’s bounding box. @height is defined in the local coordinate system u/v along the z-axis
	Height float64 `xml:"height"`
	// Unique ID within database
	Id string `xml:"id"`
	// Length of the object's bounding box, alternative to @radius. @length is defined in the local coordinate system u/v along the u-axis
	Length float64 `xml:"length"`
	// Name of the object. May be chosen freely.
	Name string `xml:"name"`
	// '+' = valid in positive s-direction '-' = valid in negative s-direction 'none' = valid in both directions (does not affect the heading)
	Orientation string `xml:"orientation"`
	// Alternative to @pitch and @roll. If true, the object is vertically perpendicular to the road surface at all points and @pitch and @roll are ignored. Default is false.
	PerpToRoad bool `xml:"perpToRoad"`
	// Pitch angle relative to the x/y-plane
	Pitch float64 `xml:"pitch"`
	// radius of the circular object's bounding box, alternative to @length and @width. @radius is defined in the local coordinate system u/v
	Radius float64 `xml:"radius"`
	// Roll angle relative to the x/y-plane
	Roll float64 `xml:"roll"`
	// s-coordinate of object's origin
	S float64 `xml:"s"`
	// Variant of a type
	Subtype string `xml:"subtype"`
	// t-coordinate of object's origin
	T float64 `xml:"t"`
	// Type of object. For a parking space, the <parkingSpace> element may be used additionally.
	Type string `xml:"type"`
	// Validity of object along s-axis (0.0 for point object)
	ValidLength float64 `xml:"validLength"`
	// Width of the object's bounding box, alternative to @radius. @width is defined in the local coordinate system u/v along the v-axis
	Width float64 `xml:"width"`
	// z-offset of object's origin relative to the elevation of the road reference line
	ZOffset      float64       `xml:"zOffset"`
	Repeat       *Repeat       `xml:""`
	Outline      *Outline      `xml:""`
	Outlines     *Outlines     `xml:""`
	Material     *Material     `xml:""`
	Validity     *Validity     `xml:""`
	ParkingSpace *ParkingSpace `xml:""`
	Markings     *Markings     `xml:""`
	Borders      *Borders      `xml:""`
	RoadSurface  *RoadSurface  `xml:"surface"`
	Skeleton     *Skeleton     `xml:""`
}

type Repeat struct {
}

type Outline struct {
	CornerRoad  *CornerRoad  `xml:""`
	CornerLocal *CornerLocal `xml:""`
}

type CornerRoad struct {
}

type CornerLocal struct {
}

type Outlines struct {
	Outline *Outline `xml:""`
}

type Validity struct {
}

type ParkingSpace struct {
}

type Markings struct {
	Marking *Marking `xml:""`
}

type Marking struct {
	CornerReference *CornerReference `xml:""`
}

type CornerReference struct {
}

type Borders struct {
	Border *Border `xml:""`
}

type RoadSurface struct {
	RoadSurfaceCRG *RoadSurfaceCRG `xml:"CRG"`
}

type Skeleton struct {
	Polyline *Polyline `xml:""`
}

type Polyline struct {
	VertexRoad  *VertexRoad  `xml:""`
	VertexLocal *VertexLocal `xml:""`
}

type VertexRoad struct {
}

type VertexLocal struct {
}

type ObjectReference struct {
	Validity *Validity `xml:""`
}

type Tunnel struct {
	Validity *Validity `xml:""`
}

type Bridge struct {
	Validity *Validity `xml:""`
}

type Signals struct {
	Signal          *Signal          `xml:""`
	SignalReference *SignalReference `xml:""`
}

type Signal struct {
	PositionInertial *PositionInertial `xml:""`
	PositionRoad     *PositionRoad     `xml:""`
	Validity         *Validity         `xml:""`
	Dependency       *Dependency       `xml:""`
	Reference        *Reference        `xml:""`
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
	Validity *Validity `xml:""`
}

type RoadSurfaceCRG struct {
	// Name of the file containing the CRG data
	File string `xml:"file"`
	// Heading offset between CRG center line and reference line of the road (only allowed for mode genuine, default = 0.0).
	HOffset float64 `xml:"hOffset"`
	// Attachment mode for the surface data, see specification.
	Mode string `xml:"mode"`
	// Orientation of the CRG data set relative to the parent <road> element. Only allowed for mode attached and attached0.
	Orientation string `xml:"orientation"`
	// Physical purpose of the data contained in the CRG file; if the attribute is missing, data will be interpreted as elevation data.
	Purpose string `xml:"purpose"`
	// End of the application of CRG (s-coordinate)
	SEnd float64 `xml:"sEnd"`
	// s-offset between CRG center line and reference line of the road (default = 0.0)
	SOffset float64 `xml:"sOffset"`
	// Start of the application of CRG data (s-coordinate)
	SStart float64 `xml:"sStart"`
	// t-offset between CRG center line and reference line of the road (default = 0.0)
	TOffset float64 `xml:"tOffset"`
	// z-offset between CRG center line and reference line of the road (default = 0.0). Only allowed for purpose elevation.
	ZOffset float64 `xml:"zOffset"`
	// z-scale factor for the surface description (default = 1.0). Only allowed for purpose elevation.
	ZScale float64 `xml:"zScale"`
}

type Railroad struct {
	Switch *Switch `xml:""`
}

type Switch struct {
	MainTrack *MainTrack `xml:""`
	SideTrack *SideTrack `xml:""`
	Partner   *Partner   `xml:""`
}

type MainTrack struct {
}

type SideTrack struct {
}

type Partner struct {
}

type JunctionController struct {
	// ID of the controller
	Id string `xml:"id"`
	// Sequence number (priority) of this controller with respect to other controllers in the same junction
	Sequence int `xml:"sequence"`
	// Type of control for this junction. Free text, depending on the application.
	Type    string   `xml:"type"`
	Control *Control `xml:""`
}

type Control struct {
}

type JunctionCrossing struct {
	JunctionRoadSection *JunctionRoadSection `xml:"roadSection"`
	Priority            *Priority            `xml:""`
	JunctionController  *JunctionController  `xml:"controller"`
	RoadSurface         *RoadSurface         `xml:"surface"`
	RoadPlanView        *RoadPlanView        `xml:"planView"`
}

type JunctionRoadSection struct {
	// Unique ID within the junction
	Id string `xml:"id"`
	// ID of the road of this roadSection element
	RoadId string `xml:"roadId"`
	// End position of the crossing junction in the road reference line coordinate system. This attribute is mandatory for crossing junctions.
	SEnd float64 `xml:"sEnd"`
	// Start position of the crossing junction in the road reference line coordinate system. This attribute is mandatory for crossing junctions.
	SStart float64 `xml:"sStart"`
}

type Priority struct {
	// ID of the prioritized road
	High string `xml:"high"`
	// ID of the road with lower priority
	Low string `xml:"low"`
}

type JunctionDefault struct {
	// ID of the junction to which the road belongs, for example connecting roads, cross paths, and roads of a junction boundary. Use -1 for none.
	Id string `xml:"id"`
	// Name of the junction. May be chosen freely.
	Name string `xml:"name"`
	// Common junctions are of type 'default'. If the attribute is not specified, the junction type is 'default'. This attribute is mandatory for all other junction types.
	Type                  string                 `xml:"type"`
	Connection            *Connection            `xml:""`
	CrossPath             *CrossPath             `xml:""`
	Priority              *Priority              `xml:""`
	JunctionController    *JunctionController    `xml:"controller"`
	RoadSurface           *RoadSurface           `xml:"surface"`
	RoadPlanView          *RoadPlanView          `xml:"planView"`
	JunctionBoundary      *JunctionBoundary      `xml:"boundary"`
	JunctionElevationGrid *JunctionElevationGrid `xml:"elevationGrid"`
}

type JunctionConnectionDirect struct {
	// Contact point on the @connectingRoad or @linkedRoad. Required for all junction types except virtual.
	ContactPoint string `xml:"contactPoint"`
	// Unique ID within the junction
	Id string `xml:"id"`
	// ID of the incoming road. Required for all junction types except virtual.
	IncomingRoad string `xml:"incomingRoad"`
	// ID of the directly linked road. Only to be used for junctions of @type='direct'.
	LinkedRoad                 string                      `xml:"linkedRoad"`
	JunctionConnectionLaneLink *JunctionConnectionLaneLink `xml:"laneLink"`
}

type Connection struct {
	// ID of the connecting road. Only to be used for junctions of @type='default'.
	ConnectingRoad string `xml:"connectingRoad"`
	// Contact point on the @connectingRoad or @linkedRoad. Required for all junction types except virtual.
	ContactPoint string `xml:"contactPoint"`
	// Unique ID within the junction
	Id string `xml:"id"`
	// ID of the incoming road. Required for all junction types except virtual.
	IncomingRoad               string                      `xml:"incomingRoad"`
	JunctionConnectionLaneLink *JunctionConnectionLaneLink `xml:"laneLink"`
}

type JunctionConnectionLaneLink struct {
	// ID of the incoming lane
	From int `xml:"from"`
	// Specifies the length of the area where traffic from both overlapping lanes shares the space. It is defined in s length relative to the position of the junction. Intended for direct junctions only. Default is 100.
	OverlapZone float64 `xml:"overlapZone"`
	// ID of the connection lane
	To int `xml:"to"`
}

type CrossPath struct {
	// ID of road defining the cross path.
	CrossingRoad string `xml:"crossingRoad"`
	// Unique ID within the junction
	Id string `xml:"id"`
	// ID of road at end point of the crossing road
	RoadAtEnd string `xml:"roadAtEnd"`
	// ID of road at start point of the crossing road
	RoadAtStart   string         `xml:"roadAtStart"`
	StartLaneLink *StartLaneLink `xml:""`
	EndLaneLink   *EndLaneLink   `xml:""`
}

type StartLaneLink struct {
	// Lane ID of either @roadAtEnd for <endLaneLink> or @roadAtStart for <startLaneLink>
	From int `xml:"from"`
	// s-coordinate of either start or end point in linked road.
	S float64 `xml:"s"`
	// Lane ID of @crossingRoad
	To int `xml:"to"`
}

type EndLaneLink struct {
	// Lane ID of either @roadAtEnd for <endLaneLink> or @roadAtStart for <startLaneLink>
	From int `xml:"from"`
	// s-coordinate of either start or end point in linked road.
	S float64 `xml:"s"`
	// Lane ID of @crossingRoad
	To int `xml:"to"`
}

type JunctionBoundary struct {
	JunctionBoundarySegmentJoint *JunctionBoundarySegmentJoint `xml:"segment type='joint'"`
	JunctionBoundarySegmentLane  *JunctionBoundarySegmentLane  `xml:"segment type='lane'"`
}

type JunctionBoundarySegmentLane struct {
	// ID of the lane of which the outer edge is the segment
	BoundaryLane int `xml:"boundaryLane"`
	// ID of the road used for the segment
	RoadId string `xml:"roadId"`
	// End of the segment (s-coordinate, begin, end)
	SEnd float64 `xml:"sEnd"`
	// Start of the segment (s-coordinate, begin, end)
	SStart float64 `xml:"sStart"`
	// Type of the segment
	Type string `xml:"type"`
}

type JunctionBoundarySegmentJoint struct {
	// Contact point on the road
	ContactPoint float64 `xml:"contactPoint"`
	// ID of the lane crossed by the segment. If missing all lanes are crossed by the segment.
	JointLaneEnd int `xml:"jointLaneEnd"`
	// ID of the lane crossed by the segment. If missing all lanes are crossed by the segment.
	JointLaneStart int `xml:"jointLaneStart"`
	// ID of the road used for the segment
	RoadId string `xml:"roadId"`
	// Length of the transition area where local height is interpolated between road data and the <elevationGrid> in order to ensure a smooth transition. The default is 0.
	TransitionLength float64 `xml:"transitionLength"`
	// Type of the segment
	Type string `xml:"type"`
}

type JunctionElevationGrid struct {
	//
	GridSpacing float64 `xml:"gridSpacing"`
	//
	SStart                         float64                         `xml:"sStart"`
	JunctionElevationGridElevation *JunctionElevationGridElevation `xml:"elevation"`
}

type JunctionDirect struct {
	// ID of the junction to which the road belongs, for example connecting roads, cross paths, and roads of a junction boundary. Use -1 for none.
	Id string `xml:"id"`
	// Name of the junction. May be chosen freely.
	Name string `xml:"name"`
	// Common junctions are of type 'default'. If the attribute is not specified, the junction type is 'default'. This attribute is mandatory for all other junction types.
	Type               string              `xml:"type"`
	Connection         *Connection         `xml:""`
	Priority           *Priority           `xml:""`
	JunctionController *JunctionController `xml:"controller"`
	RoadSurface        *RoadSurface        `xml:"surface"`
	RoadPlanView       *RoadPlanView       `xml:"planView"`
}

type JunctionVirtual struct {
	// ID of the junction to which the road belongs, for example connecting roads, cross paths, and roads of a junction boundary. Use -1 for none.
	Id string `xml:"id"`
	// The main road from which the connecting roads of the virtual junction branch off. This attribute is mandatory for virtual junctions and shall not be specified for other junction types.
	MainRoad string `xml:"mainRoad"`
	// Name of the junction. May be chosen freely.
	Name string `xml:"name"`
	// Defines the relevance of the virtual junction according to the driving direction. This attribute is mandatory for virtual junctions and shall not be specified for other junction types. The enumerator 'none' specifies that the virtual junction is valid in both directions.
	Orientation string `xml:"orientation"`
	// End position of the virtual junction in the reference line coordinate system. This attribute is mandatory for virtual junctions.
	SEnd float64 `xml:"sEnd"`
	// Start position of the virtual junction in the reference line coordinate system. This attribute is mandatory for virtual junctions.
	SStart float64 `xml:"sStart"`
	// Common junctions are of type 'default'. If the attribute is not specified, the junction type is 'default'. This attribute is mandatory for all other junction types.
	Type                      string                     `xml:"type"`
	JunctionConnectionDefault *JunctionConnectionDefault `xml:"connection type='default'"`
	JunctionConnectionVirtual *JunctionConnectionVirtual `xml:"connection type='virtual'"`
	CrossPath                 *CrossPath                 `xml:""`
	Priority                  *Priority                  `xml:""`
	JunctionController        *JunctionController        `xml:"controller"`
	RoadSurface               *RoadSurface               `xml:"surface"`
	RoadPlanView              *RoadPlanView              `xml:"planView"`
}

type JunctionConnectionDefault struct {
	//
	ConnectingRoad string `xml:"connectingRoad"`
	// Contact point on the @connectingRoad or @linkedRoad. Required for all junction types except virtual.
	ContactPoint string `xml:"contactPoint"`
	// Unique ID within the junction
	Id string `xml:"id"`
	// ID of the incoming road. Required for all junction types except virtual.
	IncomingRoad string `xml:"incomingRoad"`
	// Type of the connection. Regular connections are @type=“default” . This attribute is mandatory for virtual connections.
	Type                       string                      `xml:"type"`
	JunctionConnectionLaneLink *JunctionConnectionLaneLink `xml:"laneLink"`
}

type JunctionConnectionVirtual struct {
	//
	ConnectingRoad string `xml:"connectingRoad"`
	// Contact point on the @connectingRoad or @linkedRoad. Required for all junction types except virtual.
	ContactPoint string `xml:"contactPoint"`
	// Unique ID within the junction
	Id string `xml:"id"`
	// ID of the incoming road. Required for all junction types except virtual.
	IncomingRoad string `xml:"incomingRoad"`
	// Type of the connection. Regular connections are @type=“default” . This attribute is mandatory for virtual connections.
	Type                       string                      `xml:"type"`
	Predecessor                *Predecessor                `xml:"predecessor"`
	Successor                  *Successor                  `xml:"successor"`
	JunctionConnectionLaneLink *JunctionConnectionLaneLink `xml:"laneLink"`
}

type JunctionGroup struct {
	// Unique ID within database
	Id string `xml:"id"`
	// Name of the junction group. May be chosen freely.
	Name string `xml:"name"`
	// Type of junction group
	Type                           string                          `xml:"type"`
	JunctionGroupJunctionReference *JunctionGroupJunctionReference `xml:"junctionReference"`
}

type JunctionGroupJunctionReference struct {
	// ID of the junction
	Junction string `xml:"junction"`
}

type Station struct {
	Platform *Platform `xml:""`
}

type Platform struct {
	Segment *Segment `xml:""`
}

type Segment struct {
}
