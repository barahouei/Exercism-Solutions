// Package anagram deals with strings in order to check the anagram.
package anagram

import (
	"strings"
)

// To solve this exercise, you should consider a couple of things:
// 	1. Checking process must be case-insensitive,
//		so all phrases must be lower-case or upper-case.
// 	2. The length of the subject and candidate phrases must be equal.
//	3. Subject and candidate phrases must not be the same.
//	4. The number of characters used in phrases must be equal,
//		for example: if the character "a" is used once in the subject,
//		it should be used precisely once in the candidate phrase too, and not two times or zero.

// Detect checks if some phrases are an anagram of a given phrase.
func Detect(subject string, candidates []string) []string {
	result := []string{}

	for _, candidate := range candidates {
		if isAnagram(subject, candidate) {
			result = append(result, candidate)
		}
	}

	return result
}

// isAnagram checks if a phrase is an anagram of another phrase.
func isAnagram(subject, candidate string) bool {
	subject = strings.ToLower(subject)
	candidate = strings.ToLower(candidate)

	if len(subject) != len(candidate) {
		return false
	}

	if strings.EqualFold(subject, candidate) {
		return false
	}

	subChars := make(map[rune]int)
	candChars := make(map[rune]int)

	for _, v := range subject {
		subChars[v]++
	}

	for _, v := range candidate {
		candChars[v]++
	}

	for k, v := range subChars {
		if v != candChars[k] {
			return false
		}
	}

	return true
}
