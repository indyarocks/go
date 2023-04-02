package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var count int
var mu sync.Mutex

func homePage(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL Path: %q", r.URL)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Request Count: %d\n", count)
	mu.Unlock()
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
