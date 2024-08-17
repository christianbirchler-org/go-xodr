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
	Object          RoadObjectsObject
	ObjectReference RoadObjectsObjectReference
	Tunnel          RoadObjectsTunnel
	Bridge          RoadObjectsBridge
}

// TODO: Doc formatting needs to be implemented!
type RoadObjectsBridge struct {
	OpenDriveElement
	Validity RoadObjectsObjectLaneValidity
	Id       string
	Length   TGrEqZero
	Name     string
	S        TGrEqZero
	Type     EBridgeType
}

// TODO: Doc formatting needs to be implemented!
type RoadObjectsObject struct {
	OpenDriveElement
	Repeat       RoadObjectsObjectRepeat
	Outline      RoadObjectsObjectOutlinesOutline
	Outlines     RoadObjectsObjectOutlines
	Material     RoadObjectsObjectMaterial
	Validity     RoadObjectsObjectLaneValidity
	ParkingSpace RoadObjectsObjectParkingSpace
	Markings     RoadObjectsObjectMarkings
	Borders      RoadObjectsObjectBorders
	Surface      RoadObjectsObjectSurface
	Skeleton     RoadObjectsObjectSkeleton
	Dynamic      TYesNo
	Hdg          float64
	Height       TGrEqZero
	Id           string
	Length       TGrZero
	Name         string
	Orientation  EOrientation
	PerpToRoad   TBool
	Pitch        float64
	Radius       TGrZero
	Roll         float64
	S            TGrEqZero
	Subtype      string
	T            float64
	Type         EObjectType
	ValidLength  TGrEqZero
	Width        float64
	ZOffset      float64
}

// TODO: Doc formatting needs to be implemented!
type RoadObjectsObjectBorders struct {
	OpenDriveElement
	Border RoadObjectsObjectBordersBorder
}

// TODO: Doc formatting needs to be implemented!
type RoadObjectsObjectBordersBorder struct {
	OpenDriveElement
	CornerReference    RoadObjectsObjectMarkingsMarkingCornerReference
	OutlineId          int
	Type               EBorderType
	UseCompleteOutline TBool
	Width              TGrEqZero
}

// TODO: Doc formatting needs to be implemented!
type RoadObjectsObjectMarkings struct {
	OpenDriveElement
	Marking RoadObjectsObjectMarkingsMarking
}

// TODO: Doc formatting needs to be implemented!
type RoadObjectsObjectMarkingsMarking struct {
	OpenDriveElement
	CornerReference RoadObjectsObjectMarkingsMarkingCornerReference
	Color           ERoadMarkColor
	LineLength      TGrZero
	Side            ESideType
	SpaceLength     TGrEqZero
	StartOffset     float64
	StopOffset      float64
	Weight          ERoadMarkWeight
	Width           TGrZero
	ZOffset         TGrEqZero
}

// TODO: Doc formatting needs to be implemented!
type RoadObjectsObjectMarkingsMarkingCornerReference struct {
	OpenDriveElement
	Id int
}

// TODO: Doc formatting needs to be implemented!
type RoadObjectsObjectMaterial struct {
	OpenDriveElement
	Friction      TGrEqZero
	RoadMarkColor ERoadMarkColor
	Roughness     TGrEqZero
	Surface       string
}

// TODO: Doc formatting needs to be implemented!
type RoadObjectsObjectOutlines struct {
	OpenDriveElement
	Outline RoadObjectsObjectOutlinesOutline
}

// TODO: Doc formatting needs to be implemented!
type RoadObjectsObjectOutlinesOutline struct {
	OpenDriveElement
	Closed   TBool
	FillType EOutlineFillType
	Id       int
	LaneType ELaneType
	Outer    TBool
}

// TODO: Doc formatting needs to be implemented!
type RoadObjectsObjectOutlinesOutlineCornerLocal struct {
	OpenDriveElement
	Height TGrEqZero
	Id     int
	U      float64
	V      float64
	Z      float64
}

// TODO: Doc formatting needs to be implemented!
type RoadObjectsObjectOutlinesOutlineCornerRoad struct {
	OpenDriveElement
	Dz     float64
	Height TGrEqZero
	Id     int
	S      TGrEqZero
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
	DetachFromReferenceLine TBool
	Distance                TGrEqZero
	HeightEnd               TGrEqZero
	HeightStart             TGrEqZero
	Length                  TGrEqZero
	LengthEnd               TGrEqZero
	LengthStart             TGrEqZero
	RadiusEnd               TGrEqZero
	RadiusStart             TGrEqZero
	S                       TGrEqZero
	TEnd                    float64
	TStart                  float64
	WidthEnd                TGrEqZero
	WidthStart              TGrEqZero
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
	IntersectionPoint bool
	Radius            float64
	U                 float64
	V                 float64
	Z                 float64
}

// TODO: Doc formatting needs to be implemented!
type RoadObjectsObjectSkeletonPolylineVertexRoad struct {
	Dz                float64
	Id                int
	IntersectionPoint bool
	Radius            float64
	S                 float64
	T                 float64
}

// TODO: Doc formatting needs to be implemented!
type RoadObjectsObjectSurface struct {
	OpenDriveElement
	Crg RoadObjectsObjectSurfaceCrg
}

// TODO: Doc formatting needs to be implemented!
type RoadObjectsObjectSurfaceCrg struct {
	OpenDriveElement
	File               string
	HideRoadSurfaceCrg TBool
	ZScale             float64
}

// TODO: Doc formatting needs to be implemented!
type RoadObjectsObjectReference struct {
	OpenDriveElement
	Validity    RoadObjectsObjectLaneValidity
	Id          string
	Orientation EOrientation
	S           TGrEqZero
	T           float64
	ValidLength TGrEqZero
	ZOffset     float64
}

// TODO: Doc formatting needs to be implemented!
type RoadObjectsTunnel struct {
	OpenDriveElement
	Validity RoadObjectsObjectLaneValidity
	Daylight TZeroOne
	Id       string
	Length   TGrEqZero
	Lighting TZeroOne
	Name     string
	S        TGrEqZero
	Type     ETunnelType
}
