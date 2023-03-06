package main

import "fmt"

func main() {
	fmt.Println("Hello-world")
	var whatToSay string
	whatToSay = "This is a variable"
	sayHelloWorld(whatToSay)
}

func sayHelloWorld(whatToSay string) {
	fmt.Println(whatToSay)
}
