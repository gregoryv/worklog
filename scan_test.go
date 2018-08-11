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
		if c.letter != got {
			t.Errorf("Expected %q, got %q", c.letter, got)
		}
		if c.ok != ok {
			t.Errorf("Expected ok = %v, got %v", c.ok, ok)
		}
	}
}

func TestScanner_ScanAll(t *testing.T) {
	s := NewScanner("cab123")
	got := s.ScanAll("abcdefghijklmnopqrst")
	exp := "cab"
	if exp != got {
		t.Errorf("Expected %q, got %q", exp, got)
	}
	r := s.Next()
	if r != '1' {
		t.Fail()
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
		res := s.Next()
		assert(t, "", check(c, res, s))
	}
}

func TestScanner_Back(t *testing.T) {
	s := NewScanner("abc\nd\ne")
	s.Next()
	s.Back()
	r := s.Next()
	c := ScanCase{'a', 1, 1}
	assert(t, "", check(c, r, s))
	// Back over a newline
	s = NewScanner("\na")
	s.Next()
	s.Back()
	line, _ := s.pos.Val()
	if line != 1 || s.index != 0 {
		t.Fail()
	}
}

func TestScanner_Peek(t *testing.T) {
	s := NewScanner("12")
	res := s.Peek()
	c := ScanCase{'1', 1, 0}
	assert(t, "", check(c, res, s))
}

func check(c ScanCase, r rune, s *Scanner) (err error) {
	if c.exp != r {
		return fmt.Errorf("Expected rune %q, got %q", c.exp, string(r))
	}
	line, _ := s.pos.Val()
	if c.line != line {
		return fmt.Errorf("Expected line %v, got %v", c.line, line)
	}
	if c.index != s.index {
		return fmt.Errorf("Expected index %v, got %v", c.index, s.index)
	}
	return
}
