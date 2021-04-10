package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/points", GetPoints).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetPoints(w http.ResponseWriter, r *http.Request) {}
