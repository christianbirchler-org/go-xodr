// Code generated; DO NOT EDIT.

package xodr

type Element interface {
	Attributes() []byte
	Children() []Element
	Parent() Element
}

type Leaf interface {
	AddParent(Element)
}

type OpenDRIVE struct {
	parent   Element
	children []Element
}

func (e *OpenDRIVE) Parent() Element {
	return e.parent
}

func (e *OpenDRIVE) Children() []Element {
	return e.children
}

type Header struct {
	parent   Element
	children []Element
}

func (e *Header) Parent() Element {
	return e.parent
}

func (e *Header) Children() []Element {
	return e.children
}

type GeoReference struct {
	parent   Element
	children []Element
}

func (e *GeoReference) Parent() Element {
	return e.parent
}

func (e *GeoReference) Children() []Element {
	return e.children
}

type Offset struct {
	parent   Element
	children []Element
}

func (e *Offset) Parent() Element {
	return e.parent
}

func (e *Offset) Children() []Element {
	return e.children
}

type License struct {
	parent   Element
	children []Element
}

func (e *License) Parent() Element {
	return e.parent
}

func (e *License) Children() []Element {
	return e.children
}

type DefaultRegulations struct {
	parent   Element
	children []Element
}

func (e *DefaultRegulations) Parent() Element {
	return e.parent
}

func (e *DefaultRegulations) Children() []Element {
	return e.children
}

type RoadRegulations struct {
	parent   Element
	children []Element
}

func (e *RoadRegulations) Parent() Element {
	return e.parent
}

func (e *RoadRegulations) Children() []Element {
	return e.children
}

type SignalRegulations struct {
	parent   Element
	children []Element
}

func (e *SignalRegulations) Parent() Element {
	return e.parent
}

func (e *SignalRegulations) Children() []Element {
	return e.children
}

type Road struct {
	parent   Element
	children []Element
}

func (e *Road) Parent() Element {
	return e.parent
}

func (e *Road) Children() []Element {
	return e.children
}

type Link struct {
	parent   Element
	children []Element
}

func (e *Link) Parent() Element {
	return e.parent
}

func (e *Link) Children() []Element {
	return e.children
}

type Predecessor struct {
	parent   Element
	children []Element
}

func (e *Predecessor) Parent() Element {
	return e.parent
}

func (e *Predecessor) Children() []Element {
	return e.children
}

type Successor struct {
	parent   Element
	children []Element
}

func (e *Successor) Parent() Element {
	return e.parent
}

func (e *Successor) Children() []Element {
	return e.children
}

type Type struct {
	parent   Element
	children []Element
}

func (e *Type) Parent() Element {
	return e.parent
}

func (e *Type) Children() []Element {
	return e.children
}

type Speed struct {
	parent   Element
	children []Element
}

func (e *Speed) Parent() Element {
	return e.parent
}

func (e *Speed) Children() []Element {
	return e.children
}

type PlanView struct {
	parent   Element
	children []Element
}

func (e *PlanView) Parent() Element {
	return e.parent
}

func (e *PlanView) Children() []Element {
	return e.children
}

type Geometry struct {
	parent   Element
	children []Element
}

func (e *Geometry) Parent() Element {
	return e.parent
}

func (e *Geometry) Children() []Element {
	return e.children
}

type Line struct {
	parent   Element
	children []Element
}

func (e *Line) Parent() Element {
	return e.parent
}

func (e *Line) Children() []Element {
	return e.children
}

type Spiral struct {
	parent   Element
	children []Element
}

func (e *Spiral) Parent() Element {
	return e.parent
}

func (e *Spiral) Children() []Element {
	return e.children
}

type Arc struct {
	parent   Element
	children []Element
}

func (e *Arc) Parent() Element {
	return e.parent
}

func (e *Arc) Children() []Element {
	return e.children
}

type Poly3 struct {
	parent   Element
	children []Element
}

func (e *Poly3) Parent() Element {
	return e.parent
}

func (e *Poly3) Children() []Element {
	return e.children
}

type ParamPoly3 struct {
	parent   Element
	children []Element
}

func (e *ParamPoly3) Parent() Element {
	return e.parent
}

func (e *ParamPoly3) Children() []Element {
	return e.children
}

type ElevationProfile struct {
	parent   Element
	children []Element
}

