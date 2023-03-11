package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

var reader *bufio.Reader

type User struct {
	userName       string
	userAge        int
	favoriteNumber float64
	OwnsADog       bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)
	var user User
	user.userName = readString("What is your name?")

	user.userAge = readInt("How old are you?")

	user.favoriteNumber = readFloat("Your favorite numer?")

	user.OwnsADog = readBool("Do you own a dog (y/n)?")

	fmt.Println("Your name is", user.userName, "and you are", user.userAge, "years old.")
	fmt.Println("Your name is "+user.userName+". You are", user.userAge, "years old.")
	fmt.Println(fmt.Sprintf("Your name is %s. You are %d years old.", user.userName, user.userAge))
	fmt.Printf("Your name is %s. You are %d years old. Your favorite number is %.4f. You own a dog: %t", user.userName, user.userAge, user.favoriteNumber, user.OwnsADog)
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

func readBool(s string) bool {
	fmt.Println(s)
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()
	for {
		prompt()
		char, _, err := keyboard.GetSingleKey()

		if err != nil {
			log.Fatal(err)
		} else {
			if char == 'y' || char == 'Y' {
				return true
			} else {
				return false
			}
		}
	}
}
