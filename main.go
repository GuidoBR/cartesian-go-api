package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := NewRouter()

	fmt.Println("Listening to :8000\n")
	log.Fatal(http.ListenAndServe(":8000", router))
}