func (e *ElevationProfile) Parent() Element {
	return e.parent
}

func (e *ElevationProfile) Children() []Element {
	return e.children
}

type Elevation struct {
	parent   Element
	children []Element
}

func (e *Elevation) Parent() Element {
	return e.parent
}

func (e *Elevation) Children() []Element {
	return e.children
}

type LateralProfile struct {
	parent   Element
	children []Element
}

func (e *LateralProfile) Parent() Element {
	return e.parent
}

func (e *LateralProfile) Children() []Element {
	return e.children
}

type Superelevation struct {
	parent   Element
	children []Element
}

func (e *Superelevation) Parent() Element {
	return e.parent
}

func (e *Superelevation) Children() []Element {
	return e.children
}

type Shape struct {
	parent   Element
	children []Element
}

func (e *Shape) Parent() Element {
	return e.parent
}

func (e *Shape) Children() []Element {
	return e.children
}

type CrossSectionSurface struct {
	parent   Element
	children []Element
}

func (e *CrossSectionSurface) Parent() Element {
	return e.parent
}

func (e *CrossSectionSurface) Children() []Element {
	return e.children
}

type TOffset struct {
	parent   Element
	children []Element
}

func (e *TOffset) Parent() Element {
	return e.parent
}

func (e *TOffset) Children() []Element {
	return e.children
}

type Coefficients struct {
	parent   Element
	children []Element
}

func (e *Coefficients) Parent() Element {
	return e.parent
}

func (e *Coefficients) Children() []Element {
	return e.children
}

type SurfaceStrips struct {
	parent   Element
	children []Element
}

func (e *SurfaceStrips) Parent() Element {
	return e.parent
}

func (e *SurfaceStrips) Children() []Element {
	return e.children
}

type Strip struct {
	parent   Element
	children []Element
}

func (e *Strip) Parent() Element {
	return e.parent
}

func (e *Strip) Children() []Element {
	return e.children
}

type Width struct {
	parent   Element
	children []Element
}

func (e *Width) Parent() Element {
	return e.parent
}

func (e *Width) Children() []Element {
	return e.children
}

type Constant struct {
	parent   Element
	children []Element
}

func (e *Constant) Parent() Element {
	return e.parent
}

func (e *Constant) Children() []Element {
	return e.children
}

type Linear struct {
	parent   Element
	children []Element
}

func (e *Linear) Parent() Element {
	return e.parent
}

func (e *Linear) Children() []Element {
	return e.children
}

type Quadratic struct {
	parent   Element
	children []Element
}

func (e *Quadratic) Parent() Element {
	return e.parent
}

func (e *Quadratic) Children() []Element {
	return e.children
}

type Cubic struct {
	parent   Element
	children []Element
}

func (e *Cubic) Parent() Element {
	return e.parent
}

func (e *Cubic) Children() []Element {
	return e.children
}

type Lanes struct {
	parent   Element
	children []Element
}

func (e *Lanes) Parent() Element {
	return e.parent
}

func (e *Lanes) Children() []Element {
	return e.children
}

type LaneOffset struct {
	parent   Element
	children []Element
}

func (e *LaneOffset) Parent() Element {
	return e.parent
}

func (e *LaneOffset) Children() []Element {
	return e.children
}

type LaneSection struct {
	parent   Element
	children []Element
}

func (e *LaneSection) Parent() Element {
	return e.parent
}

func (e *LaneSection) Children() []Element {
	return e.children
}

type Left struct {
	parent   Element
	children []Element
}

func (e *Left) Parent() Element {
	return e.parent
}

func (e *Left) Children() []Element {
	return e.children
}

type Lane struct {
	parent   Element
	children []Element
}

func (e *Lane) Parent() Element {
	return e.parent
}

func (e *Lane) Children() []Element {
	return e.children
}

type Border struct {
	parent   Element
	children []Element
}

func (e *Border) Parent() Element {
	return e.parent
}

func (e *Border) Children() []Element {
	return e.children
}

type RoadMark struct {
	parent   Element
	children []Element
}

func (e *RoadMark) Parent() Element {
	return e.parent
}

func (e *RoadMark) Children() []Element {
	return e.children
}

type Sway struct {
	parent   Element
	children []Element
}

func (e *Sway) Parent() Element {
	return e.parent
}

func (e *Sway) Children() []Element {
	return e.children
}

