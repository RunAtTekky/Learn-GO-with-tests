package structs

import "testing"

func TestPerimeter(t *testing.T) {
	check_perimeter := func(t testing.TB, shape Shape, want float64) {
		got := shape.Perimeter()
		if got != want {
			t.Errorf("got %g but want %g", got, want)
		}
	}

	t.Run("Test rectangle perimeter", func(t *testing.T) {
		rect := Rectangle{12.0, 8.0}
		check_perimeter(t, rect, 40.0)
	})
}

func TestArea(t *testing.T) {

	area_tests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12.0, 8.0}, 96.0},
		{Circle{10.0}, 314.1592653589793},
	}

	check_area := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got %g but want %g", got, want)
		}
	}

	for _, tt := range area_tests {
		check_area(t, tt.shape, tt.want)
	}
}
