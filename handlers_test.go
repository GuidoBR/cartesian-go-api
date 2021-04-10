package main

import "testing"

func TestManhattanDistance(t *testing.T) {
	p1 := Point{10, 20}
	p2 := Point{30, 20}

	if ManhattanDistance(p1, p2) > 10 {
		t.Error(`ManhattanDistance(p1{10, 20}, p2{30, 20} > 10`)
	}
}
