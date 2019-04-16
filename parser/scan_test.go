// Copyright (c) 2019 Gregory Vinčić. All rights reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package parser

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
		if letter != c.letter {
			t.Errorf("%q, expected %q", letter, c.letter)
		}
		if ok != c.ok {
			t.Errorf("%v, expected %v", ok, c.ok)
		}
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
		exp := c.part
		if got != exp {
			t.Errorf("%q, expected %q", got, exp)
		}
		if ok != c.ok {
			t.Errorf("%v, expected %v", ok, c.ok)
		}
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
		if r != c.exp {
			t.Errorf("%q, expected %q", r, c.exp)
		}
		if line != c.line {
			t.Errorf("%q, expected %q", line, c.line)
		}
		if s.index != c.index {
			t.Errorf("%q, expected %q", s.index, c.index)
		}
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
	if r != c.exp {
		t.Errorf("%q, expected %q", r, c.exp)
	}
	if line != c.line {
		t.Errorf("%q, expected %q", line, c.line)
	}
	if s.index != c.index {
		t.Errorf("%q, expected %q", s.index, c.index)
	}
	// Back over a newline
	s = NewScanner("\na")
	s.Next()
	s.Back()
	pos = s.Pos()
	line, _ = pos.Val()
	exp := 1
	if line != exp {
		t.Errorf("%q, expected %q", line, exp)
	}
	exp = 0
	if s.index != exp {
		t.Errorf("%q, expected %q", s.index, exp)
	}
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
		exp := c.exp
		if got != exp {
			t.Errorf("%v, expected %v", got, exp)
		}
	}
}

func TestScanner_Peek(t *testing.T) {
	s := NewScanner("12")
	r := s.Peek()
	c := ScanCase{'1', 1, 0}
	if r != c.exp {
		t.Errorf("%q, expected %q", r, c.exp)
	}
	pos := s.Pos()
	line, _ := pos.Val()

	if line != c.line {
		t.Errorf("%q, expected %q", line, c.line)
	}
	if s.index != c.index {
		t.Errorf("%q, expected %q", s.index, c.index)
	}
}
