package main

import (
	"math/rand"
	"time"
)

// Delay for 0 to interval milliseconds
func randDelay(interval int) {
	rand.Seed(time.Now().UnixNano())
	duration := time.Duration(rand.Intn(interval)) * time.Millisecond

	time.Sleep(duration)
}
