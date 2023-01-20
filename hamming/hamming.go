// Package hamming Calculates the Hamming Distance between DNA strands.
package hamming

import (
	"fmt"
)

// notEqualErr defines a struct of two strings to handle the not equal error.
type notEqualErr struct {
	strandA string
	strandB string
}

// Error creates a custom error for notEqualErrr.
func (s *notEqualErr) Error() string {
	return fmt.Sprintf("length of strand '%s' and strand '%s' must be equal.", s.strandA, s.strandB)
}

// Distance returns the number of differences between two DNA strands.
func Distance(a, b string) (int, error) {
	diff := 0

	if len(a) != len(b) {
		return 0, &notEqualErr{strandA: a, strandB: b}
	}

	for k := range a {
		if a[k] != b[k] {
			diff++
		}
	}

	return diff, nil
}
