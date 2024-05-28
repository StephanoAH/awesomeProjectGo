package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Greetings someone", func(t *testing.T) {
		got := Hello(Params{subject: "Stephano"})
		want := "Hello, Stephano"
		correctMessage(t, got, want)
	})
	t.Run("Fallback to 'Hello, World' when param is empty", func(t *testing.T) {
		got := Hello(Params{subject: ""})
		want := "Hello, World"
		correctMessage(t, got, want)
	})
	t.Run("In spanish", func(t *testing.T) {
		got := Hello(Params{subject: "mundo", lang: "Spanish"})
		want := "Hola, mundo"
		correctMessage(t, got, want)
	})
}

func correctMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
