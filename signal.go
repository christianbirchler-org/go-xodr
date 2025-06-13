// Code generated; DO NOT EDIT.

package xodr

type ERoadSignalsDisplayType string

type ERoadSignalsSignalReferenceElementType string

type ESignalsSemanticsLane string

type ESignalsSemanticsPriority string

type ESignalsSemanticsSpeed string

type ESignalsSemanticsSupplementaryDistance string

type ESignalsSemanticsSupplementaryEnvironment string

type ESignalsSemanticsSupplementaryTime string

type EsignalsSemanticsSpeed struct {
}

type EunitSpeed struct {
}

// Controllers provide a signal program for a traffic signal or a signal group
// . The mapping of traffic signals to a signal group is done in t_controller.
// Dynamic content like the signal program itself is specified outside of this
// standard (i.e. in OpenSCENARIO).
type Controller struct {
	OpenDriveElement
	Control  []*ControllerControl `xml:"control"`
	Id       string               `xml:"id,attr"`
	Name     string               `xml:"name,attr"`
	Sequence int                  `xml:"sequence,attr"`
}

// Provides information about a single signal within a signal group controlled
// by the corresponding controller.
type ControllerControl struct {
	OpenDriveElement
	SignalId string `xml:"signalId,attr"`
	Type     string `xml:"type,attr"`
}

// Signals are traffic signs, traffic lights, and specific road markings for t
// he control and regulation of road traffic. The <signals> element is the con
// tainer for all signals along a road.
type RoadSignals struct {
	OpenDriveElement
	Signal          []*RoadSignalsSignalRoad             `xml:"signal"`
	SignalReference []*RoadSignalsSpatialSignalReference `xml:"signalReference"`
}

// Signals are not always separate signs on a single sheet of metal. Several s
// igns can be coupled on one board.
type RoadSignalsBoard struct {
	OpenDriveElement
	Validity   []*RoadObjectsObjectLaneValidity `xml:"validity"`
	Dependency []*RoadSignalsSignalDependency   `xml:"dependency"`
	Reference  []*RoadSignalsSignalReference    `xml:"reference"`
}

// A <sign> element on a static board defined in the local coordinate system o
// f the <signal> element. A <sign> element may have all attributes and child
// elements of a signal.
type RoadSignalsBoardSign struct {
	RoadSignalsSignal
	V float64 `xml:"v,attr"`
	Z float64 `xml:"z,attr"`
}

// A display area is the recommended position of the signal to be visualized i
// n the simulation. A display area is specified in the `<displayArea>` elemen
// t. A `<displayArea>` element is defined in the local coordinate system of t
// he `<signal>` element. The @index attribute can be used in ASAM OpenSCENARI
// O to reference the display area. In ASAM OpenSCENARIO a different local dis
// play area position may be specified.
type RoadSignalsDisplayArea struct {
	OpenDriveElement
	Height string  `xml:"height,attr"`
	Index  int     `xml:"index,attr"`
	V      float64 `xml:"v,attr"`
	Width  string  `xml:"width,attr"`
	Z      float64 `xml:"z,attr"`
}

// A signal along the road or on a static board.
type RoadSignalsSignal struct {
	OpenDriveElement
	Validity        []*RoadObjectsObjectLaneValidity `xml:"validity"`
	Dependency      []*RoadSignalsSignalDependency   `xml:"dependency"`
	Reference       []*RoadSignalsSignalReference    `xml:"reference"`
	StaticBoard     []*RoadSignalsStaticBoard        `xml:"staticBoard"`
	VmsBoard        []*RoadSignalsVmsBoard           `xml:"vmsBoard"`
	Semantics       *SignalsSemantics                `xml:"semantics"`
	Country         ECountryCode                     `xml:"country,attr"`
	CountryRevision string                           `xml:"countryRevision,attr"`
	Dynamic         YesNo                            `xml:"dynamic,attr"`
	Height          GrEqZero                         `xml:"height,attr"`
	HOffset         float64                          `xml:"hOffset,attr"`
	Id              string                           `xml:"id,attr"`
	Length          GrEqZero                         `xml:"length,attr"`
	Name            string                           `xml:"name,attr"`
	Orientation     EOrientation                     `xml:"orientation,attr"`
	Pitch           float64                          `xml:"pitch,attr"`
	Roll            float64                          `xml:"roll,attr"`
	Subtype         string                           `xml:"subtype,attr"`
	Text            string                           `xml:"text,attr"`
	Type            string                           `xml:"type,attr"`
	Unit            EUnit                            `xml:"unit,attr"`
	Value           float64                          `xml:"value,attr"`
	Width           GrEqZero                         `xml:"width,attr"`
}

