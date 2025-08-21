package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("RunAt")
		want := "Hello, RunAt"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("saying 'Hello, World' when supplied empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
