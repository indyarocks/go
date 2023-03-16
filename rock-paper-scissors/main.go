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
	playerValue := -1

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
		playerValue = ROCK
	} else if lowerPlayerChoice == "paper" {
		playerValue = PAPER
	} else if lowerPlayerChoice == "scissors" {
		playerValue = SCISSORS
	}

	switch computerChoice {
	case ROCK:
		fmt.Println("Computer choose: ROCK")
		break
	case PAPER:
		fmt.Println("Computer choose: PAPER")
		break
	case SCISSORS:
		fmt.Println("Computer choose: SCISSORS")
		break
	default:
	}
	fmt.Println("You choose: ", playerChoice, "Your value: ", playerValue)

	if playerValue == computerChoice {
		fmt.Println("It's a draw!")
	} else {
		switch playerValue {
		case ROCK:
			if computerChoice == PAPER {
				fmt.Println("Computer wins!")
			} else {
				fmt.Println("Player wins!")
			}
			break
		case PAPER:
			if computerChoice == SCISSORS {
				fmt.Println("Computer wins!")
			} else {
				fmt.Println("Player wins!")
			}
			break
		case SCISSORS:
			if computerChoice == ROCK {
				fmt.Println("Computer wins!")
			} else {
				fmt.Println("Player wins!")
			}
			break
		default:
			fmt.Println("Invalid choice")
		}
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
