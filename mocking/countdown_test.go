package main

import (
	"reflect"
	"testing"
	"time"
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

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.SetDurationSlept}

	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("slept for %v but should have slept for %v", spyTime.durationSlept, sleepTime)
	}
}
