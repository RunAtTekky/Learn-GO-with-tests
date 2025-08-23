package main

import (
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	spySleepPrinter := &SpyCountdownOperations{}

	Countdown(spySleepPrinter, spySleepPrinter)

	got := spySleepPrinter.Calls
	want := []string{
		"sleep",
		"write", // 3
		"sleep",
		"write", // 2
		"sleep",
		"write", // 1
		"sleep",
		"write", // GO!
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q but want %q", got, want)
	}
}
