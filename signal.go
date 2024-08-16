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

// TODO: Doc formatting needs to be implemented!
type TController struct {
	Control TControllerControl
}

// TODO: Doc formatting needs to be implemented!
type TControllerControl struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadSignals struct {
	Signal TRoadSignalsSignalRoad

	SignalReference TRoadSignalsSignalReference
}

// TODO: Doc formatting needs to be implemented!
type TRoadSignalsBoard struct {
	Validity TRoadObjectsObjectLaneValidity

	Dependency TRoadSignalsSignalDependency

	Reference TRoadSignalsSignalReference
}

// TODO: Doc formatting needs to be implemented!
type TRoadSignalsBoardSign struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadSignalsDisplayArea struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadSignalsSignal struct {
	Validity TRoadObjectsObjectLaneValidity

	Dependency TRoadSignalsSignalDependency

	Reference TRoadSignalsSignalReference

	StaticBoard TRoadSignalsStaticBoard

	VmsBoard TRoadSignalsVmsBoard

	Semantics TSignalsSemantics
}

// TODO: Doc formatting needs to be implemented!
type TRoadSignalsSignalDependency struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadSignalsSignalPositionInertial struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadSignalsSignalPositionRoad struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadSignalsSignalReference struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadSignalsSignalRoad struct {
}

// TODO: Doc formatting needs to be implemented!
type TRoadSignalsStaticBoard struct {
	Sign TRoadSignalsBoardSign
}

// TODO: Doc formatting needs to be implemented!
type TRoadSignalsVmsBoard struct {
	DisplayArea TRoadSignalsDisplayArea
}

// TODO: Doc formatting needs to be implemented!
type TSignalGroupVmsBoardReference struct {
}

// TODO: Doc formatting needs to be implemented!
type TSignalGroupVmsGroup struct {
	VmsBoardReference TSignalGroupVmsBoardReference

	OpenDriveElement OpenDriveElement
}

// TODO: Doc formatting needs to be implemented!
type TSignalsSemantics struct {
	Speed TSignalsSemanticsSpeed

	Lane TSignalsSemanticsLane

	Priority TSignalsSemanticsPriority

	Prohibited TSignalsSemanticsProhibited

	Warning TSignalsSemanticsWarning

	Routing TSignalsSemanticsRouting

	Streetname TSignalsSemanticsStreetname

	Parking TSignalsSemanticsParking

	Tourist TSignalsSemanticsTourist

	SupplementaryTime TSignalsSemanticsSupplementaryTime

	SupplementaryAllows TSignalsSemanticsSupplementaryAllows

	SupplementaryProhibits TSignalsSemanticsSupplementaryProhibits

	SupplementaryDistance TSignalsSemanticsSupplementaryDistance

	SupplementaryEnvironment TSignalsSemanticsSupplementaryEnvironment

	SupplementaryExplanatory TSignalsSemanticsSupplementaryExplanatory
}

// TODO: Doc formatting needs to be implemented!
type TSignalsSemanticsLane struct {
}

// TODO: Doc formatting needs to be implemented!
type TSignalsSemanticsParking struct {
}

// TODO: Doc formatting needs to be implemented!
type TSignalsSemanticsPriority struct {
}

// TODO: Doc formatting needs to be implemented!
type TSignalsSemanticsProhibited struct {
}

// TODO: Doc formatting needs to be implemented!
type TSignalsSemanticsRouting struct {
}

// TODO: Doc formatting needs to be implemented!
type TSignalsSemanticsSpeed struct {
	Type ESignalsSemanticsSpeed

	Unit EUnitSpeed

	Value float64
}

// TODO: Doc formatting needs to be implemented!
type TSignalsSemanticsStreetname struct {
}

// TODO: Doc formatting needs to be implemented!
type TSignalsSemanticsSupplementaryAllows struct {
}

// TODO: Doc formatting needs to be implemented!
type TSignalsSemanticsSupplementaryDistance struct {
}

// TODO: Doc formatting needs to be implemented!
type TSignalsSemanticsSupplementaryEnvironment struct {
}

// TODO: Doc formatting needs to be implemented!
type TSignalsSemanticsSupplementaryExplanatory struct {
}

// TODO: Doc formatting needs to be implemented!
type TSignalsSemanticsSupplementaryProhibits struct {
}

// TODO: Doc formatting needs to be implemented!
type TSignalsSemanticsSupplementaryTime struct {
}

// TODO: Doc formatting needs to be implemented!
type TSignalsSemanticsTourist struct {
}

// TODO: Doc formatting needs to be implemented!
type TSignalsSemanticsWarning struct {
}
