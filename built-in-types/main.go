package main

import "fmt"

func main() {
	go doSomethingInRoutine("I'm inside do something routine")

	fmt.Println("This is first another message")

	go doSomethingInRoutine("I'm inside another something routine")

	fmt.Println("This is second another message")

	for {
		// Do nothing
	}
}

func doSomethingInRoutine(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}
