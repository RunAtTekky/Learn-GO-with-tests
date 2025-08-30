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

func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	secondHandWriter(w, t)
	io.WriteString(w, svgEnd)
}

func secondHandWriter(w io.Writer, t time.Time) {
	p := SecondHand(t)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
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
