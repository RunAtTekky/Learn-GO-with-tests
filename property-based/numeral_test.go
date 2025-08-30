package propertybased

import (
	"fmt"
	"testing"
	"testing/quick"
)

var (
	tests = []struct {
		roman  string
		arabic uint16
	}{
		{roman: "I", arabic: 1},
		{roman: "II", arabic: 2},
		{roman: "III", arabic: 3},
		{roman: "IV", arabic: 4},
		{roman: "V", arabic: 5},
		{roman: "VI", arabic: 6},
		{roman: "VII", arabic: 7},
		{roman: "VIII", arabic: 8},
		{roman: "IX", arabic: 9},
		{roman: "XI", arabic: 11},
		{roman: "XII", arabic: 12},
		{roman: "XXII", arabic: 22},
		{roman: "XXXV", arabic: 35},
		{roman: "XXXIX", arabic: 39},
		{roman: "XL", arabic: 40},
		{roman: "XLVII", arabic: 47},
		{roman: "XLIX", arabic: 49},
		{roman: "L", arabic: 50},
		{roman: "M", arabic: 1000},
		{roman: "MCMLXXXIV", arabic: 1984},
		{roman: "MMMCMXCIX", arabic: 3999},
		{roman: "MCMLXXXIV", arabic: 1984},
		{roman: "MMXIV", arabic: 2014},
		{roman: "MVI", arabic: 1006},
		{roman: "DCCXCVIII", arabic: 798},
	}
)

func TestArabicToRoman(t *testing.T) {
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Convert %d to %q", tt.arabic, tt.roman), func(t *testing.T) {
			got := ConvertToRoman(tt.arabic)

			if got != tt.roman {
				t.Errorf("got %q but roman %q", got, tt.roman)
			}
		})
	}
}

func TestRomanToArabic(t *testing.T) {
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Convert %q to %d", tt.roman, tt.arabic), func(t *testing.T) {
			arabic := ConvertToArabic(tt.roman)

			if arabic != tt.arabic {
				t.Errorf("got %d but want %d", arabic, tt.arabic)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		t.Log("testing", arabic)
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, &quick.Config{MaxCount: 1000}); err != nil {
		t.Error("Failed checks", err)
	}
}
