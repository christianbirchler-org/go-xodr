// Code generated; DO NOT EDIT.

package xodr

type EAccessRestrictionType struct {
}

type ELaneDirection struct {
}

type ELaneAdvisory struct {
}

type ELaneType struct {
}

type ERoadLanesLaneSectionLcrLaneRoadMarkLaneChange struct {
}

type ERoadLanesLaneSectionLrLaneAccessRule struct {
}

type ERoadMarkColor struct {
}

type ERoadMarkRule struct {
}

type ERoadMarkType struct {
}

type ERoadMarkWeight struct {
}

type TBool struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadLanes struct {
	LaneOffset TRoadLanesLaneOffset

	LaneSection TRoadLanesLaneSection
}

// TODO: Doc formatting needs to be implemented!
type TRoadLanesLaneOffset struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadLanesLaneSection struct {
	Left TRoadLanesLaneSectionLeft

	Center TRoadLanesLaneSectionCenter

	Right TRoadLanesLaneSectionRight
}

// TODO: Doc formatting needs to be implemented!
type TRoadLanesLaneSectionCenter struct {
	Lane TRoadLanesLaneSectionCenterLane
}

// TODO: Doc formatting needs to be implemented!
type TRoadLanesLaneSectionCenterLane struct {
	Id    int
	Level TBool
	Type  ELaneType
}

// TODO: Doc formatting needs to be implemented!
type TRoadLanesLaneSectionLcrLaneLink struct {
	Predecessor TRoadLanesLaneSectionLcrLaneLinkPredecessorSuccessor

	Successor TRoadLanesLaneSectionLcrLaneLinkPredecessorSuccessor
}

// TODO: Doc formatting needs to be implemented!
type TRoadLanesLaneSectionLcrLaneLinkPredecessorSuccessor struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadLanesLaneSectionLcrLaneRoadMark struct {
	Sway TRoadLanesLaneSectionLcrLaneRoadMarkSway

	Type TRoadLanesLaneSectionLcrLaneRoadMarkType

	Explicit TRoadLanesLaneSectionLcrLaneRoadMarkExplicit
}

// TODO: Doc formatting needs to be implemented!
type TRoadLanesLaneSectionLcrLaneRoadMarkExplicit struct {
	Line TRoadLanesLaneSectionLcrLaneRoadMarkExplicitLine
}

// TODO: Doc formatting needs to be implemented!
type TRoadLanesLaneSectionLcrLaneRoadMarkExplicitLine struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadLanesLaneSectionLcrLaneRoadMarkSway struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadLanesLaneSectionLcrLaneRoadMarkType struct {
	Line TRoadLanesLaneSectionLcrLaneRoadMarkTypeLine
}

// TODO: Doc formatting needs to be implemented!
type TRoadLanesLaneSectionLcrLaneRoadMarkTypeLine struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadLanesLaneSectionLeft struct {
	Lane TRoadLanesLaneSectionLeftLane
}

// TODO: Doc formatting needs to be implemented!
type TRoadLanesLaneSectionLeftLane struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadLanesLaneSectionLrLane struct {
	Link TRoadLanesLaneSectionLcrLaneLink

	RoadMark TRoadLanesLaneSectionLcrLaneRoadMark

	Material TRoadLanesLaneSectionLrLaneMaterial

	Speed TRoadLanesLaneSectionLrLaneSpeed

	Access TRoadLanesLaneSectionLrLaneAccess

	Height TRoadLanesLaneSectionLrLaneHeight

	Rule TRoadLanesLaneSectionLrLaneRule
}

// TODO: Doc formatting needs to be implemented!
type TRoadLanesLaneSectionLrLaneAccess struct {
	Restriction TRoadLanesLaneSectionLrLaneAccessRestriction
}

// TODO: Doc formatting needs to be implemented!
type TRoadLanesLaneSectionLrLaneAccessRestriction struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadLanesLaneSectionLrLaneBorder struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadLanesLaneSectionLrLaneHeight struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadLanesLaneSectionLrLaneMaterial struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadLanesLaneSectionLrLaneRule struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadLanesLaneSectionLrLaneSpeed struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadLanesLaneSectionLrLaneWidth struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadLanesLaneSectionRight struct {
	Lane TRoadLanesLaneSectionRightLane
}

// TODO: Doc formatting needs to be implemented!
type TRoadLanesLaneSectionRightLane struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadObjectsObjectLaneValidity struct {
}
