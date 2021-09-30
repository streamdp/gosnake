package game

import "testing"

func TestGetRandomDirection(t *testing.T) {
	got := getRandomDirection()
	if got < 1 && got > 4 {
		t.Errorf("got %d; want %v", got, "should be [1..4]")
	}
}
