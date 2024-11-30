package main

import (
	"github.com/rs/zerolog"
	"net/http"
	"os"
	"os/signal"

	"github.com/rs/zerolog/log"

	"go-Kubernetes/game"
	"go-Kubernetes/handlers"
)

func main() {
	// Initialize logger
	initializeLogger()

	log.Info().Msg("Hello Kubernetes! Starting automated guessing game.")

	// Initialize random generator
	game.InitializeRandomGenerator()

	// Gracefully handle OS interrupt signals
	setupSignalHandler()

	// Start the background guessing game
	go game.StartGuessingGame()

	// Configure HTTP routes
	setupRoutes()

	// Start the HTTP server
	startHTTPServer()
}

func initializeLogger() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
}

func setupSignalHandler() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	go func() {
		<-signalChan
		log.Info().Msg("Received an interrupt, stopping services...")
		os.Exit(0)
	}()
}

func setupRoutes() {
	http.HandleFunc("/", handlers.HandleRoot)
	http.HandleFunc("/guess", handlers.HandleGuess)
}

func startHTTPServer() {
	log.Info().Msg("Starting HTTP server on port 8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