// Signal dependencies limit or extend the validity of one signal by an additi
// onal signal. For example, a speed limit sign of 60 km/h that is only valid
// for trucks, specified by a supplementary sign. One signal may have multiple
// dependencies.
type RoadSignalsSignalDependency struct {
	OpenDriveElement
	Id   string `xml:"id,attr"`
	Type string `xml:"type,attr"`
}

// Describes the reference point of the physical position in inertial coordina
// tes in cases where it deviates from the logical position. Defines the inert
// ial position.
type RoadSignalsSignalPositionInertial struct {
	OpenDriveElement
	Hdg   float64 `xml:"hdg,attr"`
	Pitch float64 `xml:"pitch,attr"`
	Roll  float64 `xml:"roll,attr"`
	X     float64 `xml:"x,attr"`
	Y     float64 `xml:"y,attr"`
	Z     float64 `xml:"z,attr"`
}

// Describes the reference point of the physical position road coordinates in
// cases where it deviates from the logical position. Defines the position on
// the road.
type RoadSignalsSignalPositionRoad struct {
	OpenDriveElement
	HOffset float64  `xml:"hOffset,attr"`
	Pitch   float64  `xml:"pitch,attr"`
	RoadId  string   `xml:"roadId,attr"`
	Roll    float64  `xml:"roll,attr"`
	S       GrEqZero `xml:"s,attr"`
	T       float64  `xml:"t,attr"`
	ZOffset float64  `xml:"zOffset,attr"`
}

// Signal references link a signal to another signal or object. One signal may
// have multiple signal references. The signal reference term should not to be
// confused with the `<signalReference>` element that is used to link a signal
// to multiple roads.
type RoadSignalsSignalReference struct {
	OpenDriveElement
	ElementId   string                                 `xml:"elementId,attr"`
	ElementType ERoadSignalsSignalReferenceElementType `xml:"elementType,attr"`
	Type        string                                 `xml:"type,attr"`
}

// Used to provide information about signals along a road. Consists of a main
// element and an optional lane validity element. The element for a signal is
// <signal>.
type RoadSignalsSignalRoad struct {
	RoadSignalsSignal
	S       GrEqZero `xml:"s,attr"`
	T       float64  `xml:"t,attr"`
	ZOffset float64  `xml:"zOffset,attr"`
}

// Refers to the same, that is, identical signal from multiple roads. The refe
// renced signals require a unique ID. The <signalReference> element consists
// of a main element and an optional lane validity element.
type RoadSignalsSpatialSignalReference struct {
	OpenDriveElement
	Validity    []*RoadObjectsObjectLaneValidity `xml:"validity"`
	Id          string                           `xml:"id,attr"`
	Orientation EOrientation                     `xml:"orientation,attr"`
	S           GrEqZero                         `xml:"s,attr"`
	T           float64                          `xml:"t,attr"`
}

// A <signal> element that contains a <staticBoard> element. The signs that ar
// e displayed on a static board are defined as separate <sign> elements.
type RoadSignalsStaticBoard struct {
	RoadSignalsBoard
	Sign []*RoadSignalsBoardSign `xml:"sign"`
}

// Variable message boards can change their values during the simulation in AS
// AM OpenSCENARIO. Variable message boards are switched off if they are not s
// pecified in ASAM OpenSCENARIO.
type RoadSignalsVmsBoard struct {
	RoadSignalsBoard
	DisplayArea   []*RoadSignalsDisplayArea `xml:"displayArea"`
	DisplayHeight float64                   `xml:"displayHeight,attr"`
	DisplayType   ERoadSignalsDisplayType   `xml:"displayType,attr"`
	DisplayWidth  float64                   `xml:"displayWidth,attr"`
	V             float64                   `xml:"v,attr"`
	Z             float64                   `xml:"z,attr"`
}

// Variable message board references list all variable message boards that bel
// ong to the same gantry.
type SignalGroupVmsBoardReference struct {
	OpenDriveElement
	GroupIndex int    `xml:"groupIndex,attr"`
	SignalId   string `xml:"signalId,attr"`
	VmsIndex   int    `xml:"vmsIndex,attr"`
}

// On a gantry there can be one large variable message board or several smalle
// r variable message boards. ASAM OpenSCENARIO requires to treat a gantry tha
// t has one large variable message board or several smaller variable message
// boards the same way. Therefore variable message boards that are on the same
// gantry shall be grouped and their indexes shall be redefined if not unique.
type SignalGroupVmsGroup struct {
	OpenDriveElement
	VmsBoardReference []*SignalGroupVmsBoardReference `xml:"vmsBoardReference"`
	Id                string                          `xml:"id,attr"`
}

