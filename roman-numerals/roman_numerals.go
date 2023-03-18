// Package romannumerals converts int to roman numeral.
package romannumerals

import "errors"

// ToRomanNumeral converts an integer to it's equivelent in Roman.
func ToRomanNumeral(input int) (string, error) {
	if input > 3999 || input <= 0 {
		return "", errors.New("input is not valid")
	}

	romanNums := []struct {
		intValue   int
		romanValue string
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

	result := ""
	repeat := 0

	for _, v := range romanNums {
		repeat = input / v.intValue
		if repeat > 0 {
			for i := 1; i <= repeat; i++ {
				result += v.romanValue
			}

			input = input % v.intValue
		}
	}

	return result, nil
}
