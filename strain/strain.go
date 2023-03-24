// Package strain checks for predicates on the items.
package strain

type Ints []int
type Lists [][]int
type Strings []string

// Keep returns a new slice of ints where the predicate is true.
func (i Ints) Keep(filter func(int) bool) Ints {
	var list []int

	for _, v := range i {
		if filter(v) {
			list = append(list, v)
		}
	}

	return list
}

// Keep returns a new slice of ints where the predicate is false.
func (i Ints) Discard(filter func(int) bool) Ints {
	var list []int

	for _, v := range i {
		if !filter(v) {
			list = append(list, v)
		}
	}

	return list
}

// Keep returns a new slice of [][]ints where the predicate is true.
func (l Lists) Keep(filter func([]int) bool) Lists {
	var list [][]int

	for _, v := range l {
		if filter(v) {
			list = append(list, v)
		}
	}

	return list
}

// Keep returns a new slice of strings where the predicate is true.
func (s Strings) Keep(filter func(string) bool) Strings {
	var list []string

	for _, v := range s {
		if filter(v) {
			list = append(list, v)
		}
	}

	return list
}
