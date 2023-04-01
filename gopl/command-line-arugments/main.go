// Echo1 prints it's command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s, sep string
	var s1, sep1 string
	arguments := os.Args[1:]
	fmt.Println(arguments)

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	for i, arg := range os.Args[1:] {
		s1 += sep1 + arg
		sep1 = " "
		fmt.Println(i, arg)
	}
	fmt.Println(s1)

	fmt.Println(strings.Join(os.Args[1:], " "))

	fmt.Println(os.Args)
}
