package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type User struct {
	userName       string
	userAge        int
	favoriteNumber float64
}

func main() {
	reader = bufio.NewReader(os.Stdin)
	var user User
	user.userName = readString("What is your name?")

	user.userAge = readInt("How old are you?")

	user.favoriteNumber = readFloat("Your favorite numer?")

	fmt.Println("Your name is", user.userName, "and you are", user.userAge, "years old.")
	fmt.Println("Your name is "+user.userName+". You are", user.userAge, "years old.")
	fmt.Println(fmt.Sprintf("Your name is %s. You are %d years old.", user.userName, user.userAge))
	fmt.Printf("Your name is %s. You are %d years old. Your favorite number is %.4f\n", user.userName, user.userAge, user.favoriteNumber)
}

func prompt() {
	fmt.Print("-> ")
}

func readString(s string) string {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "" {
			fmt.Println("Please enter your name")
		} else {
			return userInput
		}
	}
}

func readInt(s string) int {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Please enter a whole number")
		} else {
			return num
		}
	}
}

func readFloat(s string) float64 {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		num, err := strconv.ParseFloat(userInput, 64)
		if err != nil {
			fmt.Println("Please enter valid number")
		} else {
			return num
		}
	}
}
