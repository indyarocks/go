package main

import "log"

var myInt int
var myUint uint

// var myInt16 int16
// var myInt32 int32
// var myInt64 int64

var myFloat float32
var myFloat64 float64

func main() {
	myInt = 10
	myUint = 100

	myFloat = 10.1
	myFloat64 = 100.1

	log.Println(myInt, myUint, myFloat, myFloat64)

	myString := "Trevor"
	log.Println(myString)

	myString = "Chandan"
	log.Println(myString)

	var myBool = true
	myBool = false

	log.Println(myBool)
}
