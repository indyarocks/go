package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var reader *bufio.Reader

func main() {
	// read input from user 5 times and write it to log

	for i := 0; i <= 4; i++ {
		Prompt()
		input := ReadUserString("Please enter something")
		log.Println(input)
	}
}

func Prompt() {
	fmt.Println("->")
}

func ReadUserString(s string) string {
	fmt.Println(s)
	reader = bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\r\n", "", -1)
	userInput = strings.Replace(userInput, "\n", "", -1)
	return userInput
}
