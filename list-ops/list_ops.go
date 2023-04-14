package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for _, v := range s {
		initial = fn(initial, v)
	}

	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	for i := s.Length() - 1; i >= 0; i-- {
		initial = fn(s[i], initial)
	}

	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	newList := make(IntList, s.Length())

	i := 0
	for _, v := range s {
		if fn(v) {
			newList[i] = v
			i++
		}
	}

	newList = newList[:i]

	return newList
}

func (s IntList) Length() int {
	var countLenght int

	for range s {
		countLenght++
	}

	return countLenght
}

func (s IntList) Map(fn func(int) int) IntList {
	newList := make(IntList, s.Length())

	for i, v := range s {
		newList[i] = fn(v)
	}

	return newList
}

func (s IntList) Reverse() IntList {
	if s.Length() == 0 {
		return s
	}

	reversedList := make(IntList, s.Length())

	for i, j := s.Length()-1, 0; i >= 0; i, j = i-1, j+1 {
		reversedList[j] = s[i]
	}

	return reversedList
}

func (s IntList) Append(lst IntList) IntList {
	ls := s.Length() + lst.Length()
	newList := make([]int, ls)

	for k, v := range s {
		newList[k] = v
	}

	for k, v := range lst {
		newList[s.Length()+k] = v
	}

	return newList
}

func (s IntList) Concat(lists []IntList) IntList {
	for _, list := range lists {
		s = s.Append(list)
	}

	return s
}
