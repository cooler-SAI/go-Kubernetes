package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	log.Info().Msg("Hello Kubernetes! Starting automated guessing game.")

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	go func() {
		<-signalChan
		log.Info().Msg("Received an interrupt, stopping services...")
		os.Exit(0)
	}()

	go func() {
		for {
			randomNumber := r.Intn(6)
			userGuess := r.Intn(6)

			log.Info().Int("User Guess", userGuess).Msg("Generated guess")
			log.Info().Int("Random Number", randomNumber).Msg("Generated number")

			if userGuess == randomNumber {
				log.Info().Msg("You WIN!")
			} else {
				log.Info().Msg("You LOSE!")
			}

			time.Sleep(5 * time.Second)
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, "Welcome to the automated guessing game!")
		if err != nil {
			return
		}
	})

	log.Info().Msg("Starting HTTP server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
