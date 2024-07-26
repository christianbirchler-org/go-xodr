// Code generated; DO NOT EDIT.

package xodr

type Element interface {
	//Attributes() []byte
	Children() []Element
	Parent() Element
}

type Leaf interface {
	AddParent(Element)
}

type OpenDRIVE struct {
	children []Element
}

func (e *OpenDRIVE) Children() []Element {
	return e.children
}

type Header struct {
	children []Element
}

func (e *Header) Children() []Element {
	return e.children
}

type GeoReference struct {
	children []Element
}

func (e *GeoReference) Children() []Element {
	return e.children
}

type Offset struct {
	children []Element
}

func (e *Offset) Children() []Element {
	return e.children
}

type License struct {
	children []Element
}

func (e *License) Children() []Element {
	return e.children
}

type DefaultRegulations struct {
	children []Element
}

func (e *DefaultRegulations) Children() []Element {
	return e.children
}

type RoadRegulations struct {
	children []Element
}

func (e *RoadRegulations) Children() []Element {
	return e.children
}

type SignalRegulations struct {
	children []Element
}

func (e *SignalRegulations) Children() []Element {
	return e.children
}

type Road struct {
	children []Element
}

func (e *Road) Children() []Element {
	return e.children
}

type Link struct {
	children []Element
}

func (e *Link) Children() []Element {
	return e.children
}

type Predecessor struct {
	children []Element
}

func (e *Predecessor) Children() []Element {
	return e.children
}

type Successor struct {
	children []Element
}

func (e *Successor) Children() []Element {
	return e.children
}

type Type struct {
	children []Element
}

func (e *Type) Children() []Element {
	return e.children
}

type Speed struct {
	children []Element
}

func (e *Speed) Children() []Element {
	return e.children
}

type PlanView struct {
	children []Element
}

func (e *PlanView) Children() []Element {
	return e.children
}

type Geometry struct {
	children []Element
}

func (e *Geometry) Children() []Element {
	return e.children
}

type Line struct {
	children []Element
}

func (e *Line) Children() []Element {
	return e.children
}

type Spiral struct {
	children []Element
}

func (e *Spiral) Children() []Element {
	return e.children
}

type Arc struct {
	children []Element
}

func (e *Arc) Children() []Element {
	return e.children
}

type Poly3 struct {
	children []Element
}

func (e *Poly3) Children() []Element {
	return e.children
}

type ParamPoly3 struct {
	children []Element
}

func (e *ParamPoly3) Children() []Element {
	return e.children
}

type ElevationProfile struct {
	children []Element
}

func (e *ElevationProfile) Children() []Element {
	return e.children
}

type Elevation struct {
	children []Element
}

func (e *Elevation) Children() []Element {
	return e.children
}

type LateralProfile struct {
	children []Element
}

func (e *LateralProfile) Children() []Element {
	return e.children
}

type Superelevation struct {
	children []Element
}

func (e *Superelevation) Children() []Element {
	return e.children
}

type Shape struct {
	children []Element
}

func (e *Shape) Children() []Element {
	return e.children
}

type CrossSectionSurface struct {
	children []Element
}

func (e *CrossSectionSurface) Children() []Element {
	return e.children
}

type TOffset struct {
	children []Element
}

func (e *TOffset) Children() []Element {
	return e.children
}

type Coefficients struct {
	children []Element
}

func (e *Coefficients) Children() []Element {
	return e.children
}

type SurfaceStrips struct {
	children []Element
}

func (e *SurfaceStrips) Children() []Element {
	return e.children
}

type Strip struct {
	children []Element
}

func (e *Strip) Children() []Element {
	return e.children
}

type Width struct {
	children []Element
}

func (e *Width) Children() []Element {
	return e.children
}

type Constant struct {
	children []Element
}

func (e *Constant) Children() []Element {
	return e.children
}

type Linear struct {
	children []Element
}

func (e *Linear) Children() []Element {
	return e.children
}

type Quadratic struct {
	children []Element
}

func (e *Quadratic) Children() []Element {
	return e.children
}

type Cubic struct {
	children []Element
}

