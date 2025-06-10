// Code generated; DO NOT EDIT.

package xodr

type EBorderType string

type EBridgeType string

type EObjectType string

type EOrientation string

type EOutlineFillType string

type ERoadObjectsObjectParkingSpaceAccess string

type ESideType string

type ETunnelType string

// Container for all objects along a road.
type RoadObjects struct {
	OpenDriveElement
	Object          []*RoadObjectsObject
	ObjectReference []*RoadObjectsObjectReference
	Tunnel          []*RoadObjectsTunnel
	Bridge          []*RoadObjectsBridge
}

// Bridges are modeled as objects in ASAM OpenDRIVE. The road with the bridge
// object leads over a bridge. Bridges are valid for a roadâ  s complete cros
// s section unless a lane validity record with further restrictions is provid
// ed as child element.
type RoadObjectsBridge struct {
	OpenDriveElement
	Validity []*RoadObjectsObjectLaneValidity
	Id       string
	Length   GrEqZero
	Name     string
	S        GrEqZero
	Type     EBridgeType
}

// Objects influence a road by expanding, delimiting, or supplementing its cou
// rse. Objects are elements that form the environment, for example, buildings
// , guard rails, poles, and trees. Objects do not influence the behavior of t
// raffic directly, unlike signals. There are two ways to describe the boundin
// g box of objects.  - For an angular object: definition of the width, length
// and height.  - For a circular object: definition of the radius and height.
type RoadObjectsObject struct {
	OpenDriveElement
	Repeat       []*RoadObjectsObjectRepeat
	Outline      *RoadObjectsObjectOutlinesOutline
	Outlines     *RoadObjectsObjectOutlines
	Material     []*RoadObjectsObjectMaterial
	Validity     []*RoadObjectsObjectLaneValidity
	ParkingSpace *RoadObjectsObjectParkingSpace
	Markings     *RoadObjectsObjectMarkings
	Borders      *RoadObjectsObjectBorders
	Surface      *RoadObjectsObjectSurface
	Skeleton     *RoadObjectsObjectSkeleton
	Dynamic      YesNo
	Hdg          float64
	Height       GrEqZero
	Id           string
	Length       GrZero
	Name         string
	Orientation  EOrientation
	PerpToRoad   Bool
	Pitch        float64
	Radius       GrZero
	Roll         float64
	S            GrEqZero
	Subtype      string
	T            float64
	Type         EObjectType
	ValidLength  GrEqZero
	Width        float64
	ZOffset      float64
}

// Object borders are frames with a defined width, for example, to describe tr
// affic islands. Different border types are available.
type RoadObjectsObjectBorders struct {
	OpenDriveElement
	Border []*RoadObjectsObjectBordersBorder
}

// Specifies a border along certain outline points.
type RoadObjectsObjectBordersBorder struct {
	OpenDriveElement
	CornerReference    []*RoadObjectsObjectMarkingsMarkingCornerReference
	OutlineId          int
	Type               EBorderType
	UseCompleteOutline Bool
	Width              GrEqZero
}

// Object markings are road markings of any objects, for example, crosswalks,
// stopping-lines, and parking spaces.
type RoadObjectsObjectMarkings struct {
	OpenDriveElement
	Marking []*RoadObjectsObjectMarkingsMarking
}

// Specifies a marking that is either attached to one side of the object bound
// ing box or referencing outline points.
type RoadObjectsObjectMarkingsMarking struct {
	OpenDriveElement
	CornerReference []*RoadObjectsObjectMarkingsMarkingCornerReference
	Color           ERoadMarkColor
	LineLength      GrZero
	Side            ESideType
	SpaceLength     GrEqZero
	StartOffset     float64
	StopOffset      float64
	Weight          ERoadMarkWeight
	Width           GrZero
	ZOffset         GrEqZero
}

// Specifies a point by referencing an existing outline point.
type RoadObjectsObjectMarkingsMarkingCornerReference struct {
	OpenDriveElement
	Id int
}

// Describes the material properties of objects, for example, patches that are
// part of the road surface but deviate from the standard road material. Super
// sedes the material specified in the <road material> element and is valid on
// ly within the outline of the parent road object.
type RoadObjectsObjectMaterial struct {
	OpenDriveElement
	Friction      GrEqZero
	RoadMarkColor ERoadMarkColor
	Roughness     GrEqZero
	Surface       string
}

// Wrapper for the different outline entries of an object, that can contain mu
// ltiple outline definitions. If <outlines> is not used, an object can have o
// nly a single <outline> entry.
type RoadObjectsObjectOutlines struct {
	OpenDriveElement
	Outline []*RoadObjectsObjectOutlinesOutline
}

