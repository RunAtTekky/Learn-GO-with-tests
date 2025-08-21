package main

import "fmt"

const english_prefix = "Hello"

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("%s, %s", english_prefix, name)
}

func main() {
}
