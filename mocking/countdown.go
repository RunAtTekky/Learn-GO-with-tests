package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	COUNT_DOWN_START = 3
	FINAL_WORD       = "GO!"
	SLEEP            = "sleep"
	WRITE            = "write"
)

type Sleeper interface {
	Sleep()
}

// Configurable Struct where we pass the duration of each sleep required and the sleep method to be used
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// Structure to mock sleeping
type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) SetDurationSlept(duration time.Duration) {
	s.durationSlept = duration
}

// Implements sleeper and writer interfaces and keeps track of Order of operations
type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, SLEEP)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, WRITE)
	return
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := COUNT_DOWN_START; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(writer, i)
	}
	sleeper.Sleep()
	fmt.Fprint(writer, FINAL_WORD)
}

func main() {
	sleeper := &ConfigurableSleeper{duration: 1 * time.Second, sleep: time.Sleep}
	Countdown(os.Stdout, sleeper)
}
