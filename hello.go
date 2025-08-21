package main

import (
	"fmt"
)

const (
	spanish = "spanish"
	french  = "french"

	english_prefix = "Hello"
	spanish_prefix = "Hola"
	french_prefix  = "Bonjour"
)

func get_greeting(language string) (greeting string) {
	switch language {
	case spanish:
		greeting = spanish_prefix
	case french:
		greeting = french_prefix
	default:
		greeting = english_prefix
	}

	return
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return fmt.Sprintf("%s, %s", get_greeting(language), name)

}

func main() {
}
