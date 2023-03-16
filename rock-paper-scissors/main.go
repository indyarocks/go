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

var reader *bufio.Reader

func getUserInput() string {
	reader = bufio.NewReader(os.Stdin)
	playerChoice, _ := reader.ReadString('\n')
	playerChoice = strings.Replace(playerChoice, "\r\n", "", -1)
	playerChoice = strings.Replace(playerChoice, "\n", "", -1)
	return strings.ToLower(playerChoice)
}

func getPlayerValue(playerChoice string) int {
	playerValue := -1
	if playerChoice == "rock" {
		playerValue = ROCK
	} else if playerChoice == "paper" {
		playerValue = PAPER
	} else if playerChoice == "scissors" {
		playerValue = SCISSORS
	}
	return playerValue
}

func main() {
	playerChoice := ""
	playerValue := -1
	playerWins := 0
	computerWins := 0

	clearScreen()
	fmt.Println("Let the game begin. Best of 3 wins!")
	fmt.Println("===================================")
	for round := 1; round <= 3; round++ {
		fmt.Println("Round:", round)
		fmt.Println("---------------------------------")
		fmt.Print("Please enter your choice -> ")
		playerChoice = getUserInput()
		fmt.Println()
		playerValue = getPlayerValue(playerChoice)
		computerValue := rand.Intn(3)
		switch computerValue {
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
		fmt.Println("You choose: ", strings.ToUpper(playerChoice))

		if playerValue == computerValue {
			fmt.Println("It's a draw!")
			round--
		} else {
			switch playerValue {
			case ROCK:
				if computerValue == PAPER {
					computerWins += 1
					fmt.Println("Computer wins!")
				} else {
					playerWins += 1
					fmt.Println("Player wins!")
				}
				break
			case PAPER:
				if computerValue == SCISSORS {
					computerWins += 1
					fmt.Println("Computer wins!")
				} else {
					playerWins += 1
					fmt.Println("Player wins!")
				}
				break
			case SCISSORS:
				if computerValue == ROCK {
					computerWins += 1
					fmt.Println("Computer wins!")
				} else {
					playerWins += 1
					fmt.Println("Player wins!")
				}
				break
			default:
				round--
				fmt.Println("Invalid choice")
			}
		}
		fmt.Println("---------------------------------")
		fmt.Println()
	}
	fmt.Println("Game over!")
	fmt.Println("Here is the result:")
	fmt.Printf("Computer wins %d/3\n", computerWins)
	fmt.Printf("Player wins %d/3\n", playerWins)
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
