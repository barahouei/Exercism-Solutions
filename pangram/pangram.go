// Package Pangram checks if a given string contains all English alphabet.
package pangram

import (
	"strings"
)

// IsPangram checks if a string is pangram.
func IsPangram(input string) bool {
	input = strings.ToLower(input)

	for i := 'a'; i <= 'z'; i++ {
		if !strings.ContainsRune(input, i) {
			return false
		}
	}

	return true
}
