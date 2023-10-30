package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to a world", func(t *testing.T) {
		got := Hello("Bubble")
		want := "Hello, Bubble World."

		assertCorrectMessage(t, got, want)
	})
	t.Run("handle empty world string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, New World."
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
