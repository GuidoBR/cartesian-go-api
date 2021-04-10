package main

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestManhattanDistance(t *testing.T) {
	p1 := Point{100, 20}
	p2 := Point{95, -14}
	p3 := Point{83, -20}
	p4 := Point{92, -32}
	p5 := Point{30, 20}

	if ManhattanDistance(p1, p2) > 200 {
		t.Error(`TestManhattanDistance(p1{100, 20}, p2{30, 20} > 200`)
	}
	if ManhattanDistance(p1, p3) > 200 {
		t.Error(`TestManhattanDistance(p1{100, 20}, p3{95, -14} > 200`)
	}
	if ManhattanDistance(p1, p4) > 200 {
		t.Error(`TestManhattanDistance(p1{100, 20}, p4{83, -20} > 200`)
	}
	if ManhattanDistance(p1, p5) > 200 {
		t.Error(`TestManhattanDistance(p1{100, 20}, p5{30, 20} > 200`)
	}
}

func TestManhattanDistanceWithValuesFromPointsJson(t *testing.T) {
	p1 := Point{100, 20}
	d := 200

	f, _ := ioutil.ReadFile("data/points.json")
	points := Points{}
	_ = json.Unmarshal([]byte(f), &points)

	pd := Points{}
	for i := 0; i < len(points); i++ {
		p2 := Point{points[i].X, points[i].Y}
		md := ManhattanDistance(p1, p2)

		if md <= d {
			pd = append(pd, p2)
		}

		if ManhattanDistance(p1, p2) > d {
			t.Error(`TestManhattanDistance(p1{100, 20}, p5{30, 20} > 200`)
		}
	}
}
