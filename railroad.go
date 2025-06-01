// Code generated; DO NOT EDIT.

package xodr

type ERoadRailroadSwitchPosition string

type EStationPlatformSegmentSide string

type EStationType string

// Container for all railroad definitions that shall be applied along a road.
// The available set of railroad elements is currently limited to the definiti
// on of switches. All other entries shall be covered with the existing elemen
// ts, for example, track definition by <road>, signal definition by <signal>,
// etc. Railroad-specific elements are defined against the background of stree
// tcar applications.
type RoadRailroad struct {
	OpenDriveElement
	Switch []*RoadRailroadSwitch
}

// Switches change the tracks for rail-bound vehicles. Switches guide the vehi
// cles into two directions only.
type RoadRailroadSwitch struct {
	OpenDriveElement
	MainTrack []*RoadRailroadSwitchMainTrack
	SideTrack []*RoadRailroadSwitchSideTrack
	Partner   []*RoadRailroadSwitchPartner
	Id        *string
	Name      *string
	Position  *ERoadRailroadSwitchPosition
}

// Main tracks form the primary course for rail bound traffic.
type RoadRailroadSwitchMainTrack struct {
	OpenDriveElement
	Dir *EElementDir
	Id  *string
	S   *GrEqZero
}

// Partner switches are two consistently set switches linked by a side track.
type RoadRailroadSwitchPartner struct {
	OpenDriveElement
	Id   *string
	Name *string
}

// Side tracks connect two switches that are placed on main tracks.
type RoadRailroadSwitchSideTrack struct {
	OpenDriveElement
	Dir *EElementDir
	Id  *string
	S   *GrEqZero
}

// Stations are places on the rail network where passengers enter and leave ra
// il-bound vehicles at platforms. May refer to multiple tracks and is therefo
// re defined on the same level as junctions.
type Station struct {
	OpenDriveElement
	Platform []*StationPlatform
	Id       *string
	Name     *string
	Type     *EStationType
}

// Platforms are essential parts of stations for passengers to enter and leave
// rail-bound vehicles. One or more railroad tracks reference one platform.
type StationPlatform struct {
	OpenDriveElement
	Segment []*StationPlatformSegment
	Id      *string
	Name    *string
}

// Segments are parts of platforms. Each <platform> element is valid on one or
// more track segments. The <segment> element must be specified.
type StationPlatformSegment struct {
	OpenDriveElement
	RoadId *string
	SEnd   *GrEqZero
	Side   *EStationPlatformSegmentSide
	SStart *GrEqZero
}
