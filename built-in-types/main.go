package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader

func main() {
	var _ map[string]string
	intMap := make(map[string]int)

	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5

	for key, val := range intMap {
		fmt.Println(key, val)
	}

	println("Deleting four")
	delete(intMap, "four")
	for key, val := range intMap {
		fmt.Println(key, val)
	}

	el, ok := intMap["four"]
	if ok {
		fmt.Println(el, "in in map.")
	} else {
		fmt.Println(el, "is not in map")
	}

	reader = bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Do you want to check for a key in map(y/n)?")
		userInput := readAndTrimString()
		if strings.ToLower(userInput) != "y" && strings.ToLower(userInput) != "n" {
			fmt.Println("Please enter y or no")
		} else if strings.ToLower(userInput) == "y" {
			fmt.Println("Which key you're looking for?")
			check_key := readAndTrimString()
			el, ok := intMap[check_key]
			if ok {
				fmt.Println(el, "in in map.")
			} else {
				fmt.Println(check_key, "is not in map")
			}
		} else {
			break
		}
	}

}

func readAndTrimString() string {
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\r\n", "", -1)
	userInput = strings.Replace(userInput, "\n", "", -1)
	return userInput
}