type Explicit struct {
	parent   Element
	children []Element
}

func (e *Explicit) Parent() Element {
	return e.parent
}

func (e *Explicit) Children() []Element {
	return e.children
}

type Material struct {
	parent   Element
	children []Element
}

func (e *Material) Parent() Element {
	return e.parent
}

func (e *Material) Children() []Element {
	return e.children
}

type Access struct {
	parent   Element
	children []Element
}

func (e *Access) Parent() Element {
	return e.parent
}

func (e *Access) Children() []Element {
	return e.children
}

type Restriction struct {
	parent   Element
	children []Element
}

func (e *Restriction) Parent() Element {
	return e.parent
}

func (e *Restriction) Children() []Element {
	return e.children
}

type Height struct {
	parent   Element
	children []Element
}

func (e *Height) Parent() Element {
	return e.parent
}

func (e *Height) Children() []Element {
	return e.children
}

type Rule struct {
	parent   Element
	children []Element
}

func (e *Rule) Parent() Element {
	return e.parent
}

func (e *Rule) Children() []Element {
	return e.children
}

type Center struct {
	parent   Element
	children []Element
}

func (e *Center) Parent() Element {
	return e.parent
}

func (e *Center) Children() []Element {
	return e.children
}

type Right struct {
	parent   Element
	children []Element
}

func (e *Right) Parent() Element {
	return e.parent
}

func (e *Right) Children() []Element {
	return e.children
}

type Objects struct {
	parent   Element
	children []Element
}

func (e *Objects) Parent() Element {
	return e.parent
}

func (e *Objects) Children() []Element {
	return e.children
}

type Object struct {
	parent   Element
	children []Element
}

func (e *Object) Parent() Element {
	return e.parent
}

func (e *Object) Children() []Element {
	return e.children
}

type Repeat struct {
	parent   Element
	children []Element
}

func (e *Repeat) Parent() Element {
	return e.parent
}

func (e *Repeat) Children() []Element {
	return e.children
}

type Outline struct {
	parent   Element
	children []Element
}

func (e *Outline) Parent() Element {
	return e.parent
}

func (e *Outline) Children() []Element {
	return e.children
}

type CornerRoad struct {
	parent   Element
	children []Element
}

func (e *CornerRoad) Parent() Element {
	return e.parent
}

func (e *CornerRoad) Children() []Element {
	return e.children
}

type CornerLocal struct {
	parent   Element
	children []Element
}

func (e *CornerLocal) Parent() Element {
	return e.parent
}

func (e *CornerLocal) Children() []Element {
	return e.children
}

type Outlines struct {
	parent   Element
	children []Element
}

func (e *Outlines) Parent() Element {
	return e.parent
}

func (e *Outlines) Children() []Element {
	return e.children
}

type Validity struct {
	parent   Element
	children []Element
}

func (e *Validity) Parent() Element {
	return e.parent
}

func (e *Validity) Children() []Element {
	return e.children
}

type ParkingSpace struct {
	parent   Element
	children []Element
}

func (e *ParkingSpace) Parent() Element {
	return e.parent
}

func (e *ParkingSpace) Children() []Element {
	return e.children
}

type Markings struct {
	parent   Element
	children []Element
}

func (e *Markings) Parent() Element {
	return e.parent
}

func (e *Markings) Children() []Element {
	return e.children
}

type Marking struct {
	parent   Element
	children []Element
}

func (e *Marking) Parent() Element {
	return e.parent
}

func (e *Marking) Children() []Element {
	return e.children
}

type CornerReference struct {
	parent   Element
	children []Element
}

func (e *CornerReference) Parent() Element {
	return e.parent
}

func (e *CornerReference) Children() []Element {
	return e.children
}

type Borders struct {
	parent   Element
	children []Element
}

func (e *Borders) Parent() Element {
	return e.parent
}

func (e *Borders) Children() []Element {
	return e.children
}

type Surface struct {
	parent   Element
	children []Element
}

func (e *Surface) Parent() Element {
	return e.parent
}

func (e *Surface) Children() []Element {
	return e.children
}

type Skeleton struct {
	parent   Element
	children []Element
}

func (e *Skeleton) Parent() Element {
	return e.parent
}

func (e *Skeleton) Children() []Element {
	return e.children
}

