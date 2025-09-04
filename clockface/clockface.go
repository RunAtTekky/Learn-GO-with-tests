package clockface

import (
	"fmt"
	"io"
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

const (
	secondHandLength = 90
	minuteHandLength = 80
	hourHandLength   = 50
	clockCentreX     = 150
	clockCentreY     = 150

	secondsInHalfClock = 30
	secondsInClock     = 60
	minutesInHalfClock = 30
	minutesInClock     = 60
	hoursInHalfClock   = 6
	hoursInClock       = 12
)

func makeHand(p Point, length float64) Point {
	// Scale
	p = Point{p.X * length, p.Y * length}

	// Flip
	p = Point{p.X, -p.Y}

	// Translate
	p = Point{p.X + clockCentreX, p.Y + clockCentreY}

	return p
}

func SecondHand(w io.Writer, t time.Time) Point {
	p := makeHand(secondHandPoint(t), secondHandLength)

	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)

	return p
}

func MinuteHand(w io.Writer, t time.Time) Point {
	p := makeHand(minuteHandPoint(t), minuteHandLength)

	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)

	return p
}

func HourHand(w io.Writer, t time.Time) Point {
	p := makeHand(hourHandPoint(t), hourHandLength)

	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)

	return p
}

func secondsInRadians(t time.Time) float64 {
	secondInFloat := float64(t.Second())
	return math.Pi / (secondsInHalfClock / secondInFloat)
}

func minutesInRadians(t time.Time) float64 {
	minuteInFloat := float64(t.Minute())
	return (secondsInRadians(t) / secondsInClock) + (math.Pi / (minutesInHalfClock / minuteInFloat))
}

func hoursInRadians(t time.Time) float64 {
	hoursInFloat := float64(t.Hour() % 12)
	return (minutesInRadians(t) / minutesInClock) + (math.Pi / (hoursInHalfClock / hoursInFloat))
}

func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}

func hourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadians(t))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}

func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	SecondHand(w, t)
	MinuteHand(w, t)
	HourHand(w, t)
	io.WriteString(w, svgEnd)
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`
