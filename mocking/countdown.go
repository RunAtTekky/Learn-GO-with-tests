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
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := COUNT_DOWN_START; i > 0; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}
	fmt.Fprint(writer, FINAL_WORD)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
