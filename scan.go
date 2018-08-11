package timesheet

import (
	"strings"
	"unicode/utf8"
)

type Scanner struct {
	index int
	width int
	input string
	pos   *Position
}

func NewScanner(txt string) *Scanner {
	return &Scanner{input: txt, pos: NewPosition()}
}

// End Of String
const EOS = -1

// peek returns but does not consume the next rune in the input.
func (s *Scanner) Peek() rune {
	r := s.Next()
	s.Back()
	return r
}

func (s *Scanner) Back() {
	s.index -= s.width
	s.pos.Back()
}

// Reads the next rune and advances the position by 1 if not at EOS
func (s *Scanner) Next() rune {
	if s.index >= len(s.input) {
		s.width = 0
		return EOS
	}
	r, w := utf8.DecodeRuneInString(s.input[s.index:])
	s.width = w
	s.index += s.width
	if r == '\n' {
		s.pos.NextLine()
	} else {
		s.pos.Next()
	}
	return r
}

func (s *Scanner) Scan(valid string) (letter string, ok bool) {
	r := s.Next()
	if strings.ContainsRune(valid, r) {
		return string(r), true
	}
	s.Back()
	return
}

func (s *Scanner) ScanAll(valid string) (part string) {
	for r := s.Next(); strings.ContainsRune(valid, r); r = s.Next() {
		part += string(r)
	}
	s.Back()
	return
}

func (s *Scanner) Pos() *Position {
	return s.pos
}
