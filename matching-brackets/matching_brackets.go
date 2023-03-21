// Package brackets implement a way to validate brackets.
package brackets

// Bracket validates brackets such as [], {}, () to be in the correct order and paired.
func Bracket(input string) bool {
	openToClose := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	closeToOpen := map[rune]rune{
		']': '[',
		')': '(',
		'}': '{',
	}

	var brackets stack

	for _, char := range input {
		if _, ok := openToClose[char]; ok {
			brackets.push(char)
		} else if _, ok := closeToOpen[char]; ok {
			if len(brackets.items) == 0 {
				return false
			}

			top := brackets.pop()

			if top != closeToOpen[char] {
				return false
			}
		}
	}

	return len(brackets.items) == 0
}

// A collection of items Vertically.
type stack struct {
	items []rune
}

// push adds an item to the stack.
func (s *stack) push(r rune) {
	s.items = append(s.items, r)
}

// pop returns the top item of the stack.
func (s *stack) pop() rune {
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return top
}
