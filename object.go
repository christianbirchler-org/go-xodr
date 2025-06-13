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
	Object          []*RoadObjectsObject          `xml:"object"`
	ObjectReference []*RoadObjectsObjectReference `xml:"objectReference"`
	Tunnel          []*RoadObjectsTunnel          `xml:"tunnel"`
	Bridge          []*RoadObjectsBridge          `xml:"bridge"`
}

// Bridges are modeled as objects in ASAM OpenDRIVE. The road with the bridge
// object leads over a bridge. Bridges are valid for a roadâ  s complete cros
// s section unless a lane validity record with further restrictions is provid
// ed as child element.
type RoadObjectsBridge struct {
	OpenDriveElement
	Validity []*RoadObjectsObjectLaneValidity `xml:"validity"`
	Id       string                           `xml:"id,attr"`
	Length   GrEqZero                         `xml:"length,attr"`
	Name     string                           `xml:"name,attr"`
	S        GrEqZero                         `xml:"s,attr"`
	Type     EBridgeType                      `xml:"type,attr"`
}

// Objects influence a road by expanding, delimiting, or supplementing its cou
// rse. Objects are elements that form the environment, for example, buildings
// , guard rails, poles, and trees. Objects do not influence the behavior of t
// raffic directly, unlike signals. There are two ways to describe the boundin
// g box of objects.  - For an angular object: definition of the width, length
// and height.  - For a circular object: definition of the radius and height.
type RoadObjectsObject struct {
	OpenDriveElement
	Repeat       []*RoadObjectsObjectRepeat        `xml:"repeat"`
	Outline      *RoadObjectsObjectOutlinesOutline `xml:"outline"`
	Outlines     *RoadObjectsObjectOutlines        `xml:"outlines"`
	Material     []*RoadObjectsObjectMaterial      `xml:"material"`
	Validity     []*RoadObjectsObjectLaneValidity  `xml:"validity"`
	ParkingSpace *RoadObjectsObjectParkingSpace    `xml:"parkingSpace"`
	Markings     *RoadObjectsObjectMarkings        `xml:"markings"`
	Borders      *RoadObjectsObjectBorders         `xml:"borders"`
	Surface      *RoadObjectsObjectSurface         `xml:"surface"`
	Skeleton     *RoadObjectsObjectSkeleton        `xml:"skeleton"`
	Dynamic      YesNo                             `xml:"dynamic,attr"`
	Hdg          float64                           `xml:"hdg,attr"`
	Height       GrEqZero                          `xml:"height,attr"`
	Id           string                            `xml:"id,attr"`
	Length       GrZero                            `xml:"length,attr"`
	Name         string                            `xml:"name,attr"`
	Orientation  EOrientation                      `xml:"orientation,attr"`
	PerpToRoad   Bool                              `xml:"perpToRoad,attr"`
	Pitch        float64                           `xml:"pitch,attr"`
	Radius       GrZero                            `xml:"radius,attr"`
	Roll         float64                           `xml:"roll,attr"`
	S            GrEqZero                          `xml:"s,attr"`
	Subtype      string                            `xml:"subtype,attr"`
	T            float64                           `xml:"t,attr"`
	Type         EObjectType                       `xml:"type,attr"`
	ValidLength  GrEqZero                          `xml:"validLength,attr"`
	Width        float64                           `xml:"width,attr"`
	ZOffset      float64                           `xml:"zOffset,attr"`
}

// Object borders are frames with a defined width, for example, to describe tr
// affic islands. Different border types are available.
type RoadObjectsObjectBorders struct {
	OpenDriveElement
	Border []*RoadObjectsObjectBordersBorder `xml:"border"`
}

// Specifies a border along certain outline points.
type RoadObjectsObjectBordersBorder struct {
	OpenDriveElement
	CornerReference    []*RoadObjectsObjectMarkingsMarkingCornerReference `xml:"cornerReference"`
	OutlineId          int                                                `xml:"outlineId,attr"`
	Type               EBorderType                                        `xml:"type,attr"`
	UseCompleteOutline Bool                                               `xml:"useCompleteOutline,attr"`
	Width              GrEqZero                                           `xml:"width,attr"`
}

// Object markings are road markings of any objects, for example, crosswalks,
// stopping-lines, and parking spaces.
type RoadObjectsObjectMarkings struct {
	OpenDriveElement
	Marking []*RoadObjectsObjectMarkingsMarking `xml:"marking"`
}

// Specifies a marking that is either attached to one side of the object bound
// ing box or referencing outline points.
type RoadObjectsObjectMarkingsMarking struct {
	OpenDriveElement
	CornerReference []*RoadObjectsObjectMarkingsMarkingCornerReference `xml:"cornerReference"`
	Color           ERoadMarkColor                                     `xml:"color,attr"`
	LineLength      GrZero                                             `xml:"lineLength,attr"`
	Side            ESideType                                          `xml:"side,attr"`
	SpaceLength     GrEqZero                                           `xml:"spaceLength,attr"`
	StartOffset     float64                                            `xml:"startOffset,attr"`
	StopOffset      float64                                            `xml:"stopOffset,attr"`
	Weight          ERoadMarkWeight                                    `xml:"weight,attr"`
	Width           GrZero                                             `xml:"width,attr"`
	ZOffset         GrEqZero                                           `xml:"zOffset,attr"`
}

// Specifies a point by referencing an existing outline point.
type RoadObjectsObjectMarkingsMarkingCornerReference struct {
	OpenDriveElement
	Id int `xml:"id,attr"`
}

