package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	favCards := []int{2, 6, 9}

	return favCards
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if outOfBound(slice, index) {
		return -1
	}

	item := slice[index]

	return item

}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if outOfBound(slice, index) {
		slice = append(slice, value)

		return slice
	}

	slice[index] = value

	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	if values != nil {
		var newSlice []int

		newSlice = append(newSlice, values...)
		newSlice = append(newSlice, slice...)

		return newSlice
	}

	return slice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if outOfBound(slice, index) {
		return slice
	}

	newSlice := slice[:index]
	newSlice2 := slice[index+1:]

	newSlice = append(newSlice, newSlice2...)

	return newSlice
}

// outOfBound returns true if index is negative or after the end of the stack.
func outOfBound(slice []int, index int) bool {
	length := len(slice) - 1

	if index < 0 || index > length {
		return true
	}

	return false
}
