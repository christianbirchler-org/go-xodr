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
	Control  []*ControllerControl
	Id       string
	Name     string
	Sequence int
}

// Provides information about a single signal within a signal group controlled
// by the corresponding controller.
type ControllerControl struct {
	OpenDriveElement
	SignalId string
	Type     string
}

// Signals are traffic signs, traffic lights, and specific road markings for t
// he control and regulation of road traffic. The <signals> element is the con
// tainer for all signals along a road.
type RoadSignals struct {
	OpenDriveElement
	Signal          []*RoadSignalsSignalRoad
	SignalReference []*RoadSignalsSpatialSignalReference
}

// Signals are not always separate signs on a single sheet of metal. Several s
// igns can be coupled on one board.
type RoadSignalsBoard struct {
	OpenDriveElement
	Validity   []*RoadObjectsObjectLaneValidity
	Dependency []*RoadSignalsSignalDependency
	Reference  []*RoadSignalsSignalReference
}

// A <sign> element on a static board defined in the local coordinate system o
// f the <signal> element. A <sign> element may have all attributes and child
// elements of a signal.
type RoadSignalsBoardSign struct {
	RoadSignalsSignal
	V float64
	Z float64
}

// A display area is the recommended position of the signal to be visualized i
// n the simulation. A display area is specified in the `<displayArea>` elemen
// t. A `<displayArea>` element is defined in the local coordinate system of t
// he `<signal>` element. The @index attribute can be used in ASAM OpenSCENARI
// O to reference the display area. In ASAM OpenSCENARIO a different local dis
// play area position may be specified.
type RoadSignalsDisplayArea struct {
	OpenDriveElement
	Height string
	Index  int
	V      float64
	Width  string
	Z      float64
}

// A signal along the road or on a static board.
type RoadSignalsSignal struct {
	OpenDriveElement
	Validity        []*RoadObjectsObjectLaneValidity
	Dependency      []*RoadSignalsSignalDependency
	Reference       []*RoadSignalsSignalReference
	StaticBoard     []*RoadSignalsStaticBoard
	VmsBoard        []*RoadSignalsVmsBoard
	Semantics       *SignalsSemantics
	Country         ECountryCode
	CountryRevision string
	Dynamic         YesNo
	Height          GrEqZero
	HOffset         float64
	Id              string
	Length          GrEqZero
	Name            string
	Orientation     EOrientation
	Pitch           float64
	Roll            float64
	Subtype         string
	Text            string
	Type            string
	Unit            EUnit
	Value           float64
	Width           GrEqZero
}

// Signal dependencies limit or extend the validity of one signal by an additi
// onal signal. For example, a speed limit sign of 60 km/h that is only valid
// for trucks, specified by a supplementary sign. One signal may have multiple
// dependencies.
type RoadSignalsSignalDependency struct {
	OpenDriveElement
	Id   string
	Type string
}

// Describes the reference point of the physical position in inertial coordina
// tes in cases where it deviates from the logical position. Defines the inert
// ial position.
type RoadSignalsSignalPositionInertial struct {
	OpenDriveElement
	Hdg   float64
	Pitch float64
	Roll  float64
	X     float64
	Y     float64
	Z     float64
}

// Describes the reference point of the physical position road coordinates in
// cases where it deviates from the logical position. Defines the position on
// the road.
type RoadSignalsSignalPositionRoad struct {
	OpenDriveElement
	HOffset float64
	Pitch   float64
	RoadId  string
	Roll    float64
	S       GrEqZero
	T       float64
	ZOffset float64
}

// Signal references link a signal to another signal or object. One signal may
// have multiple signal references. The signal reference term should not to be
// confused with the `<signalReference>` element that is used to link a signal
// to multiple roads.
type RoadSignalsSignalReference struct {
	OpenDriveElement
	ElementId   string
	ElementType ERoadSignalsSignalReferenceElementType
	Type        string
}

// Used to provide information about signals along a road. Consists of a main
// element and an optional lane validity element. The element for a signal is
// <signal>.
type RoadSignalsSignalRoad struct {
	RoadSignalsSignal
	S       GrEqZero
	T       float64
	ZOffset float64
}

