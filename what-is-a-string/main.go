package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "Hello Chandan"

	for i := 0; i < len(name); i++ {
		fmt.Printf("%x ", name[i])
	}
	fmt.Println()
	fmt.Println("Index\tRune\tString")
	for x, y := range name {
		fmt.Println(x, "\t", y, "\t", string(y))
	}

	fmt.Println(name[6:13])

	courses := []string{
		"Learn Go for beginners crash course",
		"Learn Ruby for beginners crash course",
		"Learn c++ for beginners crash course",
		"Learn Python for beginners crash course",
		"Learn C for beginners crash course",
	}

	for _, x := range courses {
		if strings.Contains(x, "Go") {
			fmt.Println("Go is found in", x, "and index is", strings.Index(x, "Go"))
		}
	}

	array := [3]int{1, 2, 3}
	for _, x := range array {
		fmt.Println(x)
	}

	newString := "Go is a good programming language. Go for it"
	fmt.Println("Prefix Go", strings.HasPrefix(newString, "Go"))
	fmt.Println("Prefix Python", strings.HasPrefix(newString, "Python"))
	fmt.Println("Suffix it", strings.HasSuffix(newString, "it"))
	fmt.Println("Suffix Go", strings.HasSuffix(newString, "Go"))
	fmt.Println("Count of Go", strings.Count(newString, "Go"))
	fmt.Println("Count of Fish", strings.Count(newString, "Fish"))
	fmt.Println("Index of Go", strings.Index(newString, "Go"))
	fmt.Println("Index of Python", strings.Index(newString, "Python"))
	fmt.Println("Index of last Go", strings.LastIndex(newString, "Go"))
}
