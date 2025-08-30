package propertybased

import "strings"

type RomanNumeral struct {
	symbol string
	value  int
}

var allRomanNumerals = []RomanNumeral{
	{value: 10, symbol: "X"},
	{value: 9, symbol: "IX"},
	{value: 5, symbol: "V"},
	{value: 4, symbol: "IV"},
	{value: 1, symbol: "I"},
}

func ConvertToRoman(num int) string {
	var roman strings.Builder

	for _, numeral := range allRomanNumerals {
		for num >= numeral.value {
			roman.WriteString(numeral.symbol)
			num -= numeral.value
		}
	}

	return roman.String()
}
