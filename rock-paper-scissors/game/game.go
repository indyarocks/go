package game

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

var reader = bufio.NewReader(os.Stdin)

type Game struct {
	DisplayChan chan string
	RoundChan   chan int
	Round       Round
}

type Round struct {
	RoundNumber   int
	PlayerScore   int
	ComputerScore int
}

func (g *Game) Rounds() {
	//	use select to process input in channels
	//	  Print information to screen
	//	Keep track of round number
	for {
		select {
		case round := <-g.RoundChan:
			g.Round.RoundNumber = g.Round.RoundNumber + round
			g.RoundChan <- 1
		case msg := <-g.DisplayChan:
			fmt.Println(msg)
		}
	}
}

func (g *Game) ClearScreen() {
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

func (g *Game) PrintIntro() {
	g.DisplayChan <- fmt.Sprintf(`
Rock, Paper & Scissors
======================
Let the game begin. Best of 3 wins!
===================================
`)
}

func (g *Game) PlayRound() bool {
	playerValue := -1

	fmt.Println()
	fmt.Println("Round", g.Round.RoundNumber)
	fmt.Println("-----")

	fmt.Print("Please enter rock, paper or scissors -> ")
	playerChoice, _ := reader.ReadString('\n')
	playerChoice = strings.Replace(playerChoice, "\n", "", -1)
	computerValue := rand.Intn(3)

	if playerChoice == "rock" {
		playerValue = ROCK
	} else if playerChoice == "paper" {
		playerValue = PAPER
	} else if playerChoice == "scissors" {
		playerValue = SCISSORS
	}

	fmt.Println()
	g.DisplayChan <- fmt.Sprintf("Player chose %s", strings.ToUpper(playerChoice))
	//<-g.DisplayChan

	switch computerValue {
	case ROCK:
		g.DisplayChan <- "Computer choose: ROCK"
		break
	case PAPER:
		g.DisplayChan <- "Computer choose: PAPER"
		break
	case SCISSORS:
		g.DisplayChan <- "Computer choose: SCISSORS"
		break
	default:
	}

	fmt.Println()

	if playerValue == computerValue {
		g.DisplayChan <- "It's a draw!"
		return false
	} else {
		switch playerValue {
		case ROCK:
			if computerValue == PAPER {
				g.computerWins()
			} else {
				g.playerWins()
			}
			break
		case PAPER:
			if computerValue == SCISSORS {
				g.computerWins()
			} else {
				g.playerWins()
			}
			break
		case SCISSORS:
			if computerValue == ROCK {
				g.computerWins()
			} else {
				g.playerWins()
			}
			break
		default:
			g.DisplayChan <- "Invalid choice"
			return false
		}
	}
	return true
}

func (g *Game) computerWins() {
	g.Round.ComputerScore++
	g.DisplayChan <- "Computer wins!"
}

func (g *Game) playerWins() {
	g.Round.PlayerScore++
	g.DisplayChan <- "Player wins!"
}

func (g *Game) PrintSummary() {
	g.DisplayChan <- "Game over!"
	g.DisplayChan <- "Here is the result:"
	if g.Round.PlayerScore > g.Round.ComputerScore {
		g.DisplayChan <- "Player wins!"
	} else {
		g.DisplayChan <- "Computer wins!"
	}
	g.DisplayChan <- fmt.Sprintf("Computer wins %d/3\n", g.Round.ComputerScore)
	g.DisplayChan <- fmt.Sprintf("Player wins %d/3\n", g.Round.PlayerScore)
}