func (e *Cubic) Children() []Element {
	return e.children
}

type Lanes struct {
	children []Element
}

func (e *Lanes) Children() []Element {
	return e.children
}

type LaneOffset struct {
	children []Element
}

func (e *LaneOffset) Children() []Element {
	return e.children
}

type LaneSection struct {
	children []Element
}

func (e *LaneSection) Children() []Element {
	return e.children
}

type Left struct {
	children []Element
}

func (e *Left) Children() []Element {
	return e.children
}

type Lane struct {
	children []Element
}

func (e *Lane) Children() []Element {
	return e.children
}

type Border struct {
	children []Element
}

func (e *Border) Children() []Element {
	return e.children
}

type RoadMark struct {
	children []Element
}

func (e *RoadMark) Children() []Element {
	return e.children
}

type Sway struct {
	children []Element
}

func (e *Sway) Children() []Element {
	return e.children
}

type Explicit struct {
	children []Element
}

func (e *Explicit) Children() []Element {
	return e.children
}

type Material struct {
	children []Element
}

func (e *Material) Children() []Element {
	return e.children
}

type Access struct {
	children []Element
}

func (e *Access) Children() []Element {
	return e.children
}

type Restriction struct {
	children []Element
}

func (e *Restriction) Children() []Element {
	return e.children
}

type Height struct {
	children []Element
}

func (e *Height) Children() []Element {
	return e.children
}

type Rule struct {
	children []Element
}

func (e *Rule) Children() []Element {
	return e.children
}

type Center struct {
	children []Element
}

func (e *Center) Children() []Element {
	return e.children
}

type Right struct {
	children []Element
}

func (e *Right) Children() []Element {
	return e.children
}

type Objects struct {
	children []Element
}

func (e *Objects) Children() []Element {
	return e.children
}

type Object struct {
	children []Element
}

func (e *Object) Children() []Element {
	return e.children
}

type Repeat struct {
	children []Element
}

func (e *Repeat) Children() []Element {
	return e.children
}

type Outline struct {
	children []Element
}

func (e *Outline) Children() []Element {
	return e.children
}

type CornerRoad struct {
	children []Element
}

func (e *CornerRoad) Children() []Element {
	return e.children
}

type CornerLocal struct {
	children []Element
}

func (e *CornerLocal) Children() []Element {
	return e.children
}

type Outlines struct {
	children []Element
}

func (e *Outlines) Children() []Element {
	return e.children
}

type Validity struct {
	children []Element
}

func (e *Validity) Children() []Element {
	return e.children
}

type ParkingSpace struct {
	children []Element
}

func (e *ParkingSpace) Children() []Element {
	return e.children
}

type Markings struct {
	children []Element
}

func (e *Markings) Children() []Element {
	return e.children
}

type Marking struct {
	children []Element
}

func (e *Marking) Children() []Element {
	return e.children
}

type CornerReference struct {
	children []Element
}

func (e *CornerReference) Children() []Element {
	return e.children
}

type Borders struct {
	children []Element
}

func (e *Borders) Children() []Element {
	return e.children
}

type Surface struct {
	children []Element
}

func (e *Surface) Children() []Element {
	return e.children
}

type Skeleton struct {
	children []Element
}

func (e *Skeleton) Children() []Element {
	return e.children
}

type Polyline struct {
	children []Element
}

func (e *Polyline) Children() []Element {
	return e.children
}

type VertexRoad struct {
	children []Element
}

func (e *VertexRoad) Children() []Element {
	return e.children
}

type VertexLocal struct {
	children []Element
}

func (e *VertexLocal) Children() []Element {
	return e.children
}

type ObjectReference struct {
	children []Element
}

func (e *ObjectReference) Children() []Element {
	return e.children
}

type Tunnel struct {
	children []Element
}

func (e *Tunnel) Children() []Element {
	return e.children
}

type Bridge struct {
	children []Element
}

func (e *Bridge) Children() []Element {
	return e.children
}

type Signals struct {
	children []Element
}

func (e *Signals) Children() []Element {
	return e.children
}

type Signal struct {
	children []Element
}

