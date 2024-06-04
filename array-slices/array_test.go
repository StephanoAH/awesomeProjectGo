package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	errorMessage := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("Sum = %d, want %d", got, want)
		}
	}
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
}

func TestSumAllTails(t *testing.T) {
	errorMessageSlices := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Sum = %d, want %d", got, want)
		}
	}

	t.Run("Varying number of slices", func(t *testing.T) {
		sum := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
		errorMessageSlices(t, sum, want)

	})
	t.Run("Sum all tails in a varying number of slices", func(t *testing.T) {
		sum := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		errorMessageSlices(t, sum, want)
	})
	t.Run("Safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2, 3})
		want := []int{0, 5}
		errorMessageSlices(t, got, want)
	})
}
