package main

// A Stack is a  simple FIFO container
type Stack struct {
	top  *Element
	size int
}

// A Element is an element of a Stack
type Element struct {
	value *string
	next  *Element
}

// Len returns the amount of elements in the Stack
func (s Stack) Len() int {
	return s.size
}

// Push adds a new element to the Stack
func (s *Stack) Push(value string) {
	s.top = &Element{&value, s.top}
	s.size++
}

// Pop returns and removes the top element off the Stack
func (s *Stack) Pop() (value *string) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return value
	}

	return nil
}
