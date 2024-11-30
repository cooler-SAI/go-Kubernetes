package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/rs/zerolog/log"
	"go-Kubernetes/game"
)

// HandleRoot processes the root endpoint
func HandleRoot(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintln(w, "Welcome to the automated guessing game!")
	if err != nil {
		log.Error().Err(err).Msg("Error writing response")
	}
}

// HandleGuess processes the /guess endpoint
func HandleGuess(w http.ResponseWriter, r *http.Request) {
	userGuessStr := r.URL.Query().Get("guess")
	if userGuessStr == "" {
		http.Error(w, "Missing 'guess' query parameter", http.StatusBadRequest)
		return
	}

	userGuess, err := strconv.Atoi(userGuessStr)
	if err != nil {
		http.Error(w, "Invalid 'guess' parameter. Must be an integer.", http.StatusBadRequest)
		return
	}

	randomNumber := game.GenerateRandomNumber()
	log.Info().Int("User Guess", userGuess).Msg("Received user guess")
	log.Info().Int("Random Number", randomNumber).Msg("Generated random number")

	if userGuess == randomNumber {
		_, err := fmt.Fprintln(w, "You WIN!")
		if err != nil {
			log.Error().Err(err).Msg("Error writing response")
		}
	} else {
		_, err := fmt.Fprintln(w, "You LOSE!")
		if err != nil {
			log.Error().Err(err).Msg("Error writing response")
		}
	}
}
