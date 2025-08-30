package clockface

import (
	"math"
	"testing"
	"time"
)

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 45), (math.Pi / 2) * 3},
		{simpleTime(0, 0, 7), (math.Pi / 30) * 7},
	}

	for _, tt := range cases {
		t.Run(testName(tt.time), func(t *testing.T) {
			want := tt.angle
			got := secondsInRadians(tt.time)

			if got != want {
				t.Errorf("got %v radians, but want %v", got, want)
			}

		})
	}
}

func simpleTime(hour, minute, second int) time.Time {
	return time.Date(2025, time.July, 0, hour, minute, second, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}
