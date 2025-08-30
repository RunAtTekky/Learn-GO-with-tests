package clockface

import (
	"math"
	"testing"
	"time"
)

func TestMinuteHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 30, 0), Point{0, -1}},
	}

	for _, tt := range cases {
		t.Run(testName(tt.time), func(t *testing.T) {
			got := minuteHandPoint(tt.time)

			if !roughlyEqualPoint(got, tt.point) {
				t.Fatalf("got %v point, but want %v point", got, tt.point)
			}
		})
	}

}

func TestHoursInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(6, 0, 0), math.Pi},
		{simpleTime(0, 0, 0), 0},
	}

	for _, tt := range cases {
		t.Run(testName(tt.time), func(t *testing.T) {
			want := tt.angle
			got := hoursInRadians(tt.time)

			if got != want {
				t.Errorf("got %v Radians, but wanted %v", got, want)
			}

		})
	}
}

func TestMinutesInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 30, 0), math.Pi},
		{simpleTime(0, 0, 7), 7 * (math.Pi / (30 * 60))},
	}

	for _, tt := range cases {
		t.Run(testName(tt.time), func(t *testing.T) {
			want := tt.angle
			got := minutesInRadians(tt.time)

			if got != want {
				t.Errorf("got %v Radians, but wanted %v", got, want)
			}

		})
	}
}

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}

	for _, tt := range cases {
		t.Run(testName(tt.time), func(t *testing.T) {
			want := tt.point
			got := secondHandPoint(tt.time)

			if !roughlyEqualPoint(want, got) {
				t.Errorf("got %v Point, but want %v", got, want)
			}

		})
	}
}

func roughlyEqualFloat64(a, b float64) bool {
	const threshold = 1e-7
	return math.Abs(a-b) < threshold
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
}

func simpleTime(hour, minute, second int) time.Time {
	return time.Date(2025, time.July, 0, hour, minute, second, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}
