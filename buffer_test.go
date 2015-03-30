package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"testing"
)

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

func TestFromFile(t *testing.T) {
	file, _ := ioutil.TempFile(os.TempDir(), "edsrv-test")
	defer os.Remove(file.Name())

	file.WriteString("line 1\nline 2\n")

	buff := BufferFromFile(file.Name())
	assert(t, buff.GetLine(), "line 1")
	buff.NextLine()
	assert(t, buff.GetLine(), "line 2")
	buff.NextLine()
	assert(t, buff.GetLine(), "line 2")
}

func TestFromNewFile(t *testing.T) {
	file, _ := ioutil.TempFile(os.TempDir(), "edsrv-test")
	name := file.Name()
	os.Remove(file.Name())

	buff := BufferFromFile(name)
	assert(t, buff.GetLine(), "")
}

func TestFromEmptyFile(t *testing.T) {
	file, _ := ioutil.TempFile(os.TempDir(), "edsrv-test")
	defer os.Remove(file.Name())

	buff := BufferFromFile(file.Name())
	assert(t, *buff.file, file.Name())
}

func TestWriteToFile(t *testing.T) {
	buff := BufferFromFile("/tmp/edsrv-test")
	assert(t, *buff.file, "/tmp/edsrv-test")
	buff.Append("line 1")
	buff.Append("line 2")

	buff.Write()

	file, _ := os.Open("/tmp/edsrv-test")
	defer os.Remove(file.Name())
	defer file.Close()

	rd := bufio.NewReader(file)

	line, _ := rd.ReadString('\n')
	assert(t, line, "line 1\n")
	line, _ = rd.ReadString('\n')
	assert(t, line, "line 2\n")
}

func TestWriteErrors(t *testing.T) {
	buff := EmptyBuffer()
	buff.Insert("line 1")
	err := buff.Write()
	assert(t, err.Error(), "Buffer has no file associated with it")
}

func assert(t *testing.T, a, b interface{}) {
	if a != b {
		t.Errorf("Assertion failed: '%v' != '%v'", a, b)
	}
}
