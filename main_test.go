package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func TestHTTPServer(t *testing.T) {
	log.Info().Msg("Starting TestHTTPServer")

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create request")
		t.Fatalf("Can not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, "Welcome to the automated guessing game!")
	})

	log.Info().Msg("Sending request to handler")
	handler.ServeHTTP(rr, req)

	expected := "Welcome to the automated guessing game!\n"
	if rr.Body.String() != expected {
		log.Error().Str("got", rr.Body.String()).Str("expected", expected).Msg("Response mismatch")
		t.Errorf("Wrong answer: got %v, expected %v", rr.Body.String(), expected)
	} else {
		log.Info().Msg("TestHTTPServer passed successfully")
	}
}

func TestRandomNumberGeneration(t *testing.T) {
	log.Info().Msg("Starting TestRandomNumberGeneration")

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 100; i++ {
		num := r.Intn(6)
		log.Debug().Int("generated_number ", num).Msg("Generated random number")

		if num < 0 || num > 5 {
			log.Error().Int("invalid_number", num).Msg("Generated number out of range")
			t.Errorf("Number not valid: %v", num)
		}
	}

	log.Info().Msg("TestRandomNumberGeneration completed successfully")
}
