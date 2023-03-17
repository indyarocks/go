package rps

import (
	"math/rand"
)

const (
	ROCK         = 0 // beats scissors. (scissors + 1) % 3 =0
	PAPER        = 1 // beats rock. (rock + 1) % 3 = 1
	SCISSORS     = 2 // beats paper. (paper + 1) % 3 = 2
	PLAYERWINS   = 1
	COMPUTERWINS = 2
	DRAW         = 3
)

type Round struct {
	Winner         int    `json:"winner"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
}

// PlayRound - Arguments - Current Round winner, Computer Choice, Result
func PlayRound(playerValue int) Round {
	computerValue := rand.Intn(3)
	computerChoice := ""
	roundResult := ""
	winner := 0

	switch computerValue {
	case ROCK:
		computerChoice = "Computer chose ROCK"
		break
	case PAPER:
		computerChoice = "Computer chose PAPER"
		break
	case SCISSORS:
		computerChoice = "Computer chose SCISSORS"
		break
	default:
	}

	if playerValue == computerValue {
		roundResult = "It's a Draw!"
		winner = DRAW
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "Player Wins!"
		winner = PLAYERWINS
	} else {
		roundResult = "Computer Wins!"
		winner = COMPUTERWINS
	}

	result := Round{
		Winner:         winner,
		ComputerChoice: computerChoice,
		RoundResult:    roundResult,
	}
	return result
}
