package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris!"

		assertMessages(t, got, want)
	})

	t.Run("saying hello without specified name", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world!"

		assertMessages(t, got, want)
	})

}

func assertMessages(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
