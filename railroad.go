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
	Switch []*RoadRailroadSwitch `xml:"switch"`
}

// Switches change the tracks for rail-bound vehicles. Switches guide the vehi
// cles into two directions only.
type RoadRailroadSwitch struct {
	OpenDriveElement
	MainTrack *RoadRailroadSwitchMainTrack `xml:"mainTrack"`
	SideTrack *RoadRailroadSwitchSideTrack `xml:"sideTrack"`
	Partner   *RoadRailroadSwitchPartner   `xml:"partner"`
	Id        *string                      `xml:"id,attr"`
	Name      *string                      `xml:"name,attr"`
	Position  *ERoadRailroadSwitchPosition `xml:"position,attr"`
}

// Main tracks form the primary course for rail bound traffic.
type RoadRailroadSwitchMainTrack struct {
	OpenDriveElement
	Dir *EElementDir `xml:"dir,attr"`
	Id  *string      `xml:"id,attr"`
	S   *GrEqZero    `xml:"s,attr"`
}

// Partner switches are two consistently set switches linked by a side track.
type RoadRailroadSwitchPartner struct {
	OpenDriveElement
	Id   *string `xml:"id,attr"`
	Name *string `xml:"name,attr"`
}

// Side tracks connect two switches that are placed on main tracks.
type RoadRailroadSwitchSideTrack struct {
	OpenDriveElement
	Dir *EElementDir `xml:"dir,attr"`
	Id  *string      `xml:"id,attr"`
	S   *GrEqZero    `xml:"s,attr"`
}

// Stations are places on the rail network where passengers enter and leave ra
// il-bound vehicles at platforms. May refer to multiple tracks and is therefo
// re defined on the same level as junctions.
type Station struct {
	OpenDriveElement
	Platform []*StationPlatform `xml:"platform"`
	Id       *string            `xml:"id,attr"`
	Name     *string            `xml:"name,attr"`
	Type     *EStationType      `xml:"type,attr"`
}

// Platforms are essential parts of stations for passengers to enter and leave
// rail-bound vehicles. One or more railroad tracks reference one platform.
type StationPlatform struct {
	OpenDriveElement
	Segment []*StationPlatformSegment `xml:"segment"`
	Id      *string                   `xml:"id,attr"`
	Name    *string                   `xml:"name,attr"`
}

// Segments are parts of platforms. Each <platform> element is valid on one or
// more track segments. The <segment> element must be specified.
type StationPlatformSegment struct {
	OpenDriveElement
	RoadId *string                      `xml:"roadId,attr"`
	SEnd   *GrEqZero                    `xml:"sEnd,attr"`
	Side   *EStationPlatformSegmentSide `xml:"side,attr"`
	SStart *GrEqZero                    `xml:"sStart,attr"`
}
