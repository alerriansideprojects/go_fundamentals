package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	assertCorrectSum := func(t *testing.T, got, want int, numbers []int) {
		t.Helper()

		if got != want {
			t.Errorf("\nGot: %d\nWant: %d\nGiven: %v", got, want, numbers)
		}
	}

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assertCorrectSum(t, got, want, numbers)
	})
}
