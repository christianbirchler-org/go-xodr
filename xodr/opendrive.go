package xodr

import (
	"encoding/xml"
)

type OpenDrive struct {
	XMLName xml.Name `xml:"OpenDRIVE"`
	Header  Header   `xml:"header"`
	Road    Road     `xml:"road"`
}

type Header struct {
	XMLName     xml.Name    `xml:"header"`
	Name        string      `xml:"name,attr"`
	SdcTestInfo SdcTestInfo `xml:"sdc_test_info"`
}

type SdcTestInfo struct {
	XMLName xml.Name `xml:"sdc_test_info"`
}

type Road struct {
	XMLName  xml.Name `xml:"road"`
	PlanView PlanView `xml:"planView"`
}

type PlanView struct {
	XMLName    xml.Name   `xml:"planView"`
	Geometries []Geometry `xml:"geometry"`
}

type Geometry struct {
	XMLName    xml.Name   `xml:"geometry"`
	S          float64    `xml:"s,attr"`
	X          float64    `xml:"x,attr"`
	Y          float64    `xml:"y,attr"`
	Hdg        float64    `xml:"hdg,attr"`
	Length     float64    `xml:"length,attr"`
	ParamPoly3 ParamPoly3 `xml:"paramPoly3"`
}

type ParamPoly3 struct {
	PRange string  `xml:"pRange,attr"`
	AU     float64 `xml:"aU"`
	BU     float64 `xml:"bU"`
	CU     float64 `xml:"cU"`
	DU     float64 `xml:"dU"`
	AV     float64 `xml:"aV"`
	BV     float64 `xml:"bV"`
	CV     float64 `xml:"cV"`
	DV     float64 `xml:"dV"`
}

type RoadPoint struct {
	X float64
	Y float64
	Z float64
}

func (o *OpenDrive) ToPoints() ([]RoadPoint, error) {
	numberRoadPoints := len(o.Road.PlanView.Geometries)
	roadPoints := make([]RoadPoint, numberRoadPoints)
	for i, geo := range o.Road.PlanView.Geometries {
		roadPoints[i].X = geo.X
		roadPoints[i].Y = geo.Y
	}
	return roadPoints, nil
}
