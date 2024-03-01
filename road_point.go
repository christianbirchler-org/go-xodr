package xodr

import "fmt"

type RoadPoint struct {
}

func ToPoints(od *OpenDrive) ([]RoadPoint, error) {
	for _, geo := range od.Road.PlanView.Geometries {
		fmt.Println(geo.X)
	}
	return []RoadPoint{}, nil
}
