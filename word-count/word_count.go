// Package wordcount counts the words in a phrase.
package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

// WordCount returns the frequency of words in a phrase.
func WordCount(phrase string) Frequency {
	phrase = strings.ToLower(phrase)
	wordsFrequency := make(Frequency)

	words := strings.Split(phrase, " ")

	for _, word := range words {
		if regexp.MustCompile(`([a-z]+[']+[a-z])`).FindString(word) != "" {
			wordsFrequency[word]++
		} else {
			word = regexp.MustCompile(`([^a-z0-9+])`).ReplaceAllString(word, " ")

			s := strings.Split(word, " ")

			for _, v := range s {
				if v != "" {
					wordsFrequency[v]++
				}
			}
		}
	}

	return wordsFrequency
}
