package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("RunAt", "")
		want := "Hello, RunAt"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("saying 'Hello, World' when supplied empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("saying 'Hola, World' when supplied Spanish", func(t *testing.T) {
		got := Hello("", "spanish")
		want := "Hola, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("saying 'Bonjour, World' when supplied French", func(t *testing.T) {
		got := Hello("", "french")
		want := "Bonjour, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}