// Describes the material properties of objects, for example, patches that are
// part of the road surface but deviate from the standard road material. Super
// sedes the material specified in the <road material> element and is valid on
// ly within the outline of the parent road object.
type RoadObjectsObjectMaterial struct {
	OpenDriveElement
	Friction      GrEqZero       `xml:"friction,attr"`
	RoadMarkColor ERoadMarkColor `xml:"roadMarkColor,attr"`
	Roughness     GrEqZero       `xml:"roughness,attr"`
	Surface       string         `xml:"surface,attr"`
}

// Wrapper for the different outline entries of an object, that can contain mu
// ltiple outline definitions. If <outlines> is not used, an object can have o
// nly a single <outline> entry.
type RoadObjectsObjectOutlines struct {
	OpenDriveElement
	Outline []*RoadObjectsObjectOutlinesOutline `xml:"outline"`
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
	Closed   Bool             `xml:"closed,attr"`
	FillType EOutlineFillType `xml:"fillType,attr"`
	Id       int              `xml:"id,attr"`
	LaneType ELaneType        `xml:"laneType,attr"`
	Outer    Bool             `xml:"outer,attr"`
}

// Used to describe complex forms of objects. Defines a corner point on the ob
// ject outline relative to the object pivot point in local u/v-coordinates. T
// he insertion point and the orientation of the object are given by the @s, @
// t, @zOffset and @hdg attributes of the  element.
type RoadObjectsObjectOutlinesOutlineCornerLocal struct {
	OpenDriveElement
	Height GrEqZero `xml:"height,attr"`
	Id     int      `xml:"id,attr"`
	U      float64  `xml:"u,attr"`
	V      float64  `xml:"v,attr"`
	Z      float64  `xml:"z,attr"`
}

// Defines a corner point on the objectâ  s outline in road coordinates.
type RoadObjectsObjectOutlinesOutlineCornerRoad struct {
	OpenDriveElement
	Dz     float64  `xml:"dz,attr"`
	Height GrEqZero `xml:"height,attr"`
	Id     int      `xml:"id,attr"`
	S      GrEqZero `xml:"s,attr"`
	T      float64  `xml:"t,attr"`
}

// Details for a parking space may be added to the <object> element.
type RoadObjectsObjectParkingSpace struct {
	OpenDriveElement
	Access       ERoadObjectsObjectParkingSpaceAccess `xml:"access,attr"`
	Restrictions string                               `xml:"restrictions,attr"`
}

// To avoid lengthy XML code, objects of the same type may be repeated. Attrib
// utes of the repeated object shall overrule the attributes from the original
// object. If attributes are omitted in the repeated objects, the attributes f
// rom the original object apply.
type RoadObjectsObjectRepeat struct {
	OpenDriveElement
	DetachFromReferenceLine Bool     `xml:"detachFromReferenceLine,attr"`
	Distance                GrEqZero `xml:"distance,attr"`
	HeightEnd               GrEqZero `xml:"heightEnd,attr"`
	HeightStart             GrEqZero `xml:"heightStart,attr"`
	Length                  GrEqZero `xml:"length,attr"`
	LengthEnd               GrEqZero `xml:"lengthEnd,attr"`
	LengthStart             GrEqZero `xml:"lengthStart,attr"`
	RadiusEnd               GrEqZero `xml:"radiusEnd,attr"`
	RadiusStart             GrEqZero `xml:"radiusStart,attr"`
	S                       GrEqZero `xml:"s,attr"`
	TEnd                    float64  `xml:"tEnd,attr"`
	TStart                  float64  `xml:"tStart,attr"`
	WidthEnd                GrEqZero `xml:"widthEnd,attr"`
	WidthStart              GrEqZero `xml:"widthStart,attr"`
	ZOffsetEnd              float64  `xml:"zOffsetEnd,attr"`
	ZOffsetStart            float64  `xml:"zOffsetStart,attr"`
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
	Crg *RoadObjectsObjectSurfaceCrg `xml:"CRG"`
}

// Elevation data described in {GLO_VAR_STA_ASAM_OpenCRG} are represented by t
// he <CRG> element within the <surface> element.
type RoadObjectsObjectSurfaceCrg struct {
	OpenDriveElement
	File               string  `xml:"file,attr"`
	HideRoadSurfaceCrg Bool    `xml:"hideRoadSurfaceCRG,attr"`
	ZScale             float64 `xml:"zScale,attr"`
}

// An object reference refers to one identical object from multiple roads. The
// referenced objects require a unique ID. The <objectReference> element consi
// sts of a main element and an optional lane validity element.
type RoadObjectsObjectReference struct {
	OpenDriveElement
	Validity    []*RoadObjectsObjectLaneValidity `xml:"validity"`
	Id          string                           `xml:"id,attr"`
	Orientation EOrientation                     `xml:"orientation,attr"`
	S           GrEqZero                         `xml:"s,attr"`
	T           float64                          `xml:"t,attr"`
	ValidLength GrEqZero                         `xml:"validLength,attr"`
	ZOffset     float64                          `xml:"zOffset,attr"`
}

// Tunnels are modeled as objects in ASAM OpenDRIVE. The road with the tunnel
// object defines the part of the road that is the tunnel or the underpass. Tu
// nnels are valid for the complete cross section of a road unless a lane vali
// dity element with further restrictions is provided as child element
type RoadObjectsTunnel struct {
	OpenDriveElement
	Validity []*RoadObjectsObjectLaneValidity `xml:"validity"`
	Daylight ZeroOne                          `xml:"daylight,attr"`
	Id       string                           `xml:"id,attr"`
	Length   GrEqZero                         `xml:"length,attr"`
	Lighting ZeroOne                          `xml:"lighting,attr"`
	Name     string                           `xml:"name,attr"`
	S        GrEqZero                         `xml:"s,attr"`
	Type     ETunnelType                      `xml:"type,attr"`
}
