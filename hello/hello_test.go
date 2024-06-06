package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Del", "English")
		want := "Hello, Del"
		assertCorrectMessage(t, got, want)
	})

	t.Run("default to Hello, World", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Del", "Spanish")
		want := "Hola, Del"
		assertCorrectMessage(t, got, want)
	})
	t.Run(("in French"), func(t *testing.T) {
		got := Hello("Del", "French")
		want := "Bonjour, Del"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
