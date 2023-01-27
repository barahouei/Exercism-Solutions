// Package isogram checks for isogram.
package isogram

import (
	"unicode"
)

// IsIsogram checks if a word or phrase is an isogram.
func IsIsogram(word string) bool {
	iso := make(map[rune]int)

	for _, v := range word {
		v = unicode.ToLower(v)

		if unicode.IsSpace(v) || v == '-' {
			continue
		}

		if _, ok := iso[v]; ok {
			return false
		}

		iso[v] += 1
	}

	return true
}
