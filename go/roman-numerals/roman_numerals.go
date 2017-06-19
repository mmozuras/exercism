package romannumerals

import "errors"

const testVersion = 3

var conversions = []struct {
	arabic int
	roman  string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ToRomanNumeral(number int) (string, error) {
	if number <= 0 || number >= 4000 {
		return "", errors.New("Number is either too small or too large.")
	}

	result := ""
	for _, conversion := range conversions {
		for number >= conversion.arabic {
			result += conversion.roman
			number -= conversion.arabic
		}
	}

	return result, nil
}