// Refers to the same, that is, identical signal from multiple roads. The refe
// renced signals require a unique ID. The <signalReference> element consists
// of a main element and an optional lane validity element.
type RoadSignalsSpatialSignalReference struct {
	OpenDriveElement
	Validity    []*RoadObjectsObjectLaneValidity
	Id          string
	Orientation EOrientation
	S           GrEqZero
	T           float64
}

// A <signal> element that contains a <staticBoard> element. The signs that ar
// e displayed on a static board are defined as separate <sign> elements.
type RoadSignalsStaticBoard struct {
	RoadSignalsBoard
	Sign []*RoadSignalsBoardSign
}

// Variable message boards can change their values during the simulation in AS
// AM OpenSCENARIO. Variable message boards are switched off if they are not s
// pecified in ASAM OpenSCENARIO.
type RoadSignalsVmsBoard struct {
	RoadSignalsBoard
	DisplayArea   []*RoadSignalsDisplayArea
	DisplayHeight float64
	DisplayType   ERoadSignalsDisplayType
	DisplayWidth  float64
	V             float64
	Z             float64
}

// Variable message board references list all variable message boards that bel
// ong to the same gantry.
type SignalGroupVmsBoardReference struct {
	OpenDriveElement
	GroupIndex int
	SignalId   string
	VmsIndex   int
}

// On a gantry there can be one large variable message board or several smalle
// r variable message boards. ASAM OpenSCENARIO requires to treat a gantry tha
// t has one large variable message board or several smaller variable message
// boards the same way. Therefore variable message boards that are on the same
// gantry shall be grouped and their indexes shall be redefined if not unique.
type SignalGroupVmsGroup struct {
	OpenDriveElement
	VmsBoardReference []*SignalGroupVmsBoardReference
	Id                string
}

// Semantics are limited to traffic behavior that is specified just by signals
// in ASAM OpenDRIVE. Each traffic behavior is specified by a specific element
// .
type SignalsSemantics struct {
	OpenDriveElement
	Speed                    []*SignalsSemanticsSpeed
	Lane                     []*SignalsSemanticsLane
	Priority                 []*SignalsSemanticsPriority
	Prohibited               []*SignalsSemanticsProhibited
	Warning                  []*SignalsSemanticsWarning
	Routing                  []*SignalsSemanticsRouting
	Streetname               []*SignalsSemanticsStreetname
	Parking                  []*SignalsSemanticsParking
	Tourist                  []*SignalsSemanticsTourist
	SupplementaryTime        []*SignalsSemanticsSupplementaryTime
	SupplementaryAllows      []*SignalsSemanticsSupplementaryAllows
	SupplementaryProhibits   []*SignalsSemanticsSupplementaryProhibits
	SupplementaryDistance    []*SignalsSemanticsSupplementaryDistance
	SupplementaryEnvironment []*SignalsSemanticsSupplementaryEnvironment
	SupplementaryExplanatory []*SignalsSemanticsSupplementaryExplanatory
}

// Specifies lane regulations.
type SignalsSemanticsLane struct {
	OpenDriveElement
	Type ESignalsSemanticsLane
}

// Specifies parking regulations.
type SignalsSemanticsParking struct {
	OpenDriveElement
}

// Specifies priority regulations.
type SignalsSemanticsPriority struct {
	OpenDriveElement
	Type ESignalsSemanticsPriority
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
	Type  *ESignalsSemanticsSpeed
	Unit  *EUnitSpeed
	Value *float64
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
	Type  ESignalsSemanticsSupplementaryDistance
	Unit  EUnitDistance
	Value float64
}

// This signal semantic has no meaning on its own. It specifies under which en
// vironmental conditions a sign is valid.
type SignalsSemanticsSupplementaryEnvironment struct {
	OpenDriveElement
	Type ESignalsSemanticsSupplementaryEnvironment
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
	Type  ESignalsSemanticsSupplementaryTime
	Value float64
}

// Specifies tourist information.
type SignalsSemanticsTourist struct {
	OpenDriveElement
}

// Specifies warnings for traffic participant.
type SignalsSemanticsWarning struct {
	OpenDriveElement
}
