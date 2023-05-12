package main

import (
	"fmt"
	"net/http"
)

func myHome(w http.ResponseWriter, r *http.Request) {
	w.Write("Hello World")
}
func main() {
	http.HandleFunc("/", myHome)
	fmt.Println("Starting the server....")
}
