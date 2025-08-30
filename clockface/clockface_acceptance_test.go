package clockface_test

import (
	"bytes"
	"encoding/xml"
	"go_with_test/clockface"
	"slices"
	"testing"
	"time"
)

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

func TestSVGWriterAtMidnight(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			time: simpleTime(0, 0, 0),
			line: Line{150, 150, 150, 60},
		},
		{
			simpleTime(0, 0, 30),
			Line{150, 150, 150, 240},
		},
	}

	for _, tt := range cases {
		t.Run(testName(tt.time), func(t *testing.T) {
			b := bytes.Buffer{}
			clockface.SVGWriter(&b, tt.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(tt.line, svg.Line) {
				t.Errorf("Expected to find the second hand line %+v, in the SVG line %+v", tt.line, svg.Line)
			}
		})
	}

}

func containsLine(a Line, lst []Line) bool {
	return slices.Contains(lst, a)
}

func TestSecondHandAtMidnight(t *testing.T) {
	tm := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

	want := clockface.Point{X: 150, Y: 150 - 90}
	got := clockface.SecondHand(tm)

	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}

func TestSecondHandAt30Seconds(t *testing.T) {
	tm := time.Date(2025, time.July, 1, 0, 0, 30, 0, time.UTC)

	want := clockface.Point{X: 150, Y: 150 + 90}
	got := clockface.SecondHand(tm)

	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}

func simpleTime(hour, minute, second int) time.Time {
	return time.Date(2025, time.July, 0, hour, minute, second, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}
