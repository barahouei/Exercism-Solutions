// Package luhn validates credit cards.
package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

// Valid checks if a credit card number is correct.
func Valid(id string) bool {
	nums := []int{}
	sum := 0

	chars := []rune(strings.ReplaceAll(id, " ", ""))

	if len(chars) <= 1 {
		return false
	}

	for i := len(chars) - 1; i >= 0; i-- {
		if !unicode.IsNumber(chars[i]) {
			return false
		}

		number, _ := strconv.Atoi(string(chars[i]))

		if (len(chars)-i)%2 == 0 {
			number *= 2

			if number > 9 {
				number -= 9
			}

			nums = append(nums, number)
		} else {

			nums = append(nums, number)
		}
	}

	for _, v := range nums {
		sum += v
	}

	return sum%10 == 0
}
