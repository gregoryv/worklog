package timesheet

import (
	"testing"
)

func TestScanner_Next_and_Backup(t *testing.T) {
	scan := &Scanner{input: "abc\nd\ne"}
	cases := []struct {
		exp       rune
		line, pos int
		fn        func() // to run after asserts
	}{
		{'a', 1, 1, scan.Backup},
		{'a', 1, 1, nil},
		{'b', 1, 2, nil},
		{'c', 1, 3, nil},
		{'\n', 2, 4, nil},
		{'d', 2, 5, nil},
		{'\n', 3, 6, scan.Backup},
		{'\n', 3, 6, nil},
		{'e', 3, 7, nil},
		{eof, 3, 7, scan.Backup},
		{eof, 3, 7, nil},
	}

	for _, c := range cases {
		res := scan.Next()
		if res != c.exp {
			t.Errorf("Expected rune %q, got %q", c.exp, res)
		}
		if c.line != scan.line {
			t.Errorf("Expected line %v, got %v", c.line, scan.line)
		}
		if c.pos != scan.pos {
			t.Errorf("Expected pos %v, got %v", c.pos, scan.pos)
		}
		if c.fn != nil {
			c.fn()
		}
	}
}
