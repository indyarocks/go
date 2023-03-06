package main

import "fmt"

var one = "One"

func main() {
	var one = "New One"

	fmt.Println(one)

	myFunc()

}

func myFunc() {
	var one = "The number one"
	fmt.Println(one)
}
