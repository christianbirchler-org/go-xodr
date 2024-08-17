// Code generated; DO NOT EDIT.

package xodr

type ERoadRailroadSwitchPosition struct {
}

type EStationPlatformSegmentSide struct {
}

type EStationType struct {
}

// TODO: Doc formatting needs to be implemented!
type RoadRailroad struct {
	OpenDriveElement
	Switch RoadRailroadSwitch
}

// TODO: Doc formatting needs to be implemented!
type RoadRailroadSwitch struct {
	OpenDriveElement
	MainTrack RoadRailroadSwitchMainTrack
	SideTrack RoadRailroadSwitchSideTrack
	Partner   RoadRailroadSwitchPartner
	Id        string
	Name      string
	Position  ERoadRailroadSwitchPosition
}

// TODO: Doc formatting needs to be implemented!
type RoadRailroadSwitchMainTrack struct {
	OpenDriveElement
	Dir EElementDir
	Id  string
	S   TGrEqZero
}

// TODO: Doc formatting needs to be implemented!
type RoadRailroadSwitchPartner struct {
	OpenDriveElement
	Id   string
	Name string
}

// TODO: Doc formatting needs to be implemented!
type RoadRailroadSwitchSideTrack struct {
	OpenDriveElement
	Dir EElementDir
	Id  string
	S   TGrEqZero
}

// TODO: Doc formatting needs to be implemented!
type Station struct {
	OpenDriveElement
	Platform StationPlatform
	Id       string
	Name     string
	Type     EStationType
}

// TODO: Doc formatting needs to be implemented!
type StationPlatform struct {
	OpenDriveElement
	Segment StationPlatformSegment
	Id      string
	Name    string
}

// TODO: Doc formatting needs to be implemented!
type StationPlatformSegment struct {
	OpenDriveElement
	RoadId string
	SEnd   TGrEqZero
	Side   EStationPlatformSegmentSide
	SStart TGrEqZero
}
