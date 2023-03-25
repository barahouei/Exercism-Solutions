// Package linkedlist works with linked list.
package linkedlist

import "errors"

// Node is a structure of data and next node.
type Node struct {
	Data int
	Next *Node
}

// List is a collection of nodes and their number.
type List struct {
	Head   *Node
	Length int
}

// New returns a linked list of the given elements.
func New(elements []int) *List {
	var list List

	for _, v := range elements {
		list.Push(v)
	}

	return &list
}

// Size returns the number of nodes in the linked list.
func (l *List) Size() int {
	return l.Length
}

// Push adds an item to the linked list.
func (l *List) Push(element int) {
	second := l.Head
	node := &Node{Data: element}

	l.Head = node
	l.Head.Next = second

	l.Length++
}

// Pop returns the top item of the linked list.
func (l *List) Pop() (int, error) {
	if l.Length == 0 {
		return 0, errors.New("can't Pop from empty list")
	}

	toDelete := l.Head.Data
	l.Head = l.Head.Next

	l.Length--

	return toDelete, nil
}

// Array returns an array of the linked lists items
// in the original order.
func (l *List) Array() []int {
	var array []int
	reversedList := l.Reverse()

	for reversedList.Length != 0 {
		n, err := reversedList.Pop()
		if err != nil {
			return []int{}
		}

		array = append(array, n)
	}

	return array
}

// Reverse returns a reversed linked list.
func (l *List) Reverse() *List {
	var reversedList List

	for l.Head != nil {
		n, _ := l.Pop()
		reversedList.Push(n)
	}

	return &reversedList
}
