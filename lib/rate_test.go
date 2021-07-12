package lib

import (
	"testing"
	"time"
)

func TestNewRateBucket(t *testing.T) {
	success, failed := 0, 0
	// Create a new rate-limiter, allowing up-to 10 calls
	rl := NewRateBucket(10, time.Second)
	for i := 0; i < 20; i++ {
		if rl.Limit() {
			success++
		} else {
			failed++
		}
	}
	if success != 10 || failed != 10 {
		t.Fatal(success, failed)
	}

	success = 0
	rl.UpdateRate(15)
	for i := 0; i < 20; i++ {
		if rl.Limit() {
			success++
		}
	}
	if success != 15 {
		t.Fatal(success)
	}
}
