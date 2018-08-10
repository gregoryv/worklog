package timesheet

import (
	"unicode/utf8"
)

type Scanner struct {
	line  int
	pos   int
	width int
	input string
}

const EOS = -1

// peek returns but does not consume the next rune in the input.
func (s *Scanner) Peek() rune {
	r := s.Next()
	s.Back()
	return r
}

func (s *Scanner) Back() {
	s.pos -= s.width
	// Correct newline count.
	if s.width == 1 && s.input[s.pos] == '\n' {
		s.line--
	}
}

// Reads the next rune and advances the position by 1 if not at eos
func (s *Scanner) Next() rune {
	if s.line == 0 {
		s.line = 1
	}
	if s.pos >= len(s.input) {
		s.width = 0
		return EOS
	}
	r, w := utf8.DecodeRuneInString(s.input[s.pos:])
	s.width = w
	s.pos += s.width
	if r == '\n' {
		s.line++
	}
	return r
}
