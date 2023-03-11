package main

import "fmt"

func main() {
	z := addTwoNumbers(2, 4)
	fmt.Println(z)
}

func addTwoNumbers(x, y int) (sum int) {
	sum = x + y
	return
}
