package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Greetings someone", func(t *testing.T) {
		got := Hello(Params{subject: "Stephano"})
		want := "Hello, Stephano"
		errorMessage(t, got, want)
	})
	t.Run("Fallback to 'Hello, World' when param is empty", func(t *testing.T) {
		got := Hello(Params{subject: ""})
		want := "Hello, World"
		errorMessage(t, got, want)
	})
	t.Run("In spanish", func(t *testing.T) {
		got := Hello(Params{subject: "mundo", lang: "Spanish"})
		want := "Hola, mundo"
		errorMessage(t, got, want)
	})
}

func errorMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
