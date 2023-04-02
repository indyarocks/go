package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"webServer/lisssajous"
)

var count int
var mu sync.Mutex

func homePage(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	requestDetail(w, r)
}
func requestDetail(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL Path: %q\n", r.URL)
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Request Count: %d\n", count)
	mu.Unlock()
	requestDetail(w, r)
}

func lissajous(w http.ResponseWriter, r *http.Request) {
	lisssajous.Lissajous(w)
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/lissajous", lissajous)
	fmt.Println("Starting server on port 8080....")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
