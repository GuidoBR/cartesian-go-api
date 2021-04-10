package main

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type PointWithDistance struct {
	p Point
	d int
}

type Points []Point
type PointsWithDistance []PointWithDistance