// Defines a series of corner points, including the height of the object relat
// ive to the road reference line. For areas, the points should be listed in c
// ounter-clockwise order. The inner area of the described outline may be fill
// ed with a filling type, such as grass, concrete, asphalt, or pavement. An e
// lement shall be followed by two or more <cornerRoad> elements or by two or
// more <cornerLocal> elements. ASAM OpenDRIVE 1.4 outline definitions (withou
// t <outlines> parent element) shall still be supported.
type RoadObjectsObjectOutlinesOutline struct {
	OpenDriveElement
	Closed   Bool
	FillType EOutlineFillType
	Id       int
	LaneType ELaneType
	Outer    Bool
}

// Used to describe complex forms of objects. Defines a corner point on the ob
// ject outline relative to the object pivot point in local u/v-coordinates. T
// he insertion point and the orientation of the object are given by the @s, @
// t, @zOffset and @hdg attributes of the  element.
type RoadObjectsObjectOutlinesOutlineCornerLocal struct {
	OpenDriveElement
	Height GrEqZero
	Id     int
	U      float64
	V      float64
	Z      float64
}

// Defines a corner point on the objectâ  s outline in road coordinates.
type RoadObjectsObjectOutlinesOutlineCornerRoad struct {
	OpenDriveElement
	Dz     float64
	Height GrEqZero
	Id     int
	S      GrEqZero
	T      float64
}

// Details for a parking space may be added to the <object> element.
type RoadObjectsObjectParkingSpace struct {
	OpenDriveElement
	Access       ERoadObjectsObjectParkingSpaceAccess
	Restrictions string
}

// To avoid lengthy XML code, objects of the same type may be repeated. Attrib
// utes of the repeated object shall overrule the attributes from the original
// object. If attributes are omitted in the repeated objects, the attributes f
// rom the original object apply.
type RoadObjectsObjectRepeat struct {
	OpenDriveElement
	DetachFromReferenceLine Bool
	Distance                GrEqZero
	HeightEnd               GrEqZero
	HeightStart             GrEqZero
	Length                  GrEqZero
	LengthEnd               GrEqZero
	LengthStart             GrEqZero
	RadiusEnd               GrEqZero
	RadiusStart             GrEqZero
	S                       GrEqZero
	TEnd                    float64
	TStart                  float64
	WidthEnd                GrEqZero
	WidthStart              GrEqZero
	ZOffsetEnd              float64
	ZOffsetStart            float64
}

// Wrapper for the object polylines, that can be used to describe the actual s
// hape inside the bounding box more closely
type RoadObjectsObjectSkeleton struct {
}

// Defines a series of points relative to the road reference line.  An <polyli
// ne> element shall be followed by either two or more <vertexRoad> elements o
// r by two or more <vertexLocal> elements.
type RoadObjectsObjectSkeletonPolyline struct {
	Id int
}

// Defines a vertex point on the object polyline relative to the object pivot
// point in local u/v-coordinates. The insertion point and the orientation of
// the object are given by the @s, @t, @zOffset and @hdg attributes of the ele
// ment. <vertexLocal> can use either radius or length/width within one <polyl
// ine> element.
type RoadObjectsObjectSkeletonPolylineVertexLocal struct {
	Id                int
	IntersectionPoint Bool
	Radius            float64
	U                 float64
	V                 float64
	Z                 float64
}

// Defines a point on the objectâ  s polyline in road coordinates. <vertexRoa
// d> can use either radius or length/width within one <polyline> element.
type RoadObjectsObjectSkeletonPolylineVertexRoad struct {
	Dz                float64
	Id                int
	IntersectionPoint Bool
	Radius            float64
	S                 GrEqZero
	T                 float64
}

// Used to describe the road surface elevation of an object.
type RoadObjectsObjectSurface struct {
	OpenDriveElement
	Crg *RoadObjectsObjectSurfaceCrg
}

// Elevation data described in {GLO_VAR_STA_ASAM_OpenCRG} are represented by t
// he <CRG> element within the <surface> element.
type RoadObjectsObjectSurfaceCrg struct {
	OpenDriveElement
	File               string
	HideRoadSurfaceCrg Bool
	ZScale             float64
}

// An object reference refers to one identical object from multiple roads. The
// referenced objects require a unique ID. The <objectReference> element consi
// sts of a main element and an optional lane validity element.
type RoadObjectsObjectReference struct {
	OpenDriveElement
	Validity    []*RoadObjectsObjectLaneValidity
	Id          string
	Orientation EOrientation
	S           GrEqZero
	T           float64
	ValidLength GrEqZero
	ZOffset     float64
}

// Tunnels are modeled as objects in ASAM OpenDRIVE. The road with the tunnel
// object defines the part of the road that is the tunnel or the underpass. Tu
// nnels are valid for the complete cross section of a road unless a lane vali
// dity element with further restrictions is provided as child element
type RoadObjectsTunnel struct {
	OpenDriveElement
	Validity []*RoadObjectsObjectLaneValidity
	Daylight ZeroOne
	Id       string
	Length   GrEqZero
	Lighting ZeroOne
	Name     string
	S        GrEqZero
	Type     ETunnelType
}
