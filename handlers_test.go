package main

import (
	"testing"
)

func TestManhattanDistanceWithValuesFromPointsJson(t *testing.T) {
	p1 := Point{100, 20}
	d := 35
	expectedPoints := []Point{
		{90, 43},
		{78, 8},
	}

	pd := GetPointsInsideDistance(p1, d)

	if len(pd) != len(expectedPoints) {
		t.Error(`Length of points within distance is larger than the expected`)
	}

	for i := 0; i < len(pd); i++ {
		if (pd[i].X != expectedPoints[i].X) || (pd[i].Y != expectedPoints[i].Y) {
			t.Error(`Points different from expected`)
		}
	}
}

func TestManhattanDistanceWithDifferentValuesFromPointsJson(t *testing.T) {
	p1 := Point{100, 20}
	d := 45
	expectedPoints := []Point{
		{90, 43},
		{78, 8},
		{91, 49},
		{95, -14},
		{64, 17},
		{79, 39},
		{74, 37},
	}

	pd := GetPointsInsideDistance(p1, d)

	if len(pd) != len(expectedPoints) {
		t.Error(`Length of points within distance is larger than the expected`)
	}

	for i := 0; i < len(pd); i++ {
		if (pd[i].X != expectedPoints[i].X) || (pd[i].Y != expectedPoints[i].Y) {
			t.Error(`Points different from expected`)
		}
	}
}
