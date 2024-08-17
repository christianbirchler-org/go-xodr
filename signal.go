// Code generated; DO NOT EDIT.

package xodr

type ERoadSignalsDisplayType struct {
}

type ERoadSignalsSignalReferenceElementType struct {
}

type ESignalsSemanticsLane struct {
}

type ESignalsSemanticsPriority struct {
}

type ESignalsSemanticsSpeed struct {
}

type ESignalsSemanticsSupplementaryDistance struct {
}

type ESignalsSemanticsSupplementaryEnvironment struct {
}

type ESignalsSemanticsSupplementaryTime struct {
}

type EsignalsSemanticsSpeed struct {
}

type EunitSpeed struct {
}

// TODO: Doc formatting needs to be implemented!
type Controller struct {
	OpenDriveElement
	Control  *ControllerControl
	Id       string
	Name     string
	Sequence int
}

// TODO: Doc formatting needs to be implemented!
type ControllerControl struct {
	OpenDriveElement
	SignalId string
	Type     string
}

// TODO: Doc formatting needs to be implemented!
type RoadSignals struct {
	OpenDriveElement
	Signal          *RoadSignalsSignalRoad
	SignalReference *RoadSignalsSpatialSignalReference
}

// TODO: Doc formatting needs to be implemented!
type RoadSignalsBoard struct {
	OpenDriveElement
	Validity   *RoadObjectsObjectLaneValidity
	Dependency *RoadSignalsSignalDependency
	Reference  *RoadSignalsSignalReference
}

// TODO: Doc formatting needs to be implemented!
type RoadSignalsBoardSign struct {
	RoadSignalsSignal
	V float64
	Z float64
}

// TODO: Doc formatting needs to be implemented!
type RoadSignalsDisplayArea struct {
	OpenDriveElement
	Height string
	Index  int
	V      float64
	Width  string
	Z      float64
}

// TODO: Doc formatting needs to be implemented!
type RoadSignalsSignal struct {
	OpenDriveElement
	Validity        *RoadObjectsObjectLaneValidity
	Dependency      *RoadSignalsSignalDependency
	Reference       *RoadSignalsSignalReference
	StaticBoard     *RoadSignalsStaticBoard
	VmsBoard        *RoadSignalsVmsBoard
	Semantics       *SignalsSemantics
	Country         ECountryCode
	CountryRevision string
	Dynamic         TYesNo
	Height          TGrEqZero
	HOffset         float64
	Id              string
	Length          TGrEqZero
	Name            string
	Orientation     EOrientation
	Pitch           float64
	Roll            float64
	Subtype         string
	Text            string
	Type            string
	Unit            EUnit
	Value           float64
	Width           TGrEqZero
}

// TODO: Doc formatting needs to be implemented!
type RoadSignalsSignalDependency struct {
	OpenDriveElement
	Id   string
	Type string
}

// TODO: Doc formatting needs to be implemented!
type RoadSignalsSignalPositionInertial struct {
	OpenDriveElement
	Hdg   float64
	Pitch float64
	Roll  float64
	X     float64
	Y     float64
	Z     float64
}

// TODO: Doc formatting needs to be implemented!
type RoadSignalsSignalPositionRoad struct {
	OpenDriveElement
	HOffset float64
	Pitch   float64
	RoadId  string
	Roll    float64
	S       TGrEqZero
	T       float64
	ZOffset float64
}

// TODO: Doc formatting needs to be implemented!
type RoadSignalsSignalReference struct {
	OpenDriveElement
	ElementId   string
	ElementType ERoadSignalsSignalReferenceElementType
	Type        string
}

// TODO: Doc formatting needs to be implemented!
type RoadSignalsSignalRoad struct {
	RoadSignalsSignal
	S       TGrEqZero
	T       float64
	ZOffset float64
}

