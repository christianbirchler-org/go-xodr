// Code generated; DO NOT EDIT.

package xodr

type EBorderType struct {
}

type EBridgeType struct {
}

type EObjectType struct {
}

type EOrientation struct {
}

type EOutlineFillType struct {
}

type ERoadObjectsObjectParkingSpaceAccess struct {
}

type ESideType struct {
}

type ETunnelType struct {
}

// TODO: Doc formatting needs to be implemented!
type RoadObjects struct {
	OpenDriveElement
	Object          []*RoadObjectsObject
	ObjectReference []*RoadObjectsObjectReference
	Tunnel          []*RoadObjectsTunnel
	Bridge          []*RoadObjectsBridge
}

// TODO: Doc formatting needs to be implemented!
type RoadObjectsBridge struct {
	OpenDriveElement
	Validity []*RoadObjectsObjectLaneValidity
	Id       string
	Length   GrEqZero
	Name     string
	S        GrEqZero
	Type     EBridgeType
}

// TODO: Doc formatting needs to be implemented!
type RoadObjectsObject struct {
	OpenDriveElement
	Repeat       []*RoadObjectsObjectRepeat
	Outline      []*RoadObjectsObjectOutlinesOutline
	Outlines     []*RoadObjectsObjectOutlines
	Material     []*RoadObjectsObjectMaterial
	Validity     []*RoadObjectsObjectLaneValidity
	ParkingSpace []*RoadObjectsObjectParkingSpace
	Markings     []*RoadObjectsObjectMarkings
	Borders      []*RoadObjectsObjectBorders
	Surface      []*RoadObjectsObjectSurface
	Skeleton     []*RoadObjectsObjectSkeleton
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

// TODO: Doc formatting needs to be implemented!
type RoadObjectsObjectBorders struct {
	OpenDriveElement
	Border []*RoadObjectsObjectBordersBorder
}

// TODO: Doc formatting needs to be implemented!
type RoadObjectsObjectBordersBorder struct {
	OpenDriveElement
	CornerReference    []*RoadObjectsObjectMarkingsMarkingCornerReference
	OutlineId          int
	Type               EBorderType
	UseCompleteOutline Bool
	Width              GrEqZero
}

// TODO: Doc formatting needs to be implemented!
type RoadObjectsObjectMarkings struct {
	OpenDriveElement
	Marking []*RoadObjectsObjectMarkingsMarking
}

// TODO: Doc formatting needs to be implemented!
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

// TODO: Doc formatting needs to be implemented!
type RoadObjectsObjectMarkingsMarkingCornerReference struct {
	OpenDriveElement
	Id int
}

// TODO: Doc formatting needs to be implemented!
type RoadObjectsObjectMaterial struct {
	OpenDriveElement
	Friction      GrEqZero
	RoadMarkColor ERoadMarkColor
	Roughness     GrEqZero
	Surface       string
}

// TODO: Doc formatting needs to be implemented!
type RoadObjectsObjectOutlines struct {
	OpenDriveElement
	Outline []*RoadObjectsObjectOutlinesOutline
}

// TODO: Doc formatting needs to be implemented!
type RoadObjectsObjectOutlinesOutline struct {
	OpenDriveElement
	Closed   Bool
	FillType EOutlineFillType
	Id       int
	LaneType ELaneType
	Outer    Bool
}

// TODO: Doc formatting needs to be implemented!
type RoadObjectsObjectOutlinesOutlineCornerLocal struct {
	OpenDriveElement
	Height GrEqZero
	Id     int
	U      float64
	V      float64
	Z      float64
}

// TODO: Doc formatting needs to be implemented!
type RoadObjectsObjectOutlinesOutlineCornerRoad struct {
	OpenDriveElement
	Dz     float64
	Height GrEqZero
	Id     int
	S      GrEqZero
	T      float64
}

// TODO: Doc formatting needs to be implemented!
type RoadObjectsObjectParkingSpace struct {
	OpenDriveElement
	Access       ERoadObjectsObjectParkingSpaceAccess
	Restrictions string
}

// TODO: Doc formatting needs to be implemented!
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

// TODO: Doc formatting needs to be implemented!
type RoadObjectsObjectSkeleton struct {
}

// TODO: Doc formatting needs to be implemented!
type RoadObjectsObjectSkeletonPolyline struct {
	Id int
}

// TODO: Doc formatting needs to be implemented!
type RoadObjectsObjectSkeletonPolylineVertexLocal struct {
	Id                int
	IntersectionPoint Bool
	Radius            float64
	U                 float64
	V                 float64
	Z                 float64
}

// TODO: Doc formatting needs to be implemented!
type RoadObjectsObjectSkeletonPolylineVertexRoad struct {
	Dz                float64
	Id                int
	IntersectionPoint Bool
	Radius            float64
	S                 GrEqZero
	T                 float64
}

// TODO: Doc formatting needs to be implemented!
type RoadObjectsObjectSurface struct {
	OpenDriveElement
	Crg []*RoadObjectsObjectSurfaceCrg
}

// TODO: Doc formatting needs to be implemented!
type RoadObjectsObjectSurfaceCrg struct {
	OpenDriveElement
	File               string
	HideRoadSurfaceCrg Bool
	ZScale             float64
}

// TODO: Doc formatting needs to be implemented!
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

// TODO: Doc formatting needs to be implemented!
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
