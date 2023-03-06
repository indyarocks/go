package main

import "fmt"

func main() {
	fmt.Println("Hello-world")
	sayHelloWorld("This is Chandan")
}

func sayHelloWorld(whatToSay string) {
	fmt.Println(whatToSay)
}
