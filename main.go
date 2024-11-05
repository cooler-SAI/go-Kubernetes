package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"math/rand"
	"os"
	"time"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	log.Info().Msg("Hello Kubernetes!")
	log.Info().Msg("Please Enter the value from 0 to 5...")

	var userGuess int
	log.Info().Msg("Enter your guess:")
	_, err := fmt.Scan(&userGuess)
	if err != nil {
		log.Error().Err(err).Msg("Failed to read user guess")
		return
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := r.Intn(6)

	log.Info().Int("RandomNumber", randomNumber).Msg("Generated number")

	if userGuess == randomNumber {
		log.Info().Msg("You WIN!")

	} else {
		log.Info().Msg("You LOSE!")
	}

}
