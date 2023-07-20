package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		gpt := Hello("Leonardo", english)
		want := "Hello, Leonardo"

		assertCorrectMessage(t, gpt, want)
	})

	t.Run("say 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		gpt := Hello("", english)
		want := "Hello, World!"

		assertCorrectMessage(t, gpt, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		gpt := Hello("Leonardo", spanish)
		want := "Hola, Leonardo"

		assertCorrectMessage(t, gpt, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		gpt := Hello("Leonardo", french)
		want := "Bonjour, Leonardo"

		assertCorrectMessage(t, gpt, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
