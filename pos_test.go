package timesheet

import (
	"fmt"
	"testing"
)

type TestCase struct {
	msg       string
	pos       Position
	line, col int
}

func TestPosition_Back(t *testing.T) {
	cases := []TestCase{
		{"Stay on first, when already there", Position{1, 1}, 1, 1},
		{"Only backup column", Position{1, 5}, 1, 4},
		{"Backup line when possible", Position{2, 1}, 1, 1},
	}
	for _, c := range cases {
		line, col := c.pos.Back()
		assert(t, compareLineCol(c.msg, c.line, line, c.col, col))
	}
}

func TestPosition_NextLine(t *testing.T) {
	cases := []struct {
		msg       string
		pos       Position
		line, col int
	}{
		{"", Position{1, 1}, 2, 1},
		{"Reset column when moving to next line", Position{1, 5}, 2, 1},
	}
	for _, c := range cases {
		line, col := c.pos.NextLine()
		assert(t, compareLineCol(c.msg, c.line, line, c.col, col))
	}
}

func TestPosition_Next(t *testing.T) {
	cases := []struct {
		msg       string
		pos       Position
		line, col int
	}{
		{"Advance column by 1", Position{1, 1}, 1, 2},
		{"", Position{3, 5}, 3, 6},
	}
	for _, c := range cases {
		line, col := c.pos.Next()
		assert(t, compareLineCol(c.msg, c.line, line, c.col, col))
	}
}

func assert(t *testing.T, errors ...error) {
	t.Helper()
	for _, err := range errors {
		if err != nil {
			t.Error(err)
		}
	}
}

func compareLineCol(msg string, expLine, line, expCol, col int) (err error) {
	switch {
	case expLine != line:
		err = fmt.Errorf("%s\nExpected line %v, got %v", msg, expLine, line)
	case expCol != col:
		err = fmt.Errorf("%s\nExpected col %v, got %v", msg, expCol, col)
	}
	return
}

func ExamplePosition_String() {
	pos := NewPosition()
	fmt.Println(pos)
	//output:
	//1,1
}

func TestNewPosition(t *testing.T) {
	if pos := NewPosition(); pos == nil {
		t.Fail()
	}
}
