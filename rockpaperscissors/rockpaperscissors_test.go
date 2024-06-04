package rockpaperscissors

import (
	"fmt"
	"testing"
)

func TestPlayRound(t *testing.T) {
	for i := 0; i < 3; i++ {
		round := PlayRound(i)

		switch i {
		case ROCK:
			fmt.Println("Player chose rock")
		case PAPER:
			fmt.Println("Player chose paper")
		case SCISSORS:
			fmt.Println("Player chose scissors")
		}

		fmt.Println("Message:", round.Message)
		fmt.Println("Computer Choice:", round.ComputerChoice)
		fmt.Println("Computer ChoiceInt:", round.ComputerChoiceInt)
		fmt.Println("Round Result:", round.RoundResult)
		fmt.Println("Player score:", round.PlayerScore)
		fmt.Println("Computer score:", round.ComputerScore)
		fmt.Println()
	}
}
