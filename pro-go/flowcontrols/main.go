package main

import "fmt"

func main() {
	name := "ChAndaN"

	for i, character := range name {
		switch character {
		case 'A', 'a':
			fmt.Println("a at index", i)
			//fallthrough // causes only next statement to go through
		case 'n':
			fmt.Println("n at index", i)
		case 'N':
			fmt.Println("N at index", i)
		default:
			fmt.Println(string(character), "at index", i)
		}
	}
}
