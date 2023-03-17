package main

import (
	"myapp/game"
)

var chan1 chan string
var chan2 chan string

func main() {
	displayChan := make(chan string)
	roundChan := make(chan int)

	game := game.Game{
		DisplayChan: displayChan,
		RoundChan:   roundChan,
		Round: game.Round{
			RoundNumber:   0,
			PlayerScore:   0,
			ComputerScore: 0,
		},
	}

	go game.Rounds()

	game.ClearScreen()

	game.PrintIntro()

	for {
		game.RoundChan <- 1
		<-game.RoundChan // Wait for response from channel

		if game.Round.RoundNumber > 3 {
			break
		}

		if !game.PlayRound() {
			game.RoundChan <- -1
			<-game.RoundChan
		}
	}
	//fmt.Println("Game over!")
	//fmt.Println("Here is the result:")
	//if playerScore > computerScore {
	//	fmt.Println("Player wins!")
	//} else {
	//	fmt.Println("Computer wins!")
	//}
	//fmt.Printf("Computer wins %d/3\n", computerScore)
	//fmt.Printf("Player wins %d/3\n", playerScore)
}
