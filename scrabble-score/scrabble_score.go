// Package scrabble scores a Scrabble word.
package scrabble

import "unicode"

// Score computes the Scrabble score for the word.
func Score(word string) int {
	total := 0

	for _, v := range word {
		switch unicode.ToUpper(v) {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			total += 1
		case 'D', 'G':
			total += 2
		case 'B', 'C', 'M', 'P':
			total += 3
		case 'F', 'H', 'V', 'W', 'Y':
			total += 4
		case 'K':
			total += 5
		case 'J', 'X':
			total += 8
		case 'Q', 'Z':
			total += 10
		}
	}

	return total
}
