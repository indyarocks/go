package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")

	var mySlice = make([]string, 3)
	fmt.Println("mySlice", mySlice)

	mySlice[0] = "Chandan"
	mySlice[1] = "Alok"
	mySlice[2] = "Ruby"
	fmt.Println("mySlice", mySlice)

	newSlice := make([]string, len(mySlice))
	fmt.Println("newSlice", newSlice)
	copy(newSlice, mySlice)
	fmt.Println("newSlice", newSlice)

}
