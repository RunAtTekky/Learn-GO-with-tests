package interactions_test

import (
	"testing"

	"github.com/runattekky/go-specs-greet/domain/interactions"
	"github.com/runattekky/go-specs-greet/specifications"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(
		t,
		specifications.GreetAdapter(interactions.Greet),
	)
}

func TestGreetEmpty(t *testing.T) {
	got := interactions.Greet("")
	want := "Hello, World"

	if got != want {
		t.Fatalf("want %s but got %s", want, got)
	}
}
