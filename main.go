package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Global random number generator
var randomGenerator *rand.Rand

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	log.Info().Msg("Hello Kubernetes! Starting automated guessing game.")

	// Initialize the global random number generator
	randomGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))

	// Channel to handle OS interrupt signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	// Gracefully handle OS interrupt signals
	go func() {
		<-signalChan
		log.Info().Msg("Received an interrupt, stopping services...")
		os.Exit(0)
	}()

	// Start the background guessing game
	go startGuessingGame()

	// HTTP routes
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/guess", handleGuess)

	// Start the HTTP server
	log.Info().Msg("Starting HTTP server on port 8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}

// Background guessing game logic
func startGuessingGame() {
	for {
		randomNumber := randomGenerator.Intn(6)
		userGuess := randomGenerator.Intn(6)

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

// Root HTTP handler
func handleRoot(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintln(w, "Welcome to the automated guessing game!")
	if err != nil {
		log.Error().Err(err).Msg("Error writing response")
	}
}

// Guess handling logic
func handleGuess(w http.ResponseWriter, r *http.Request) {
	// Retrieve the 'guess' query parameter
	userGuessStr := r.URL.Query().Get("guess")
	if userGuessStr == "" {
		http.Error(w, "Missing 'guess' query parameter", http.StatusBadRequest)
		return
	}

	// Convert 'guess' to an integer
	userGuess, err := strconv.Atoi(userGuessStr)
	if err != nil {
		http.Error(w, "Invalid 'guess' parameter. Must be an integer.", http.StatusBadRequest)
		return
	}

	// Generate a random number
	randomNumber := randomGenerator.Intn(6)
	log.Info().Int("User Guess", userGuess).Msg("Received user guess")
	log.Info().Int("Random Number", randomNumber).Msg("Generated random number")

	// Determine if the user wins or loses
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
