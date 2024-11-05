package main

import (
	"fmt"
	"math/rand"
	"os/signal"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	log.Info().Msg("Hello Kubernetes!")
	log.Info().Msg("Guess the number from 0 to 5. Enter -1 to exit.")

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	go func() {
		<-signalChan
		log.Info().Msg("Received an interrupt, stopping services...")
		os.Exit(0)
	}()

	for {
		randomNumber := r.Intn(6)

		var userGuess int
		fmt.Print("Enter your guess: ")
		_, err := fmt.Scan(&userGuess)
		if err != nil {
			log.Error().Err(err).Msg("Failed to read input")
			continue
		}

		if userGuess == -1 {
			log.Info().Msg("Exiting the game. Goodbye!")
			fmt.Println("Exiting the game. Goodbye!")
			break
		}

		log.Info().Int("Random Number", randomNumber).Msg("Generated number")

		if userGuess == randomNumber {
			log.Info().Msg("You WIN!")
			fmt.Println("You WIN!")
		} else {
			log.Info().Msg("You LOSE!")
			fmt.Println("You LOSE!")
		}
	}
}
