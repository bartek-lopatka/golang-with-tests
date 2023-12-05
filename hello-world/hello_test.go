package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "en")
		want := "Hello, Chris!"

		assertMessages(t, got, want)
	})

	t.Run("saying hello without specified name", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world!"

		assertMessages(t, got, want)
	})

	t.Run("saying hello in spanish", func(t *testing.T) {
		got := Hello("Jose", "es")
		want := "Hola, Jose!"

		assertMessages(t, got, want)
	})

	t.Run("saying hello in french", func(t *testing.T) {
		got := Hello("Jerome", "fr")
		want := "Bonjour, Jerome!"

		assertMessages(t, got, want)
	})

}

func assertMessages(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
