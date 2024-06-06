package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Del")
		want := "Hello, Del"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("default to Hello, World", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
