package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 7)
	expected := 9

	if sum != expected {
		t.Errorf("got %d but want %d", sum, expected)
	}
}

func ExampleAdd() {
	sum := Add(1, 9)
	fmt.Println(sum)
	// Output: 10
}
