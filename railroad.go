// Code generated; DO NOT EDIT.

package xodr

type ERoadRailroadSwitchPosition struct {
}

type EStationPlatformSegmentSide struct {
}

type EStationType struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadRailroad struct {
	Switch TRoadRailroadSwitch
}

// TODO: Doc formatting needs to be implemented!
type TRoadRailroadSwitch struct {
	MainTrack TRoadRailroadSwitchMainTrack

	SideTrack TRoadRailroadSwitchSideTrack

	Partner TRoadRailroadSwitchPartner
}

// TODO: Doc formatting needs to be implemented!
type TRoadRailroadSwitchMainTrack struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadRailroadSwitchPartner struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadRailroadSwitchSideTrack struct {
}

// TODO: Doc formatting needs to be implemented!
type TStation struct {
	Platform TStationPlatform
}

// TODO: Doc formatting needs to be implemented!
type TStationPlatform struct {
	Segment TStationPlatformSegment
}

// TODO: Doc formatting needs to be implemented!
type TStationPlatformSegment struct {
}
