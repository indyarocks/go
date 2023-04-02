package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL Path: %q", r.URL)
}

func main() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
