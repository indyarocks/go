package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

func main() {
	playerChoice := ""
	//playerValue := -1

	computerChoice := rand.Intn(3)

	clearScreen()

	fmt.Print("Please enter your choice -> ")
	reader := bufio.NewReader(os.Stdin)
	playerChoice, _ = reader.ReadString('\n')
	playerChoice = strings.Replace(playerChoice, "\r\n", "", -1)
	playerChoice = strings.Replace(playerChoice, "\n", "", -1)
	lowerPlayerChoice := strings.ToLower(playerChoice)

	fmt.Println()
	if lowerPlayerChoice == "rock" {
		fmt.Println("You choose: ", playerChoice, "Your value: ", ROCK)
		fmt.Println("Computer choose:", computerChoice)
	} else if lowerPlayerChoice == "paper" {
		fmt.Println("You choose: ", playerChoice, "Your value: ", PAPER)
		fmt.Println("Computer choose:", computerChoice)
	} else if lowerPlayerChoice == "scissors" {
		fmt.Println("You choose: ", playerChoice, "Your value: ", SCISSORS)
		fmt.Println("Computer choose:", computerChoice)
	}

}

func clearScreen() {
	if strings.Contains(runtime.GOOS, "window") {
		//windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		//	linux or mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

}
