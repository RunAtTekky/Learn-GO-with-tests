package specifications

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

type Greeter interface {
	Greet(name string) (string, error)
}

func GreetSpecification(t testing.TB, greeter Greeter) {
	got, err := greeter.Greet("RunAt")

	assert.NoError(t, err)
	assert.Equal(t, "Hello, RunAt", got)
}

type MeanGreeter interface {
	Curse(name string) (string, error)
}

func MeanSpecification(t testing.TB, meany MeanGreeter) {
	got, err := meany.Curse("RunAt")

	assert.NoError(t, err)
	assert.Equal(t, "Go to hell, RunAt", got)
}
