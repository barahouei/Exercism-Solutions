// Package twofer provides a way to share items between individuals.
package twofer

import "fmt"

// ShareWith returns a string that shows one item is for me and one for another.
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
