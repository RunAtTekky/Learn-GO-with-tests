package main

import "fmt"

func Hello(name string) string {
	if name == "" {
		return fmt.Sprintf("Hello, World")
	}
	return fmt.Sprintf("Hello, %s", name)
}

func main() {
}
