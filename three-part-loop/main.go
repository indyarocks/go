package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Looping with index", i)
	}

	var x int = 1000

	rand.Seed(time.Now().UnixNano())

	for x > 100 {
		x = rand.Intn(1000) + 1
		fmt.Println("Value of x is:", x, "and loop continues..")
	}

	fmt.Println("Got", x, "and broke us out of loop.")

}
