package parser

import (
	"fmt"
	"testing"
)

type TestCase struct {
	msg       string
	pos       Position
	line, col int
}

func p(line, col int) Position {
	p := NewPosition()
	p.line = line
	p.col = col
	return *p
}

func TestPosition_Equals(t *testing.T) {
	cases := []struct {
		a, b Position
		exp  bool
	}{
		{Position{1, 1}, Position{1, 1}, true},
		{Position{1, 1}, Position{1, 2}, false},
		{Position{2, 1}, Position{1, 1}, false},
	}
	for _, c := range cases {
		got := c.a.Equals(c.b)
		exp := c.exp
		if got != exp {
			t.Errorf("%v, expected %v", got, exp)
		}

	}
}

func TestPosition_Val(t *testing.T) {
	p := NewPosition()
	line, col := p.Val()
	expLine := 1
	if line != expLine {
		t.Errorf("%q, expected %q", line, expLine)
	}
	expCol := 1
	if col != expCol {
		t.Errorf("%q, expected %q", col, expCol)
	}
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