func (e *Signal) Children() []Element {
	return e.children
}

type PositionInertial struct {
	children []Element
}

func (e *PositionInertial) Children() []Element {
	return e.children
}

type PositionRoad struct {
	children []Element
}

func (e *PositionRoad) Children() []Element {
	return e.children
}

type Dependency struct {
	children []Element
}

func (e *Dependency) Children() []Element {
	return e.children
}

type Reference struct {
	children []Element
}

func (e *Reference) Children() []Element {
	return e.children
}

type SignalReference struct {
	children []Element
}

func (e *SignalReference) Children() []Element {
	return e.children
}

type CRG struct {
	children []Element
}

func (e *CRG) Children() []Element {
	return e.children
}

type Railroad struct {
	children []Element
}

func (e *Railroad) Children() []Element {
	return e.children
}

type Switch struct {
	children []Element
}

func (e *Switch) Children() []Element {
	return e.children
}

type MainTrack struct {
	children []Element
}

func (e *MainTrack) Children() []Element {
	return e.children
}

type SideTrack struct {
	children []Element
}

func (e *SideTrack) Children() []Element {
	return e.children
}

type Partner struct {
	children []Element
}

func (e *Partner) Children() []Element {
	return e.children
}

type Controller struct {
	children []Element
}

func (e *Controller) Children() []Element {
	return e.children
}

type Control struct {
	children []Element
}

func (e *Control) Children() []Element {
	return e.children
}

type JunctiontypeCrossing struct {
	children []Element
}

func (e *JunctiontypeCrossing) Children() []Element {
	return e.children
}

type RoadSection struct {
	children []Element
}

func (e *RoadSection) Children() []Element {
	return e.children
}

type Priority struct {
	children []Element
}

func (e *Priority) Children() []Element {
	return e.children
}

type JunctiontypeDefault struct {
	children []Element
}

func (e *JunctiontypeDefault) Children() []Element {
	return e.children
}

type Connection struct {
	children []Element
}

func (e *Connection) Children() []Element {
	return e.children
}

type LaneLink struct {
	children []Element
}

func (e *LaneLink) Children() []Element {
	return e.children
}

type CrossPath struct {
	children []Element
}

func (e *CrossPath) Children() []Element {
	return e.children
}

type StartLaneLink struct {
	children []Element
}

func (e *StartLaneLink) Children() []Element {
	return e.children
}

type EndLaneLink struct {
	children []Element
}

func (e *EndLaneLink) Children() []Element {
	return e.children
}

type Boundary struct {
	children []Element
}

func (e *Boundary) Children() []Element {
	return e.children
}

type SegmenttypeJoint struct {
	children []Element
}

func (e *SegmenttypeJoint) Children() []Element {
	return e.children
}

type SegmenttypeLane struct {
	children []Element
}

func (e *SegmenttypeLane) Children() []Element {
	return e.children
}

type ElevationGrid struct {
	children []Element
}

func (e *ElevationGrid) Children() []Element {
	return e.children
}

type JunctiontypeDirect struct {
	children []Element
}

func (e *JunctiontypeDirect) Children() []Element {
	return e.children
}

type JunctiontypeVirtual struct {
	children []Element
}

func (e *JunctiontypeVirtual) Children() []Element {
	return e.children
}

type ConnectiontypeDefault struct {
	children []Element
}

func (e *ConnectiontypeDefault) Children() []Element {
	return e.children
}

type ConnectiontypeVirtual struct {
	children []Element
}

func (e *ConnectiontypeVirtual) Children() []Element {
	return e.children
}

type JunctionGroup struct {
	children []Element
}

func (e *JunctionGroup) Children() []Element {
	return e.children
}

type JunctionReference struct {
	children []Element
}

func (e *JunctionReference) Children() []Element {
	return e.children
}

type Station struct {
	children []Element
}

func (e *Station) Children() []Element {
	return e.children
}

type Platform struct {
	children []Element
}

func (e *Platform) Children() []Element {
	return e.children
}

type Segment struct {
	children []Element
}

func (e *Segment) Children() []Element {
	return e.children
}
