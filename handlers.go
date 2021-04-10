package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetPoints(w http.ResponseWriter, r *http.Request) {
	f, _ := ioutil.ReadFile("data/points.json")
	points := Points{}
	_ = json.Unmarshal([]byte(f), &points)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(points); err != nil {
		panic(err)
	}
}

func GetIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Points API!\n")
}
