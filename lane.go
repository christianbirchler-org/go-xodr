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

// Lanes are an essential part of all roads. Lanes are attached to the road re
// ference line and are defined from inside to outside. Lanes contain a series
// of lane section elements that define the characteristics of the road cross
// sections with respect to the lanes along the road reference line.
type RoadLanes struct {
	OpenDriveElement
	LaneOffset  []*RoadLanesLaneOffset  `xml:"laneOffset"`
	LaneSection []*RoadLanesLaneSection `xml:"laneSection"`
}

// Lane offsets shift the center lane away from the road reference line.
type RoadLanesLaneOffset struct {
	OpenDriveElement
}

// A lane section splits a road into multiple parts whenever the number of lan
// es or their function changes. The distance between two succeeding lane sect
// ions shall not be zero. For easier navigation through an ASAM OpenDRIVE roa
// d description, the lanes within a lane section are grouped into left, cente
// r, and right lanes. Each lane section shall contain one <center> element an
// d at least one <right> or <left> element.
type RoadLanesLaneSection struct {
	OpenDriveElement
	Left   *RoadLanesLaneSectionLeft   `xml:"left"`
	Center *RoadLanesLaneSectionCenter `xml:"center"`
	Right  *RoadLanesLaneSectionRight  `xml:"right"`
}

// Contains the center lane, which must be defined for all roads.
type RoadLanesLaneSectionCenter struct {
	OpenDriveElement
	Lane *RoadLanesLaneSectionCenterLane `xml:"lane"`
}

// Center lane element with ID zero. Has no width attribute. Mainly used for r
// oad marks.
type RoadLanesLaneSectionCenterLane struct {
	Id       int       `xml:"id,attr"`
	Level    Bool      `xml:"level,attr"`
	Type     ELaneType `xml:"type,attr"`
	Link     RoadLanesLaneSectionLcrLaneLink
	RoadMark RoadLanesLaneSectionLcrLaneRoadMark
}

// For links between lanes with an identical road reference line, the lane pre
// decessor and successor information provide the IDs of lanes on the precedin
// g or following lane section. For links between lanes with different road re
// ference line,  the lane predecessor and successor information provide the I
// Ds of lanes on the first or last lane section of the other road reference l
// ine depending on the contact point of the road linkage. This element may on
// ly be omitted, if lanes end at a junction or have no physical link.
type RoadLanesLaneSectionLcrLaneLink struct {
	OpenDriveElement
	Predecessor []*RoadLanesLaneSectionLcrLaneLinkPredecessorSuccessor `xml:"predecessor"`
	Successor   []*RoadLanesLaneSectionLcrLaneLinkPredecessorSuccessor `xml:"successor"`
}

//
type RoadLanesLaneSectionLcrLaneLinkPredecessorSuccessor struct {
	OpenDriveElement
}

// Defines the style of the line at the outer border of a lane. The style of t
// he center line that separates left and right lanes is determined by the roa
// d mark element for the center lane.
type RoadLanesLaneSectionLcrLaneRoadMark struct {
	OpenDriveElement
	Sway     []*RoadLanesLaneSectionLcrLaneRoadMarkSway   `xml:"sway"`
	Type     *RoadLanesLaneSectionLcrLaneRoadMarkType     `xml:"type"`
	Explicit *RoadLanesLaneSectionLcrLaneRoadMarkExplicit `xml:"explicit"`
}

// Irregular road markings that cannot be described by repetitive line pattern
// s may be described by individual road marking elements. These explicit defi
// nitions also contain <line> elements for the line definition, however, thes
// e lines will not be repeated automatically as in repetitive road marking ty
// pes. In ASAM OpenDRIVE, irregular road marking types and lines are represen
// ted by <explicit> elements within elements. The line definitions are contai
// ned in <line> elements within the <explicit> element. The <explicit> elemen
// t should specifically be used for measurement data.
type RoadLanesLaneSectionLcrLaneRoadMarkExplicit struct {
	OpenDriveElement
	Line []*RoadLanesLaneSectionLcrLaneRoadMarkExplicitLine `xml:"line"`
}

// Specifies a single line in an explicit road mark definition.
type RoadLanesLaneSectionLcrLaneRoadMarkExplicitLine struct {
	OpenDriveElement
}

// Relocates the lateral reference position for the following (explicit) type
// definition and thus defines an offset. The sway offset is relative to the n
// ominal reference position of the lane marking, meaning the lane border.
type RoadLanesLaneSectionLcrLaneRoadMarkSway struct {
	OpenDriveElement
}

// Each type definition shall contain one or more line definitions with additi
// onal information about the lines that the road mark is composed of.
type RoadLanesLaneSectionLcrLaneRoadMarkType struct {
	OpenDriveElement
	Line []*RoadLanesLaneSectionLcrLaneRoadMarkTypeLine `xml:"line"`
}

