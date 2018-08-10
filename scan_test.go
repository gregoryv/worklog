package timesheet

import (
	"testing"
)

type ScanCase struct {
	exp         rune
	line, index int
}

func TestScanner_Next(t *testing.T) {
	scan := &Scanner{input: "abc\nd\ne"}
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
		check(c, res, scan, t)
	}
}

func TestScanner_Back(t *testing.T) {
	scan := &Scanner{input: "abc\nd\ne"}
	scan.Next()
	scan.Back()
	r := scan.Next()
	c := ScanCase{'a', 1, 1}
	check(c, r, scan, t)

	// Back over a newline
	scan = &Scanner{input: "\na"}
	scan.Next()
	scan.Back()
	if scan.line != 1 || scan.index != 0 {
		t.Fail()
	}
}

func TestScanner_Peek(t *testing.T) {
	scan := &Scanner{input: "12"}
	res := scan.Peek()
	c := ScanCase{'1', 1, 0}
	check(c, res, scan, t)
}

func check(c ScanCase, r rune, scan *Scanner, t *testing.T) {
	t.Helper()
	if c.exp != r {
		t.Errorf("Expected rune %q, got %q", c.exp, string(r))
	}
	if c.line != scan.line {
		t.Errorf("Expected line %v, got %v", c.line, scan.line)
	}
	if c.index != scan.index {
		t.Errorf("Expected pos %v, got %v", c.index, scan.index)
	}
}
