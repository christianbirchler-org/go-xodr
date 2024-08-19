// Code generated; DO NOT EDIT.

package xodr

type EAccessRestrictionType string

type ELaneDirection string

type ELaneAdvisory string

type ELaneType string

type ERoadLanesLaneSectionLcrLaneRoadMarkLaneChange string

type ERoadLanesLaneSectionLrLaneAccessRule string

type ERoadMarkColor string

type ERoadMarkRule string

type ERoadMarkType string

type ERoadMarkWeight string

type Bool string

// TODO: Doc formatting needs to be implemented!
type RoadLanes struct {
	OpenDriveElement
	LaneOffset  []*RoadLanesLaneOffset
	LaneSection []*RoadLanesLaneSection
}

// TODO: Doc formatting needs to be implemented!
type RoadLanesLaneOffset struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type RoadLanesLaneSection struct {
	OpenDriveElement
	Left   []*RoadLanesLaneSectionLeft
	Center []*RoadLanesLaneSectionCenter
	Right  []*RoadLanesLaneSectionRight
}

// TODO: Doc formatting needs to be implemented!
type RoadLanesLaneSectionCenter struct {
	OpenDriveElement
	Lane []*RoadLanesLaneSectionCenterLane
}

// TODO: Doc formatting needs to be implemented!
type RoadLanesLaneSectionCenterLane struct {
	Id       int
	Level    Bool
	Type     ELaneType
	Link     RoadLanesLaneSectionLcrLaneLink
	RoadMark RoadLanesLaneSectionLcrLaneRoadMark
}

// TODO: Doc formatting needs to be implemented!
type RoadLanesLaneSectionLcrLaneLink struct {
	OpenDriveElement
	Predecessor []*RoadLanesLaneSectionLcrLaneLinkPredecessorSuccessor
	Successor   []*RoadLanesLaneSectionLcrLaneLinkPredecessorSuccessor
}

// TODO: Doc formatting needs to be implemented!
type RoadLanesLaneSectionLcrLaneLinkPredecessorSuccessor struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type RoadLanesLaneSectionLcrLaneRoadMark struct {
	OpenDriveElement
	Sway     []*RoadLanesLaneSectionLcrLaneRoadMarkSway
	Type     []*RoadLanesLaneSectionLcrLaneRoadMarkType
	Explicit []*RoadLanesLaneSectionLcrLaneRoadMarkExplicit
}

// TODO: Doc formatting needs to be implemented!
type RoadLanesLaneSectionLcrLaneRoadMarkExplicit struct {
	OpenDriveElement
	Line []*RoadLanesLaneSectionLcrLaneRoadMarkExplicitLine
}

// TODO: Doc formatting needs to be implemented!
type RoadLanesLaneSectionLcrLaneRoadMarkExplicitLine struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type RoadLanesLaneSectionLcrLaneRoadMarkSway struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type RoadLanesLaneSectionLcrLaneRoadMarkType struct {
	OpenDriveElement
	Line []*RoadLanesLaneSectionLcrLaneRoadMarkTypeLine
}

// TODO: Doc formatting needs to be implemented!
type RoadLanesLaneSectionLcrLaneRoadMarkTypeLine struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type RoadLanesLaneSectionLeft struct {
	OpenDriveElement
	Lane []*RoadLanesLaneSectionLeftLane
}

// TODO: Doc formatting needs to be implemented!
type RoadLanesLaneSectionLeftLane struct {
	RoadLanesLaneSectionLrLane
}

// TODO: Doc formatting needs to be implemented!
type RoadLanesLaneSectionLrLane struct {
	OpenDriveElement
	Link     []*RoadLanesLaneSectionLcrLaneLink
	RoadMark []*RoadLanesLaneSectionLcrLaneRoadMark
	Material []*RoadLanesLaneSectionLrLaneMaterial
	Speed    []*RoadLanesLaneSectionLrLaneSpeed
	Access   []*RoadLanesLaneSectionLrLaneAccess
	Height   []*RoadLanesLaneSectionLrLaneHeight
	Rule     []*RoadLanesLaneSectionLrLaneRule
}

// TODO: Doc formatting needs to be implemented!
type RoadLanesLaneSectionLrLaneAccess struct {
	OpenDriveElement
	Restriction []*RoadLanesLaneSectionLrLaneAccessRestriction
}

// TODO: Doc formatting needs to be implemented!
type RoadLanesLaneSectionLrLaneAccessRestriction struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type RoadLanesLaneSectionLrLaneBorder struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type RoadLanesLaneSectionLrLaneHeight struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type RoadLanesLaneSectionLrLaneMaterial struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type RoadLanesLaneSectionLrLaneRule struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type RoadLanesLaneSectionLrLaneSpeed struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type RoadLanesLaneSectionLrLaneWidth struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type RoadLanesLaneSectionRight struct {
	OpenDriveElement
	Lane []*RoadLanesLaneSectionRightLane
}

// TODO: Doc formatting needs to be implemented!
type RoadLanesLaneSectionRightLane struct {
	RoadLanesLaneSectionLrLane
}

// TODO: Doc formatting needs to be implemented!
type RoadObjectsObjectLaneValidity struct {
	OpenDriveElement
}
