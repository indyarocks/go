package main

import (
	"bufio"
	"fmt"
	"myapp/mylogger"
	"os"
	"strings"
	"time"
)

var reader *bufio.Reader

func main() {
	// read input from user 5 times and write it to log

	ch := make(chan string)

	go mylogger.ListenForLog(ch)

	fmt.Println("Going inside the loop..")

	for i := 0; i <= 4; i++ {
		Prompt()
		ch <- ReadUserString()
		time.Sleep(time.Second)
	}
}

func Prompt() {
	fmt.Println("->")
}

func ReadUserString() string {
	reader = bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\r\n", "", -1)
	userInput = strings.Replace(userInput, "\n", "", -1)
	return userInput
}
