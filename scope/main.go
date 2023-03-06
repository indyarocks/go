package main

import "fmt"

var one = "One"

func main() {
	var one = "New One"

	fmt.Println(one)

	myFunc()

}

func myFunc() {
	fmt.Println(one)
}
