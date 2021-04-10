package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetPoints(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	x, _ := strconv.Atoi(v.Get("x"))
	y, _ := strconv.Atoi(v.Get("y"))
	d, _ := strconv.Atoi(v.Get("distance"))

	// TODO: What if we don't have x, y, d?
	p1 := Point{x, y}

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
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(pd); err != nil {
		panic(err)
	}
}

func ManhattanDistance(p1 Point, p2 Point) int {
	d := abs(p1.X-p2.X) + abs(p1.Y-p2.Y)
	return d
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
