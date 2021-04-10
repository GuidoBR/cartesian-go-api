package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
)

func GetPoints(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	x, _ := strconv.Atoi(v.Get("x"))
	y, _ := strconv.Atoi(v.Get("y"))
	d, _ := strconv.Atoi(v.Get("distance"))

	p1 := Point{x, y}

	pd := GetPointsInsideDistance(p1, d)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(pd); err != nil {
		panic(err)
	}
}

func GetPointsInsideDistance(p1 Point, d int) Points {
	points := readPointsFromJSON("data/points.json")
	pd := pointsInsideDistance(points, p1, d)

	return pd
}

func readPointsFromJSON(p string) Points {
	f, _ := ioutil.ReadFile(p)
	points := Points{}
	_ = json.Unmarshal([]byte(f), &points)

	return points
}

func pointsInsideDistance(points Points, p1 Point, d int) Points {
	pd := PointsWithDistance{}

	// Filters points based on given distance d
	for i := 0; i < len(points); i++ {
		p2 := Point{points[i].X, points[i].Y}
		md := manhattanDistance(p1, p2)

		if md <= d {
			pd = append(pd, PointWithDistance{p2, md})
		}
	}

	// Sorts array of points based on their distance
	sort.Slice(pd, func(i, j int) bool {
		return pd[i].d < pd[j].d
	})

	// Remove distance from return array
	p := Points{}
	for i := 0; i < len(pd); i++ {
		p = append(p, Point{pd[i].p.X, pd[i].p.Y})
	}

	return p
}

func manhattanDistance(p1 Point, p2 Point) int {
	return abs(p1.X-p2.X) + abs(p1.Y-p2.Y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
