package interactions_test

import (
	"testing"

	"github.com/runattekky/go-specs-greet/domain/interactions"
	"github.com/runattekky/go-specs-greet/specifications"
)

func TestCurse(t *testing.T) {
	specifications.MeanSpecification(
		t, specifications.GreetAdapter(interactions.Curse),
	)
}