// Semantics are limited to traffic behavior that is specified just by signals
// in ASAM OpenDRIVE. Each traffic behavior is specified by a specific element
// .
type SignalsSemantics struct {
	OpenDriveElement
	Speed                    []*SignalsSemanticsSpeed                    `xml:"speed"`
	Lane                     []*SignalsSemanticsLane                     `xml:"lane"`
	Priority                 []*SignalsSemanticsPriority                 `xml:"priority"`
	Prohibited               []*SignalsSemanticsProhibited               `xml:"prohibited"`
	Warning                  []*SignalsSemanticsWarning                  `xml:"warning"`
	Routing                  []*SignalsSemanticsRouting                  `xml:"routing"`
	Streetname               []*SignalsSemanticsStreetname               `xml:"streetname"`
	Parking                  []*SignalsSemanticsParking                  `xml:"parking"`
	Tourist                  []*SignalsSemanticsTourist                  `xml:"tourist"`
	SupplementaryTime        []*SignalsSemanticsSupplementaryTime        `xml:"supplementaryTime"`
	SupplementaryAllows      []*SignalsSemanticsSupplementaryAllows      `xml:"supplementaryAllows"`
	SupplementaryProhibits   []*SignalsSemanticsSupplementaryProhibits   `xml:"supplementaryProhibits"`
	SupplementaryDistance    []*SignalsSemanticsSupplementaryDistance    `xml:"supplementaryDistance"`
	SupplementaryEnvironment []*SignalsSemanticsSupplementaryEnvironment `xml:"supplementaryEnvironment"`
	SupplementaryExplanatory []*SignalsSemanticsSupplementaryExplanatory `xml:"supplementaryExplanatory"`
}

// Specifies lane regulations.
type SignalsSemanticsLane struct {
	OpenDriveElement
	Type ESignalsSemanticsLane `xml:"type,attr"`
}

// Specifies parking regulations.
type SignalsSemanticsParking struct {
	OpenDriveElement
}

// Specifies priority regulations.
type SignalsSemanticsPriority struct {
	OpenDriveElement
	Type ESignalsSemanticsPriority `xml:"type,attr"`
}

// Specifies that certain types of traffic participants are not allowed to ent
// er. Signal semantics for traffic participants in {THIS_STANDARD} are curren
// tly not defined because traffic participants are not harmonized for all sta
// ndards.
type SignalsSemanticsProhibited struct {
	OpenDriveElement
}

// Specifies routing information.
type SignalsSemanticsRouting struct {
	OpenDriveElement
}

// Specifies speed regulations.
type SignalsSemanticsSpeed struct {
	OpenDriveElement
	Type  *ESignalsSemanticsSpeed `xml:"type"`
	Unit  *EUnitSpeed             `xml:"unit"`
	Value *float64                `xml:"value"`
}

// Specifies the name of a street.
type SignalsSemanticsStreetname struct {
	OpenDriveElement
}

// This signal semantic has no meaning on its own. It specifies the type of th
// e traffic participant an exception is made for. Signal semantics for traffi
// c participants in {THIS_STANDARD} are currently not specified because traff
// ic participants are not harmonized for all standards.
type SignalsSemanticsSupplementaryAllows struct {
	OpenDriveElement
}

// This signal semantic has no meaning on its own. It specifies the distance a
// fter a sign becomes valid or the range in which the sign is valid.
type SignalsSemanticsSupplementaryDistance struct {
	OpenDriveElement
	Type  ESignalsSemanticsSupplementaryDistance `xml:"type,attr"`
	Unit  EUnitDistance                          `xml:"unit,attr"`
	Value float64                                `xml:"value,attr"`
}

// This signal semantic has no meaning on its own. It specifies under which en
// vironmental conditions a sign is valid.
type SignalsSemanticsSupplementaryEnvironment struct {
	OpenDriveElement
	Type ESignalsSemanticsSupplementaryEnvironment `xml:"type,attr"`
}

// This signal semantic has no meaning on its own. It specifies explanations f
// or a sign.
type SignalsSemanticsSupplementaryExplanatory struct {
	OpenDriveElement
}

// This signal semantic has no meaning on its own. It specifies the type of th
// e traffic participant a restriction is made for. Signal semantics for traff
// ic participants in {THIS_STANDARD} are currently not specified because traf
// fic participants are not harmonized for all standards.
type SignalsSemanticsSupplementaryProhibits struct {
	OpenDriveElement
}

// This signal semantic has no meaning on its own. It specifies the time or da
// te a sign is valid.
type SignalsSemanticsSupplementaryTime struct {
	OpenDriveElement
	Type  ESignalsSemanticsSupplementaryTime `xml:"type,attr"`
	Value float64                            `xml:"value,attr"`
}

// Specifies tourist information.
type SignalsSemanticsTourist struct {
	OpenDriveElement
}

// Specifies warnings for traffic participant.
type SignalsSemanticsWarning struct {
	OpenDriveElement
}
