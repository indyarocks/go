package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Println("Calling dup1()....")
	dup1()
	fmt.Println("Ending dup1()....")
	fmt.Println("Args[1:]", os.Args[1:])
	for _, arg := range os.Args[1:] {
		fmt.Printf("Argument: %s | %q\t Type: %T\n", arg, arg, arg)
	}
	fmt.Println("Calling dup2()....")
	dup2(os.Args[1:])
	fmt.Println("Ending dup2()....")
	fmt.Println("Calling dup3()....")
	dup3(os.Args[1:])
	fmt.Println("Ending dup3()....")
}
func dup3(args []string) {
	counts := make(map[string]int)
	for _, fileName := range args {
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, count := range counts {
		fmt.Printf("%d \t %s\n", count, line)
	}
}

func dup2(args []string) {
	fmt.Println(args)
	counts := make(map[string]int)
	if len(args) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, fileName := range args {
			file, err := os.Open(fileName)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dups: %v\n", err)
				continue
			}
			countLines(file, counts)
			file.Close()
		}
		for line, count := range counts {
			fmt.Printf("%d \t %s\n", count, line)
		}
	}

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

func countLines(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		counts[input.Text()]++
	}
}
