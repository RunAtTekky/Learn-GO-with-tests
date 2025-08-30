package propertybased

import "testing"

func TestRomanNumeral(t *testing.T) {
	tests := []struct {
		name string
		want string
		num  int
	}{
		{name: "Convert 1 to I", want: "I", num: 1},
		{name: "Convert 2 to II", want: "II", num: 2},
		{name: "Convert 3 to III", want: "III", num: 3},
		{name: "Convert 4 to IV", want: "IV", num: 4},
		{name: "Convert 5 to V", want: "V", num: 5},
		{name: "Convert 6 to VI", want: "VI", num: 6},
		{name: "Convert 7 to VII", want: "VII", num: 7},
		{name: "Convert 8 to VIII", want: "VIII", num: 8},
		{name: "Convert 9 to IX", want: "IX", num: 9},
		{name: "Convert 11 to XI", want: "XI", num: 11},
		{name: "Convert 12 to XII", want: "XII", num: 12},
		{name: "Convert 22 to XXII", want: "XXII", num: 22},
		{name: "Convert 35 to XXXV", want: "XXXV", num: 35},
		{name: "Convert 39 to XXXIX", want: "XXXIX", num: 39},
	}

	for _, tt := range tests {
		got := ConvertToRoman(tt.num)

		if got != tt.want {
			t.Errorf("got %q but want %q", got, tt.want)
		}
	}
}
