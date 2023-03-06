package main

import "fmt"

var one = "One"

func main() {
	var somethingElse = "This is a block level variable"

	fmt.Println(somethingElse)

	myFunc()

}

func myFunc() {
	fmt.Println(one)
}
