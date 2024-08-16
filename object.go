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
type TRoadObjects struct {
	Object TRoadObjectsObject

	ObjectReference TRoadObjectsObjectReference

	Tunnel TRoadObjectsTunnel

	Bridge TRoadObjectsBridge
}

// TODO: Doc formatting needs to be implemented!
type TRoadObjectsBridge struct {
	Validity TRoadObjectsObjectLaneValidity
}

// TODO: Doc formatting needs to be implemented!
type TRoadObjectsObject struct {
	Repeat TRoadObjectsObjectRepeat

	Outline TRoadObjectsObjectOutlinesOutline

	Outlines TRoadObjectsObjectOutlines

	Material TRoadObjectsObjectMaterial

	Validity TRoadObjectsObjectLaneValidity

	ParkingSpace TRoadObjectsObjectParkingSpace

	Markings TRoadObjectsObjectMarkings

	Borders TRoadObjectsObjectBorders

	Surface TRoadObjectsObjectSurface

	Skeleton TRoadObjectsObjectSkeleton
}

// TODO: Doc formatting needs to be implemented!
type TRoadObjectsObjectBorders struct {
	Border TRoadObjectsObjectBordersBorder
}

// TODO: Doc formatting needs to be implemented!
type TRoadObjectsObjectBordersBorder struct {
	CornerReference TRoadObjectsObjectMarkingsMarkingCornerReference
}

// TODO: Doc formatting needs to be implemented!
type TRoadObjectsObjectMarkings struct {
	Marking TRoadObjectsObjectMarkingsMarking
}

// TODO: Doc formatting needs to be implemented!
type TRoadObjectsObjectMarkingsMarking struct {
	CornerReference TRoadObjectsObjectMarkingsMarkingCornerReference
}

// TODO: Doc formatting needs to be implemented!
type TRoadObjectsObjectMarkingsMarkingCornerReference struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadObjectsObjectMaterial struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadObjectsObjectOutlines struct {
	Outline TRoadObjectsObjectOutlinesOutline
}

// TODO: Doc formatting needs to be implemented!
type TRoadObjectsObjectOutlinesOutline struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadObjectsObjectOutlinesOutlineCornerLocal struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadObjectsObjectOutlinesOutlineCornerRoad struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadObjectsObjectParkingSpace struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadObjectsObjectRepeat struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadObjectsObjectSkeleton struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadObjectsObjectSkeletonPolyline struct {
	Id int
}

// TODO: Doc formatting needs to be implemented!
type TRoadObjectsObjectSkeletonPolylineVertexLocal struct {
	Id                int
	IntersectionPoint TBool
	Radius            float64
	U                 float64
	V                 float64
	Z                 float64
}

// TODO: Doc formatting needs to be implemented!
type TRoadObjectsObjectSkeletonPolylineVertexRoad struct {
	Dz                float64
	Id                int
	IntersectionPoint TBool
	Radius            float64
	S                 TGrEqZero
	T                 float64
}

// TODO: Doc formatting needs to be implemented!
type TRoadObjectsObjectSurface struct {
	Crg TRoadObjectsObjectSurfaceCrg
}

// TODO: Doc formatting needs to be implemented!
type TRoadObjectsObjectSurfaceCrg struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadObjectsObjectReference struct {
	Validity TRoadObjectsObjectLaneValidity
}

// TODO: Doc formatting needs to be implemented!
type TRoadObjectsTunnel struct {
	Validity TRoadObjectsObjectLaneValidity
}
