package timesheet

import (
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
	if s.width == 1 && s.input[s.index] == '\n' {
		s.pos.Back()
	}
}

// Reads the next rune and advances the position by 1 if not at eos
func (s *Scanner) Next() rune {
	if s.index >= len(s.input) {
		s.width = 0
		return EOS
	}
	r, w := utf8.DecodeRuneInString(s.input[s.index:])
	s.width = w
	s.index += s.width
	s.pos.Next()
	if r == '\n' {
		s.pos.NextLine()
	}
	return r
}