// TODO: Doc formatting needs to be implemented!
type RoadSignalsSpatialSignalReference struct {
	OpenDriveElement
	Validity    *RoadObjectsObjectLaneValidity
	Id          string
	Orientation EOrientation
	S           TGrEqZero
	T           float64
}

// TODO: Doc formatting needs to be implemented!
type RoadSignalsStaticBoard struct {
	RoadSignalsBoard
	Sign *RoadSignalsBoardSign
}

// TODO: Doc formatting needs to be implemented!
type RoadSignalsVmsBoard struct {
	RoadSignalsBoard
	DisplayArea   *RoadSignalsDisplayArea
	DisplayHeight float64
	DisplayType   ERoadSignalsDisplayType
	DisplayWidth  float64
	V             float64
	Z             float64
}

// TODO: Doc formatting needs to be implemented!
type SignalGroupVmsBoardReference struct {
	OpenDriveElement
	GroupIndex int
	SignalId   string
	VmsIndex   int
}

// TODO: Doc formatting needs to be implemented!
type SignalGroupVmsGroup struct {
	OpenDriveElement
	VmsBoardReference *SignalGroupVmsBoardReference
	Id                string
}

// TODO: Doc formatting needs to be implemented!
type SignalsSemantics struct {
	OpenDriveElement
	Speed                    *SignalsSemanticsSpeed
	Lane                     *SignalsSemanticsLane
	Priority                 *SignalsSemanticsPriority
	Prohibited               *SignalsSemanticsProhibited
	Warning                  *SignalsSemanticsWarning
	Routing                  *SignalsSemanticsRouting
	Streetname               *SignalsSemanticsStreetname
	Parking                  *SignalsSemanticsParking
	Tourist                  *SignalsSemanticsTourist
	SupplementaryTime        *SignalsSemanticsSupplementaryTime
	SupplementaryAllows      *SignalsSemanticsSupplementaryAllows
	SupplementaryProhibits   *SignalsSemanticsSupplementaryProhibits
	SupplementaryDistance    *SignalsSemanticsSupplementaryDistance
	SupplementaryEnvironment *SignalsSemanticsSupplementaryEnvironment
	SupplementaryExplanatory *SignalsSemanticsSupplementaryExplanatory
}

// TODO: Doc formatting needs to be implemented!
type SignalsSemanticsLane struct {
	OpenDriveElement
	Type ESignalsSemanticsLane
}

// TODO: Doc formatting needs to be implemented!
type SignalsSemanticsParking struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type SignalsSemanticsPriority struct {
	OpenDriveElement
	Type ESignalsSemanticsPriority
}

// TODO: Doc formatting needs to be implemented!
type SignalsSemanticsProhibited struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type SignalsSemanticsRouting struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type SignalsSemanticsSpeed struct {
	OpenDriveElement
	Type  *EsignalsSemanticsSpeed
	Unit  *EunitSpeed
	Value *float64
}

// TODO: Doc formatting needs to be implemented!
type SignalsSemanticsStreetname struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type SignalsSemanticsSupplementaryAllows struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type SignalsSemanticsSupplementaryDistance struct {
	OpenDriveElement
	Type  ESignalsSemanticsSupplementaryDistance
	Unit  EUnitDistance
	Value float64
}

// TODO: Doc formatting needs to be implemented!
type SignalsSemanticsSupplementaryEnvironment struct {
	OpenDriveElement
	Type ESignalsSemanticsSupplementaryEnvironment
}

// TODO: Doc formatting needs to be implemented!
type SignalsSemanticsSupplementaryExplanatory struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type SignalsSemanticsSupplementaryProhibits struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type SignalsSemanticsSupplementaryTime struct {
	OpenDriveElement
	Type  ESignalsSemanticsSupplementaryTime
	Value float64
}

// TODO: Doc formatting needs to be implemented!
type SignalsSemanticsTourist struct {
	OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type SignalsSemanticsWarning struct {
	OpenDriveElement
}
