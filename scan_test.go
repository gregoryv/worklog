package timesheet

import (
	"testing"
)

type ScanCase struct {
	exp       rune
	line, pos int
}

func TestScanner_movement(t *testing.T) {
	scan := &Scanner{input: "abc\nd\ne"}
	cases := []ScanCase{
		{'a', 1, 1},
		{'b', 1, 2},
		{'c', 1, 3},
		{'\n', 2, 4},
		{'d', 2, 5},
		{'\n', 3, 6},
		{'e', 3, 7},
		//		{eof, 3, 7},
	}

	for _, c := range cases {
		r := scan.Next()
		check(c, r, scan, t)
	}

	// Now check them by going backward and forwards
	for i := len(cases) - 1; i >= 0; i-- {
		scan.Backup()
		r := scan.Next()
		c := cases[i]
		t.Logf("%v", c)
		check(c, r, scan, t)
		scan.Backup()
	}
}

func check(c ScanCase, res rune, scan *Scanner, t *testing.T) {
	t.Helper()
	if res != c.exp {
		t.Errorf("Expected rune %q, got %q", c.exp, string(res))
	}
	if c.line != scan.line {
		t.Errorf("Expected line %v, got %v", c.line, scan.line)
	}
	if c.pos != scan.pos {
		t.Errorf("Expected pos %v, got %v", c.pos, scan.pos)
	}
}
