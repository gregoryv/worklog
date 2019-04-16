// Copyright (c) 2019 Gregory Vinčić. All rights reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package parser

import (
	"strings"
	"unicode/utf8"
)

type Scanner struct {
	index int
	width int
	input string
	line  int
	inTag bool
}

func NewScanner(txt string) *Scanner {
	return &Scanner{input: txt, line: 1}
}

// End Of String
const EOS = -1

func (p *Scanner) PeekIs(valid string) bool {
	return strings.ContainsRune(valid, p.Peek())
}

// peek returns but does not consume the next rune in the input.
func (s *Scanner) Peek() rune {
	r := s.Next()
	s.Back()
	return r
}

func (s *Scanner) Back() {
	s.index -= s.width
	r, _ := utf8.DecodeRuneInString(s.input[s.index:])
	if r == '\n' {
		s.line--
	}
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
		s.line++
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

func (s *Scanner) ScanAll(valid string) (part string, ok bool) {
	for r := s.Next(); strings.ContainsRune(valid, r); r = s.Next() {
		part += string(r)
	}
	s.Back()
	ok = part != ""
	return
}

func (s *Scanner) Pos() Position {
	col := s.index - strings.LastIndex(s.input[:s.index], "\n")
	p := Position{line: s.line, col: col}
	return p
}
