package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	fmt.Println("Listening to :8000\n")
	log.Fatal(http.ListenAndServe(":8000", router))
}
