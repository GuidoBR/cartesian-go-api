package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetPoints(w http.ResponseWriter, r *http.Request) {
	points := Points{
		Point{X: 1, Y: 1},
		Point{X: 1, Y: 2},
		Point{X: 1, Y: 3},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(points); err != nil {
		panic(err)
	}
}

func GetIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Points API!\n")
}
