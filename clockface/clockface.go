package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

const (
	secondHandLength = 90
	clockCentreX     = 150
	clockCentreY     = 150
)

func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)

	// Scale
	p = Point{p.X * secondHandLength, p.Y * secondHandLength}

	// Flip
	p = Point{p.X, -p.Y}

	// Translate
	p = Point{p.X + clockCentreX, p.Y + clockCentreY}

	return p
}

func secondsInRadians(t time.Time) float64 {
	secondInFloat := float64(t.Second())
	return math.Pi / (30 / secondInFloat)
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
