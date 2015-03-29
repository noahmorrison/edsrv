package main

// A Buffer is a view into a file on the servers computer
type Buffer struct {
	prev *Stack
	curr *string
	next *Stack
}

// EmptyBuffer returns a new, empty, buffer
func EmptyBuffer() *Buffer {
	return &Buffer{
		prev: new(Stack),
		curr: nil,
		next: new(Stack),
	}
}

// GetLine returns the current line from the buffer
func (buff *Buffer) GetLine() string {
	if buff.curr != nil {
		return *buff.curr
	}

	return ""
}

// GetLineNum returns the current line number
func (buff Buffer) GetLineNum() int {
	return buff.prev.Len() + 1
}

// GetTotalLines returns the total line count
func (buff Buffer) GetTotalLines() int {
	var current int
	if buff.curr != nil {
		current = 1
	}
	return buff.prev.Len() + buff.next.Len() + current
}

// NextLine shifts the lines forwards
func (buff *Buffer) NextLine() {
	el := buff.next.Pop()
	if el == nil {
		return
	}
	buff.prev.Push(*buff.curr)
	buff.curr = el
}

// PrevLine shifts the lines backwards
func (buff *Buffer) PrevLine() {
	el := buff.prev.Pop()
	if el == nil {
		return
	}

	buff.next.Push(*buff.curr)
	buff.curr = el
}

// Goto shifts the lines to the specified line
func (buff *Buffer) Goto(num int) {
	if num < 1 {
		num = 1
	} else if num > buff.GetTotalLines() {
		num = buff.GetTotalLines()
	}

	for buff.GetLineNum() != num {
		if buff.GetLineNum() > num {
			buff.PrevLine()
		} else {
			buff.NextLine()
		}
	}
}

// Insert puts text before the current line
// And sets the current line to the inserted text
func (buff *Buffer) Insert(text string) {
	if buff.curr != nil {
		buff.next.Push(*buff.curr)
	}

	buff.curr = &text
}

// Append puts text after the current line
// And sets the current line to the appended text
func (buff *Buffer) Append(text string) {
	if buff.curr != nil {
		buff.prev.Push(*buff.curr)
	}

	buff.curr = &text
}
