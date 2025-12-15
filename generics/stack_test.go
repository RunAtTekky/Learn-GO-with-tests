package generics

import "testing"

func TestStack(t *testing.T) {
	t.Run("Test Integer Stack", func(t *testing.T) {
		myNewStack := new(Stack[int])

		AssertTrue(t, myNewStack.IsEmpty())

		myNewStack.Push(7)
		AssertFalse(t, myNewStack.IsEmpty())

		myNewStack.Push(10)

		val, ok := myNewStack.Pop()
		AssertTrue(t, ok)

		AssertEqual(t, val, 10)

		val, ok = myNewStack.Pop()
		AssertTrue(t, ok)

		AssertEqual(t, val, 7)
		AssertTrue(t, myNewStack.IsEmpty())
	})

	t.Run("Test String Stack", func(t *testing.T) {
		myNewStack := new(Stack[string])

		const (
			messi   = `MESSI`
			ronaldo = `RONALDO`
		)

		AssertTrue(t, myNewStack.IsEmpty())

		myNewStack.Push(ronaldo)
		AssertFalse(t, myNewStack.IsEmpty())

		myNewStack.Push(messi)

		val, ok := myNewStack.Pop()
		AssertTrue(t, ok)

		AssertEqual(t, val, messi)

		val, ok = myNewStack.Pop()
		AssertTrue(t, ok)

		AssertEqual(t, val, ronaldo)
		AssertTrue(t, myNewStack.IsEmpty())
	})
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Fatalf("Did not want %+v", got)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Fatalf("got %v but want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Fatalf("got %v but want false", got)
	}
}