// A road mark may consist of one or more elements. Multiple elements are usua
// lly positioned side-by-side. A line definition is valid for a given length
// of the lane and will be repeated automatically.
type RoadLanesLaneSectionLcrLaneRoadMarkTypeLine struct {
	OpenDriveElement
}

// Contains all lanes left to the center lane.
type RoadLanesLaneSectionLeft struct {
	OpenDriveElement
	Lane []*RoadLanesLaneSectionLeftLane `xml:"lane"`
}

// Left lanes numbered with positive IDs in ascending order from center lane t
// o left border.
type RoadLanesLaneSectionLeftLane struct {
	RoadLanesLaneSectionLrLane
}

// Lane elements are included in left/center/right elements. Lane elements sho
// uld represent the lanes from left to right, that is, with descending ID.
type RoadLanesLaneSectionLrLane struct {
	OpenDriveElement
	Link     *RoadLanesLaneSectionLcrLaneLink       `xml:"link"`
	RoadMark []*RoadLanesLaneSectionLcrLaneRoadMark `xml:"roadMark"`
	Material []*RoadLanesLaneSectionLrLaneMaterial  `xml:"material"`
	Speed    []*RoadLanesLaneSectionLrLaneSpeed     `xml:"speed"`
	Access   []*RoadLanesLaneSectionLrLaneAccess    `xml:"access"`
	Height   []*RoadLanesLaneSectionLrLaneHeight    `xml:"height"`
	Rule     []*RoadLanesLaneSectionLrLaneRule      `xml:"rule"`
}

// Defines access restrictions for certain types of road users. Each element i
// s valid in direction of the increasing s coordinate until a new element is
// defined. If multiple elements are defined, they shall be listed in ascendin
// g order.
type RoadLanesLaneSectionLrLaneAccess struct {
	OpenDriveElement
	Restriction []*RoadLanesLaneSectionLrLaneAccessRestriction `xml:"restriction"`
}

// Defines access restrictions for certain types of road users. Each restricti
// on element defines one type that is either allowed or denied according to t
// he parent access element.
type RoadLanesLaneSectionLrLaneAccessRestriction struct {
	OpenDriveElement
}

// Lane borders set the width of lanes. Lane borders describe the outer limits
// of lanes, independent of the parameters of their inner borders. In this cas
// e, inner lanes are defined as lanes which have the same sign for their ID a
// s the lane currently defined, but with a smaller absolute value for their I
// D. Especially when road data is derived from automatic measurements, this t
// ype of definition is easier than specifying the lane width because it avoid
// s creating many lane sections. Lane width and lane border elements are mutu
// ally exclusive within the same lane group. If both width and lane border el
// ements are present for a lane section in the ASAM OpenDRIVE file, the appli
// cation shall use the information from the <width> elements.
type RoadLanesLaneSectionLrLaneBorder struct {
	OpenDriveElement
}

// Lane heights elevate lanes along the h-coordinate within a lane section ind
// ependent from the road elevation.  Lane height is used to implement small-s
// cale elevation such as raising pedestrian walkways. Lane height is specifie
// d as offset from the road (including elevation, superelevation, shape, cros
// s section surface) in h-direction.
type RoadLanesLaneSectionLrLaneHeight struct {
	OpenDriveElement
}

// Stores information about the material of lanes. Each element is valid until
// a new element is defined. If multiple elements are defined, they must be li
// sted in ascending order.
type RoadLanesLaneSectionLrLaneMaterial struct {
	OpenDriveElement
}

// Used to add rules that are not covered by any of the other lane attributes
// that are described in this specification.
type RoadLanesLaneSectionLrLaneRule struct {
	OpenDriveElement
}

// Defines the maximum allowed speed on a given lane. Each element is valid in
// direction of the increasing s-coordinate until a new element is defined.
type RoadLanesLaneSectionLrLaneSpeed struct {
	OpenDriveElement
}

// Lane widths widen or narrow lanes along the t-coordinate within lane sectio
// ns. Lane width and lane border elements are mutually exclusive within the s
// ame lane group. If both width and lane border elements are present for a la
// ne section in the ASAM OpenDRIVE file, the application must use the informa
// tion from the <width> elements.
type RoadLanesLaneSectionLrLaneWidth struct {
	OpenDriveElement
}

// Contains all lanes right to the center lane.
type RoadLanesLaneSectionRight struct {
	OpenDriveElement
	Lane []*RoadLanesLaneSectionRightLane `xml:"lane"`
}

// Right lanes numbered with negative IDs in descending order from center lane
// to right border.
type RoadLanesLaneSectionRightLane struct {
	RoadLanesLaneSectionLrLane
}

// Lane validities restrict signals and objects to specific lanes.
type RoadObjectsObjectLaneValidity struct {
	OpenDriveElement
}
