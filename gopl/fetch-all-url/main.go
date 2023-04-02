package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	startTime := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Println("Total time taken:", time.Since(startTime))
}

func fetch(url string, ch chan<- string) {
	startTime := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprintf("Error while fetching URL: %s \n %v\n", url, err)
		return
	}
	nbytes, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("Error while fetching URL: %s \n %v\n", url, err)
		return
	}
	ch <- fmt.Sprintf("Url: %s \t Respons Size: %d \t Time Taken: %2f", url, nbytes, time.Since(startTime).Seconds())
}
