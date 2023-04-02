package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	//	Read arguments os.Args[1:]
	for _, url := range os.Args[1:] {
		if !(strings.HasPrefix(url, "http://")) {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		//body, err := ioutil.ReadAll(resp.Body)
		//if err != nil {
		//	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		//	os.Exit(1)
		//}
		//resp.Body.Close()
		//fmt.Printf("%s", body)
	}
}
