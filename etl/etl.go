// Package etl converts the legacy Scrabble scores format to a new format.
package etl

import "strings"

// Transform returns a map of each character to its point.
func Transform(in map[int][]string) map[string]int {
	m := make(map[string]int)

	for point, list := range in {
		for _, char := range list {
			char = strings.ToLower(char)
			m[char] = point
		}
	}

	return m
}
