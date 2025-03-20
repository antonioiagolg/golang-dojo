package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Antonio")
		want := "Hello, Antonio"
		assertCorrectMessage(t, got, want)

	})

	t.Run("Saying Hello, World! when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q, want %q\n", got, want)
	}
}
