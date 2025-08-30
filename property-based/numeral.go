package propertybased

import "strings"

type RomanNumeral struct {
	symbol string
	value  int
}

var allRomanNumerals = []RomanNumeral{
	{value: 1000, symbol: "M"},
	{value: 900, symbol: "CM"},
	{value: 500, symbol: "D"},
	{value: 400, symbol: "CD"},
	{value: 100, symbol: "C"},
	{value: 90, symbol: "XC"},
	{value: 50, symbol: "L"},
	{value: 40, symbol: "XL"},
	{value: 10, symbol: "X"},
	{value: 9, symbol: "IX"},
	{value: 5, symbol: "V"},
	{value: 4, symbol: "IV"},
	{value: 1, symbol: "I"},
}

func ConvertToArabic(roman string) (arabic int) {
	for _, numeral := range allRomanNumerals {
		for strings.HasPrefix(roman, numeral.symbol) {
			arabic += numeral.value
			roman = strings.TrimPrefix(roman, numeral.symbol)
		}
	}
	return arabic
}

func ConvertToRoman(arabic int) string {
	var roman strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.value {
			roman.WriteString(numeral.symbol)
			arabic -= numeral.value
		}
	}

	return roman.String()
}