type Polyline struct {
	parent   Element
	children []Element
}

func (e *Polyline) Parent() Element {
	return e.parent
}

func (e *Polyline) Children() []Element {
	return e.children
}

type VertexRoad struct {
	parent   Element
	children []Element
}

func (e *VertexRoad) Parent() Element {
	return e.parent
}

func (e *VertexRoad) Children() []Element {
	return e.children
}

type VertexLocal struct {
	parent   Element
	children []Element
}

func (e *VertexLocal) Parent() Element {
	return e.parent
}

func (e *VertexLocal) Children() []Element {
	return e.children
}

type ObjectReference struct {
	parent   Element
	children []Element
}

func (e *ObjectReference) Parent() Element {
	return e.parent
}

func (e *ObjectReference) Children() []Element {
	return e.children
}

type Tunnel struct {
	parent   Element
	children []Element
}

func (e *Tunnel) Parent() Element {
	return e.parent
}

func (e *Tunnel) Children() []Element {
	return e.children
}

type Bridge struct {
	parent   Element
	children []Element
}

func (e *Bridge) Parent() Element {
	return e.parent
}

func (e *Bridge) Children() []Element {
	return e.children
}

type Signals struct {
	parent   Element
	children []Element
}

func (e *Signals) Parent() Element {
	return e.parent
}

func (e *Signals) Children() []Element {
	return e.children
}

type Signal struct {
	parent   Element
	children []Element
}

func (e *Signal) Parent() Element {
	return e.parent
}

func (e *Signal) Children() []Element {
	return e.children
}

type PositionInertial struct {
	parent   Element
	children []Element
}

func (e *PositionInertial) Parent() Element {
	return e.parent
}

func (e *PositionInertial) Children() []Element {
	return e.children
}

type PositionRoad struct {
	parent   Element
	children []Element
}

func (e *PositionRoad) Parent() Element {
	return e.parent
}

func (e *PositionRoad) Children() []Element {
	return e.children
}

type Dependency struct {
	parent   Element
	children []Element
}

func (e *Dependency) Parent() Element {
	return e.parent
}

func (e *Dependency) Children() []Element {
	return e.children
}

type Reference struct {
	parent   Element
	children []Element
}

func (e *Reference) Parent() Element {
	return e.parent
}

func (e *Reference) Children() []Element {
	return e.children
}

type SignalReference struct {
	parent   Element
	children []Element
}

func (e *SignalReference) Parent() Element {
	return e.parent
}

func (e *SignalReference) Children() []Element {
	return e.children
}

type CRG struct {
	parent   Element
	children []Element
}

func (e *CRG) Parent() Element {
	return e.parent
}

func (e *CRG) Children() []Element {
	return e.children
}

type Railroad struct {
	parent   Element
	children []Element
}

func (e *Railroad) Parent() Element {
	return e.parent
}

func (e *Railroad) Children() []Element {
	return e.children
}

type Switch struct {
	parent   Element
	children []Element
}

func (e *Switch) Parent() Element {
	return e.parent
}

func (e *Switch) Children() []Element {
	return e.children
}

type MainTrack struct {
	parent   Element
	children []Element
}

func (e *MainTrack) Parent() Element {
	return e.parent
}

func (e *MainTrack) Children() []Element {
	return e.children
}

type SideTrack struct {
	parent   Element
	children []Element
}

func (e *SideTrack) Parent() Element {
	return e.parent
}

func (e *SideTrack) Children() []Element {
	return e.children
}

type Partner struct {
	parent   Element
	children []Element
}

func (e *Partner) Parent() Element {
	return e.parent
}

func (e *Partner) Children() []Element {
	return e.children
}

type Controller struct {
	parent   Element
	children []Element
}

func (e *Controller) Parent() Element {
	return e.parent
}

func (e *Controller) Children() []Element {
	return e.children
}

type Control struct {
	parent   Element
	children []Element
}

func (e *Control) Parent() Element {
	return e.parent
}

func (e *Control) Children() []Element {
	return e.children
}

type JunctiontypeCrossing struct {
	parent   Element
	children []Element
}

func (e *JunctiontypeCrossing) Parent() Element {
	return e.parent
}

func (e *JunctiontypeCrossing) Children() []Element {
	return e.children
}

type RoadSection struct {
	parent   Element
	children []Element
}

