package timesheet

import (
	"fmt"
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
		s      *Scanner
		valid  string
		letter string
		ok     bool
	}{
		{"", s, "xyz", "", false},
		{"", s, "s", "s", true},
		{"", s, "thrsi", "h", true},
		{"", s, "thrsi", "i", true},
		{"End of string", s, "thrsi", "", false},
	}

	for _, c := range cases {
		got, ok := c.s.Scan(c.valid)
		assert(t, c.msg,
			equals("ok", c.ok, ok),
			equals("letter", c.letter, got),
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
		assert(t, c.msg,
			equals("part", c.part, got),
			equals("valid scan", c.ok, ok),
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
		line, _ := s.pos.Val()
		assert(t, "",
			equals("rune", c.exp, r),
			equals("line", c.line, line),
			equals("index", c.index, s.index),
		)
	}
}

func TestScanner_Back(t *testing.T) {
	s := NewScanner("abc\nd\ne")
	s.Next()
	s.Back()
	r := s.Next()
	c := ScanCase{'a', 1, 1}
	line, _ := s.pos.Val()
	assert(t, "",
		equals("rune", c.exp, r),
		equals("line", c.line, line),
		equals("index", c.index, s.index),
	)
	// Back over a newline
	s = NewScanner("\na")
	s.Next()
	s.Back()
	line, _ = s.pos.Val()
	assert(t, "Back over a newline",
		equals("line", line, 1),
		equals("index", s.index, 0),
	)
}

func TestScanner_Peek(t *testing.T) {
	s := NewScanner("12")
	r := s.Peek()
	c := ScanCase{'1', 1, 0}
	line, _ := s.pos.Val()
	assert(t, "",
		equals("rune", c.exp, r),
		equals("line", c.line, line),
		equals("index", c.index, s.index),
	)
}
