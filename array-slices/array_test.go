package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Sum 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		sum := Sum(numbers)
		want := 15
		errorMessage(t, sum, want)
	})
	t.Run("Array of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		sum := Sum(numbers)
		want := 6
		errorMessage(t, sum, want)
	})
	t.Run("Varying number of slices", func(t *testing.T) {
		sum := Sum([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
		errorMessage(t, sum, want)
	})
}

func errorMessage(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Sum = %d, want %d", got, want)
	}
}
