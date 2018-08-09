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

const eof = -1

// peek returns but does not consume the next rune in the input.
func (s *Scanner) Peek() rune {
	r := s.Next()
	s.Backup()
	return r
}

func (s *Scanner) Backup() {
	s.pos -= s.width
	// Correct newline count.
	if s.width == 1 && s.input[s.pos] == '\n' {
		s.line--
	}
}

func (s *Scanner) Next() rune {
	if s.line == 0 {
		s.line = 1
	}
	if s.pos >= len(s.input) {
		s.width = 0
		return eof
	}
	r, w := utf8.DecodeRuneInString(s.input[s.pos:])
	s.width = w
	s.pos += s.width
	if r == '\n' {
		s.line++
	}
	return r
}
