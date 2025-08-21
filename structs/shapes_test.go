package structs

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("Test rectangle perimeter", func(t *testing.T) {
		rect := Rectangle{12.0, 8.0}
		got := rect.Perimeter()

		want := 40.0
		if got != want {
			t.Errorf("got %g but want %g", got, want)
		}
	})
}

func TestArea(t *testing.T) {

	area_tests := []struct {
		name     string
		shape    Shape
		has_area float64
	}{
		{name: "Rectangle", shape: Rectangle{width: 12.0, height: 8.0}, has_area: 96.0},
		{name: "Circle", shape: Circle{radius: 10.0}, has_area: 314.1592653589793},
		{name: "Triangle", shape: Triangle{height: 10.0, base: 2.0}, has_area: 10.0},
	}

	for _, tt := range area_tests {
		t.Run(tt.name, func(t *testing.T) {
			area := tt.shape.Area()
			if area != tt.has_area {
				t.Errorf("%#v got %g but want %g", tt.shape, area, tt.has_area)
			}
		})
	}
}
