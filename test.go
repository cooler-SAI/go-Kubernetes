package main

import (
	"fmt"
	"math/rand"
	"time"
)

func testApp() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println("Random Number:", r.Intn(10))
}
