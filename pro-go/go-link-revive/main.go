// main Checking revive go linter
package main

import "fmt"

func main() {
	PrintHello()
	for i := 0; i < 10; i++ {
		PrintNum(i)
	}
}

func PrintHello() {
	fmt.Println("Hello")
}

func PrintNum(num int) {
	fmt.Println(num)
}
