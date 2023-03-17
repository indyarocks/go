package main

import (
	"errors"
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
	x := 10
	y := 0

	c, err := divideTwoNumbers(x, y)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(c)
	}
}

func divideTwoNumbers(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("can't divide by 0")
	} else {
		return x / y, nil
	}
}
