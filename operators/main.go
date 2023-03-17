package main

import (
	"fmt"
	"math"
)

func main() {
	answer := 7 + 3*4
	fmt.Println("Answer", answer)

	answer = (7 + 3) * 4
	fmt.Println("Answer is now", answer)

	var radius = 12.0
	area := math.Pi * radius * radius
	fmt.Println("Area", area)

	goodThreeSquared := math.Pow(3.0, 2.0)
	fmt.Println("goodThreeSquared", goodThreeSquared)
}
