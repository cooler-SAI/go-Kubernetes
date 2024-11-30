package game

import (
	"math/rand"
	"time"

	"github.com/rs/zerolog/log"
)

var randomGenerator *rand.Rand

// InitializeRandomGenerator initializes the global random number generator
func InitializeRandomGenerator() {
	randomGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// GenerateRandomNumber generates a random number between 0 and 5
func GenerateRandomNumber() int {
	return randomGenerator.Intn(6)
}

// StartGuessingGame runs the background guessing game
func StartGuessingGame() {
	for {
		randomNumber := GenerateRandomNumber()
		userGuess := GenerateRandomNumber()

		log.Info().Int("User Guess", userGuess).Msg("Generated guess")
		log.Info().Int("Random Number", randomNumber).Msg("Generated number")

		if userGuess == randomNumber {
			log.Info().Msg("You WIN!")
		} else {
			log.Info().Msg("You LOSE!")
		}

		time.Sleep(5 * time.Second)
	}
}
