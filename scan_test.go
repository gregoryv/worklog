package timesheet

import (
	"fmt"
	"testing"
)

type ScanCase struct {
	exp         rune
	line, index int
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
	scan := NewScanner("abc\nd\ne")
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
		res := scan.Next()
		assert(t, "", check(c, res, scan))
	}
}

func TestScanner_Back(t *testing.T) {
	scan := NewScanner("abc\nd\ne")
	scan.Next()
	scan.Back()
	r := scan.Next()
	c := ScanCase{'a', 1, 1}
	assert(t, "", check(c, r, scan))
	// Back over a newline
	scan = NewScanner("\na")
	scan.Next()
	scan.Back()
	line, _ := scan.pos.Val()
	if line != 1 || scan.index != 0 {
		t.Fail()
	}
}

func TestScanner_Peek(t *testing.T) {
	scan := NewScanner("12")
	res := scan.Peek()
	c := ScanCase{'1', 1, 0}
	assert(t, "", check(c, res, scan))
}

func check(c ScanCase, r rune, scan *Scanner) (err error) {
	if c.exp != r {
		return fmt.Errorf("Expected rune %q, got %q", c.exp, string(r))
	}
	line, _ := scan.pos.Val()
	if c.line != line {
		return fmt.Errorf("Expected line %v, got %v", c.line, line)
	}
	if c.index != scan.index {
		return fmt.Errorf("Expected index %v, got %v", c.index, scan.index)
	}
	return
}
