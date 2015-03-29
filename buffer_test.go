package main

import "testing"

func TestInsert(t *testing.T) {
	buff := EmptyBuffer()
	buff.Insert("World!")
	assert(t, buff.GetLine(), "World!")
	assert(t, buff.GetLineNum(), 1)

	buff.Insert("Hello,")
	assert(t, buff.GetLine(), "Hello,")
	assert(t, buff.GetLineNum(), 1)

	buff.NextLine()
	assert(t, buff.GetLine(), "World!")
	assert(t, buff.GetLineNum(), 2)
}

func TestAppend(t *testing.T) {
	buff := EmptyBuffer()
	buff.Append("Hello,")
	assert(t, buff.GetLine(), "Hello,")
	assert(t, buff.GetLineNum(), 1)

	buff.Append("World!")
	assert(t, buff.GetLine(), "World!")
	assert(t, buff.GetLineNum(), 2)

	buff.PrevLine()
	assert(t, buff.GetLine(), "Hello,")
	assert(t, buff.GetLineNum(), 1)
}

func TestLineMovement(t *testing.T) {
	buff := EmptyBuffer()
	buff.Append("line 1")
	buff.Append("line 2")
	buff.Append("line 3")

	assert(t, buff.GetLine(), "line 3")
	buff.NextLine()
	assert(t, buff.GetLine(), "line 3")
	buff.PrevLine()
	assert(t, buff.GetLine(), "line 2")
	buff.PrevLine()
	assert(t, buff.GetLine(), "line 1")
	buff.PrevLine()
	assert(t, buff.GetLine(), "line 1")
	buff.NextLine()
	assert(t, buff.GetLine(), "line 2")
}

func TestGoto(t *testing.T) {
	buff := EmptyBuffer()

	buff.Append("line 1")
	buff.Append("line 2")
	buff.Append("line 3")
	buff.Append("line 4")

	buff.Goto(1)
	assert(t, buff.GetLine(), "line 1")
	buff.Goto(3)
	assert(t, buff.GetLine(), "line 3")
	buff.PrevLine()
	assert(t, buff.GetLine(), "line 2")

	buff.Goto(42)
	assert(t, buff.GetLine(), "line 4")
	buff.Goto(-42)
	assert(t, buff.GetLine(), "line 1")
}

func TestGetLine(t *testing.T) {
	buff := EmptyBuffer()
	assert(t, buff.GetLine(), "")
}

func TestLineNumber(t *testing.T) {
	buff := EmptyBuffer()

	buff.Insert("")
	assert(t, buff.GetLineNum(), 1)
	buff.Append("")
	assert(t, buff.GetLineNum(), 2)
	buff.PrevLine()
	assert(t, buff.GetLineNum(), 1)
}

func TestTotalLines(t *testing.T) {
	buff := EmptyBuffer()

	assert(t, buff.GetTotalLines(), 0)
	buff.Append("line 1")
	assert(t, buff.GetTotalLines(), 1)
	buff.Append("line 3")
	assert(t, buff.GetTotalLines(), 2)
	buff.Insert("line 2")
	assert(t, buff.GetTotalLines(), 3)
}

func assert(t *testing.T, a, b interface{}) {
	if a != b {
		t.Errorf("Assertion failed: '%v' != '%v'", a, b)
	}
}
