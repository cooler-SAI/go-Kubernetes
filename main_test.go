package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHTTPServer(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Can not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "Welcome to the automated guessing game!")
	})

	handler.ServeHTTP(rr, req)

	expected := "Welcome to the automated guessing game!\n"
	if rr.Body.String() != expected {
		t.Errorf("Wrong answer: got %v, guess %v", rr.Body.String(), expected)
	}
}

func TestRandomNumberGeneration(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 100; i++ {
		num := r.Intn(6)
		if num < 0 || num > 5 {
			t.Errorf("Number not valid")
		}
	}
}
