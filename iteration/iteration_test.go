package iteration

import (
	"testing"
)

func TestIteration(t *testing.T) {
	t.Run("Iterate 5 times", func(t *testing.T) {
		repeat := Iteration(Params{element: "a", times: 5})
		expected := "aaaaa"

		if repeat != expected {
			t.Errorf("Got only %s, expected %s", repeat, expected)
		}
	})
}

func BenchmarkIteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iteration(Params{element: "a", times: 5})
	}
}
