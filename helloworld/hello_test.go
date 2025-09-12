package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris"
		assertCorrectionMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("World", "English")
		want := "Hello, World"
		assertCorrectionMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectionMessage(t, got, want)
	})

	t.Run("in German", func(t *testing.T) {
		got := Hello("Fritz", "German")
		want := "Hallo, Fritz"
		assertCorrectionMessage(t, got, want)
	})

	t.Run("in Italian", func(t *testing.T) {
		got := Hello("Francesca", "Italian")
		want := "Ciao, Francesca"
		assertCorrectionMessage(t, got, want)
	})
}

func assertCorrectionMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
