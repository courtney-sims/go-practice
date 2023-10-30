package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to a world", func(t *testing.T) {
		got := Hello("Bubble")
		want := "Hello, Bubble World."

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("handle empty world string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, New World."

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}