func (e *RoadSection) Parent() Element {
	return e.parent
}

func (e *RoadSection) Children() []Element {
	return e.children
}

type Priority struct {
	parent   Element
	children []Element
}

func (e *Priority) Parent() Element {
	return e.parent
}

func (e *Priority) Children() []Element {
	return e.children
}

type JunctiontypeDefault struct {
	parent   Element
	children []Element
}

func (e *JunctiontypeDefault) Parent() Element {
	return e.parent
}

func (e *JunctiontypeDefault) Children() []Element {
	return e.children
}

type Connection struct {
	parent   Element
	children []Element
}

func (e *Connection) Parent() Element {
	return e.parent
}

func (e *Connection) Children() []Element {
	return e.children
}

type LaneLink struct {
	parent   Element
	children []Element
}

func (e *LaneLink) Parent() Element {
	return e.parent
}

func (e *LaneLink) Children() []Element {
	return e.children
}

type CrossPath struct {
	parent   Element
	children []Element
}

func (e *CrossPath) Parent() Element {
	return e.parent
}

func (e *CrossPath) Children() []Element {
	return e.children
}

type StartLaneLink struct {
	parent   Element
	children []Element
}

func (e *StartLaneLink) Parent() Element {
	return e.parent
}

func (e *StartLaneLink) Children() []Element {
	return e.children
}

type EndLaneLink struct {
	parent   Element
	children []Element
}

func (e *EndLaneLink) Parent() Element {
	return e.parent
}

func (e *EndLaneLink) Children() []Element {
	return e.children
}

type Boundary struct {
	parent   Element
	children []Element
}

func (e *Boundary) Parent() Element {
	return e.parent
}

func (e *Boundary) Children() []Element {
	return e.children
}

type SegmenttypeJoint struct {
	parent   Element
	children []Element
}

func (e *SegmenttypeJoint) Parent() Element {
	return e.parent
}

func (e *SegmenttypeJoint) Children() []Element {
	return e.children
}

type SegmenttypeLane struct {
	parent   Element
	children []Element
}

func (e *SegmenttypeLane) Parent() Element {
	return e.parent
}

func (e *SegmenttypeLane) Children() []Element {
	return e.children
}

type ElevationGrid struct {
	parent   Element
	children []Element
}

func (e *ElevationGrid) Parent() Element {
	return e.parent
}

func (e *ElevationGrid) Children() []Element {
	return e.children
}

type JunctiontypeDirect struct {
	parent   Element
	children []Element
}

func (e *JunctiontypeDirect) Parent() Element {
	return e.parent
}

func (e *JunctiontypeDirect) Children() []Element {
	return e.children
}

type JunctiontypeVirtual struct {
	parent   Element
	children []Element
}

func (e *JunctiontypeVirtual) Parent() Element {
	return e.parent
}

func (e *JunctiontypeVirtual) Children() []Element {
	return e.children
}

type ConnectiontypeDefault struct {
	parent   Element
	children []Element
}

func (e *ConnectiontypeDefault) Parent() Element {
	return e.parent
}

func (e *ConnectiontypeDefault) Children() []Element {
	return e.children
}

type ConnectiontypeVirtual struct {
	parent   Element
	children []Element
}

func (e *ConnectiontypeVirtual) Parent() Element {
	return e.parent
}

func (e *ConnectiontypeVirtual) Children() []Element {
	return e.children
}

type JunctionGroup struct {
	parent   Element
	children []Element
}

func (e *JunctionGroup) Parent() Element {
	return e.parent
}

func (e *JunctionGroup) Children() []Element {
	return e.children
}

type JunctionReference struct {
	parent   Element
	children []Element
}

func (e *JunctionReference) Parent() Element {
	return e.parent
}

func (e *JunctionReference) Children() []Element {
	return e.children
}

type Station struct {
	parent   Element
	children []Element
}

func (e *Station) Parent() Element {
	return e.parent
}

func (e *Station) Children() []Element {
	return e.children
}

type Platform struct {
	parent   Element
	children []Element
}

func (e *Platform) Parent() Element {
	return e.parent
}

func (e *Platform) Children() []Element {
	return e.children
}

type Segment struct {
	parent   Element
	children []Element
}

func (e *Segment) Parent() Element {
	return e.parent
}

func (e *Segment) Children() []Element {
	return e.children
}
