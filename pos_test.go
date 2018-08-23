package timesheet

import (
	"fmt"
	. "github.com/gregoryv/qual"
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
		Assert(t, Vars{c.a, c.b, got},
			got == c.exp,
		)
	}
}

func TestPosition_Val(t *testing.T) {
	p := NewPosition()
	line, col := p.Val()
	Assert(t, Vars{line, col},
		line == 1,
		col == 1,
	)
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
