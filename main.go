package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := NewRouter()

	fmt.Println("Listening to :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
