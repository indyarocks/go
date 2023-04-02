package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Calling dup1()....")
	dup1()
	fmt.Println("Ending dup1()....")
	for _, arg := range os.Args[1:] {
		fmt.Printf("Argument: %s | %q\t Type: %T\n", arg, arg, arg)
	}
	fmt.Println("Calling dup2()....")
	dup2()
	fmt.Println("Ending dup2()....")
}

func dup2() {

}

func dup1() {
	counts := make(map[string]int)
	counts["hi"] = 1
	counts["bye"] = 2
	//input := bufio.NewScanner(os.Stdin)
	//
	//for input.Scan() {
	//	counts[input.Text()]++
	//}

	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\t%T\t%T\n", count, line, count, line)
		}
	}
}
