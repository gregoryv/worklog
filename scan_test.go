package timesheet

import (
	"fmt"
	. "github.com/gregoryv/qual"
	"testing"
)

type ScanCase struct {
	exp         rune
	line, index int
}

func TestScanner_Scan(t *testing.T) {
	s := NewScanner("shi")
	cases := []struct {
		msg    string
		valid  string
		letter string
		ok     bool
	}{
		{"", "xyz", "", false},
		{"", "s", "s", true},
		{"", "thrsi", "h", true},
		{"", "thrsi", "i", true},
		{"End of string", "thrsi", "", false},
	}

	for _, c := range cases {
		letter, ok := s.Scan(c.valid)
		Assert(t, Vars{c, letter, ok},
			c.ok == ok,
			c.letter == letter,
		)
	}
}

func TestScanner_ScanAll(t *testing.T) {
	s := NewScanner("cab123")
	cases := []struct {
		msg   string
		s     *Scanner
		valid string
		part  string
		ok    bool
	}{
		{"", s, "abcdefghijklmnopqrst", "cab", true},
	}
	for _, c := range cases {
		got, ok := c.s.ScanAll(c.valid)
		Assert(t, Vars{c, got, ok},
			c.part == got,
			c.ok == ok,
		)
	}
}

func ExampleScanner_ScanAll() {
	s := NewScanner("cab123")
	fmt.Print(s.ScanAll("abcdefg"))
	//ouput:
	//cab
}

func ExampleScanner_Next() {
	s := NewScanner("ab\nc")
	for {
		r := s.Peek()
		if r == EOS {
			break
		}
		v := string(r)
		if r == '\n' {
			v = "\\n"
		}

		fmt.Printf("%s: %s\n", s.Pos().String(), v)
		s.Next()
	}
	//output:
	// 1,1: a
	// 1,2: b
	// 1,3: \n
	// 2,1: c
}

func TestNewScanner(t *testing.T) {
	s := NewScanner("")
	if s == nil {
		t.Fail()
	}
}

func TestScanner_Next(t *testing.T) {
	s := NewScanner("abc\nd\ne")
	cases := []ScanCase{
		{'a', 1, 1},
		{'b', 1, 2},
		{'c', 1, 3},
		{'\n', 2, 4},
		{'d', 2, 5},
		{'\n', 3, 6},
		{'e', 3, 7},
		{EOS, 3, 7},
	}
	for _, c := range cases {
		r := s.Next()
		pos := s.Pos()
		line, _ := pos.Val()
		Assert(t, Vars{c, line, r},
			c.exp == r,
			c.line == line,
			c.index == s.index,
		)
	}
}

func TestScanner_Back(t *testing.T) {
	s := NewScanner("abc\nd\ne")
	s.Next()
	s.Back()
	r := s.Next()
	c := ScanCase{'a', 1, 1}
	pos := s.Pos()
	line, _ := pos.Val()
	Assert(t, Vars{c, r, line},
		c.exp == r,
		c.line == line,
		c.index == s.index,
	)
	// Back over a newline
	s = NewScanner("\na")
	s.Next()
	s.Back()
	pos = s.Pos()
	line, _ = pos.Val()
	Assert(t, Vars{line, s.index},
		line == 1,
		s.index == 0,
	)
}

func TestScanner_PeekIs(t *testing.T) {
	for _, c := range []struct {
		txt, valid string
		exp        bool
	}{
		{"abc", "a", true},
		{"abc", "csa", true},
		{"abc", "q", false},
		{"abc", "qbc", false},
	} {
		s := NewScanner(c.txt)
		got := s.PeekIs(c.valid)
		Assert(t, Vars{c, got},
			c.exp == got,
		)
	}
}

func TestScanner_Peek(t *testing.T) {
	s := NewScanner("12")
	r := s.Peek()
	c := ScanCase{'1', 1, 0}
	pos := s.Pos()
	line, _ := pos.Val()
	Assert(t, Vars{c, line},
		c.exp == r,
		c.line == line,
		c.index == s.index,
	)
}
