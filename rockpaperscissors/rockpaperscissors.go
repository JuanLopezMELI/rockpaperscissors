package rockpaperscissors

import (
	"math/rand"
	"strconv"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

type Round struct {
	Message           string `json:"message"`
	ComputerChoice    string `json:"computer_choice"`
	RoundResult       string `json:"round_result"`
	ComputerChoiceInt int    `json:"computer_choice_int"`
	ComputerScore     string `json:"computer_score"`
	PlayerScore       string `json:"player_score"`
}

var victoryMessages = []string{
	"Good job! You won the round!",
	"Well done! You won the round!",
	"You should buy a lottery ticket! You won the round!",
}

var defeatMessages = []string{
	"Sorry, you lost the round.",
	"Better luck next time!",
	"Too bad, you lost the round.",
}

var drawMessages = []string{
	"It's a draw!",
	"Nobody wins this round.",
	"Nobody loses this round.",
}

var (
	PlayerScore   int
	ComputerScore int
)

func PlayRound(playerValue int) Round {
	computerValue := rand.Intn(3)

	var computerChoice, roundResult string
	var computerChoiceInt int

	switch computerValue {
	case ROCK:
		computerChoice = "Computer chose rock"
		computerChoiceInt = ROCK
	case PAPER:
		computerChoice = "Computer chose paper"
		computerChoiceInt = PAPER
	case SCISSORS:
		computerChoice = "Computer chose scissors"
		computerChoiceInt = SCISSORS
	}

	messageInt := rand.Intn(3)

	var message string

	if playerValue == computerValue {
		roundResult = "It's a draw!"
		message = drawMessages[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		PlayerScore++
		roundResult = "You won the round!"
		message = victoryMessages[messageInt]
	} else {
		ComputerScore++
		roundResult = "You lost the round."
		message = defeatMessages[messageInt]
	}

	return Round{
		Message:           message,
		ComputerChoice:    computerChoice,
		RoundResult:       roundResult,
		ComputerChoiceInt: computerChoiceInt,
		ComputerScore:     strconv.Itoa(ComputerScore),
		PlayerScore:       strconv.Itoa(PlayerScore),
	}
}
