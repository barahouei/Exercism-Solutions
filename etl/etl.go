package etl

import "strings"

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
