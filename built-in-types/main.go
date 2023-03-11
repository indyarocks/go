package main

import "fmt"

type Car struct {
	NumberOfTires int
	Luxury        bool
	BucketSeats   bool
	Make          string
	Model         string
	Year          int
}

func main() {
	var myString [3]string

	myString[0] = "cat"
	myString[1] = "dog"
	myString[2] = "fish"

	fmt.Println("First element in array is", myString[0])

	var myCar Car

	myCar.NumberOfTires = 4
	myCar.Luxury = false
	myCar.Make = "Volksvagan"
}
