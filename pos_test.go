package timesheet

import (
	"fmt"
	"testing"
)

func TestPos_NextLine(t *testing.T) {
	cases := []struct {
		msg       string
		p         Pos
		line, col int
	}{
		{"", Pos{1, 1}, 2, 1},
		{"Reset column when moving to next line", Pos{1, 5}, 2, 1},
	}
	for _, c := range cases {
		line, col := c.p.NextLine()
		if err := compareLineCol(c.msg, c.line, line, c.col, col); err != nil {
			t.Error(err)
		}
	}
}

func TestPos_Next(t *testing.T) {
	cases := []struct {
		msg       string
		p         Pos
		line, col int
	}{
		{"Advance column by 1", Pos{1, 1}, 1, 2},
		{"", Pos{3, 5}, 3, 6},
	}
	for _, c := range cases {
		line, col := c.p.Next()
		if err := compareLineCol(c.msg, c.line, line, c.col, col); err != nil {
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

func ExamplePos_String() {
	p := NewPos()
	fmt.Println(p)
	//output:
	//1,1
}

func TestNewPos(t *testing.T) {
	if p := NewPos(); p == nil {
		t.Fail()
	}
}
