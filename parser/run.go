// Copyright (c) 2019 Gregory Vinčić. All rights reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package parser

import (
	"strings"

	"github.com/gregoryv/go-timesheet/token"
)

func lexYear(s *Scanner) (p Part, next lexFn) {
	p, next = ScanPart(s, token.Year), lexMonth
	s.Scan(" ")
	return
}

func ScanPart(s *Scanner, tok token.Token) (p Part) {
	p = Part{Tok: tok, Pos: s.Pos()}
	var valid string
	switch tok {
	case token.Hours, token.Minutes, token.Year, token.Date, token.Week:
		valid = digits
	case token.Separator:
		valid = "-"
	}
	val, ok := s.ScanAll(valid)
	if !ok {
		p.Errorf("invalid %s", tok)
	} else {
		p.Val = val
	}
	return
}

const validMonths = "JanuaryFebruaryMarchAprilMayJune" +
	"JulyAugustSeptemberOctoberNovemberDecember"

func lexMonth(s *Scanner) (p Part, next lexFn) {
	m := token.Month
	p, next = Part{Tok: m, Pos: s.Pos()}, lexSep
	val, ok := s.Scan("JFMASOND")
	if !ok {
		p.Errorf("invalid %s", m)
		return
	}
	rest, _ := s.ScanAll("abcdefghijklmnopqrstuvxyz")
	p.Val = val + rest
	if !strings.Contains(validMonths, p.Val) {
		p.Errorf("invalid %s", m)
		return
	}
	if p := skipToNextLine(s); p.Defined() {
		return p, next
	}
	return
}

func skipToNextLine(s *Scanner) (p Part) {
	pos := s.Pos()
	s.ScanAll(" \t")
	_, ok := s.Scan("\n")
	if !ok {
		p = Part{Tok: token.Error, Pos: pos, Val: "expect newline"}
	}
	return
}

func lexSep(s *Scanner) (p Part, next lexFn) {
	p, next = ScanPart(s, token.Separator), lexWeek
	s.Scan("\n")
	return
}

func lexWeek(s *Scanner) (p Part, next lexFn) {
	s.Scan(" ") // eg. for week numbers 1-9
	next = lexDate
	r := s.Next()
	if r == EOS {
		return p, nil
	}
	s.Back()
	if s.PeekIs(" ") {
		s.ScanAll(" ")
		return
	}
	p = ScanPart(s, token.Week)
	s.ScanAll(" ")
	return
}

func lexDate(s *Scanner) (p Part, next lexFn) {
	p, next = ScanPart(s, token.Date), lexDay
	s.Scan(" ")
	return
}

const validDays = "MonTueWedThuFriSatSun"

func lexDay(s *Scanner) (p Part, next lexFn) {
	p, next = Part{Tok: token.Day, Pos: s.Pos()}, lexNote
	val, ok := s.Scan("MTWFS")
	if !ok {
		p.Errorf("invalid %s", token.Day)
	} else {
		rest, _ := s.ScanAll("aedhniortu")
		p.Val = val + rest
		if len(p.Val) != 3 || !strings.Contains(validDays, p.Val) {
			p.Errorf("invalid %s", token.Day)
		}
	}
	s.Scan(" ")
	switch {
	case s.PeekIs("-+"):
		next = lexOperator
	case s.PeekIs("\n"): // no hours reported
		s.Scan("\n")
		next = lexWeek
	case s.PeekIs(digits):
		next = lexHours
	}
	return
}

const digits = "0123456789"

func lexOperator(s *Scanner) (p Part, next lexFn) {
	p, next = Part{Tok: token.Operator, Pos: s.Pos()}, lexHours
	r, ok := s.Scan("+-")
	p.Val = string(r)
	if !ok {
		p.Errorf("invalid %s", token.Operator)
	}
	return
}

func lexHours(s *Scanner) (p Part, next lexFn) {
	p = ScanPart(s, token.Hours)
	if s.PeekIs(":") {
		return p, lexColon
	}
	s.ScanAll(" ")
	if s.inTag {
		return p, lexTag
	}
	return p, lexNote
}

func lexColon(s *Scanner) (p Part, next lexFn) {
	p = Part{Tok: token.Colon, Pos: s.Pos()}
	next = lexMinutes
	r := s.Next()
	if r != ':' {
		p.Tok = token.Undefined
		return
	}
	p.Val = string(r)
	return
}

func lexMinutes(s *Scanner) (Part, lexFn) {
	p := ScanPart(s, token.Minutes)
	s.ScanAll(" ")
	if s.inTag {
		return p, lexTag
	}
	return p, lexNote
}

func lexRightParen(s *Scanner) (p Part, next lexFn) {
	p = Part{Tok: token.RightParenthesis, Pos: s.Pos()}
	val, ok := s.Scan(")")
	if !ok {
		p.Errorf("invalid %s", token.RightParenthesis)
		return
	}
	p.Val = val
	next = lexNote
	s.inTag = false
	s.ScanAll(" ")
	return
}

func lexLeftParen(s *Scanner) (p Part, next lexFn) {
	p, next = Part{Tok: token.LeftParenthesis, Pos: s.Pos()}, lexHours
	val, ok := s.Scan("(")
	if !ok {
		p.Errorf("invalid %s", token.LeftParenthesis)
		return
	}
	s.Scan(" ")
	p.Val = val
	if s.PeekIs("+-") {
		next = lexOperator
	}
	s.inTag = true
	return
}

func lexNote(s *Scanner) (p Part, next lexFn) {
	p = Part{Tok: token.Note, Pos: s.Pos()}
	var r rune
	for r = s.Next(); r != '('; r = s.Next() {
		if r == '\n' {
			if p.Val == "" { // skip notes with only newline
				p.Tok = token.Undefined
				p.Val += string(r)
			}
			next = lexWeek
			return
		}
		if r == EOS {
			return
		}
		p.Val += string(r)
	}
	// found a left parenthesis
	s.Back()
	if len(p.Val) == 0 {
		p.Tok = token.Undefined
	}
	next = lexLeftParen
	return
}

func lexTag(s *Scanner) (p Part, next lexFn) {
	p = Part{Tok: token.Tag, Pos: s.Pos()}
	var r rune
	for r = s.Next(); r != ')'; r = s.Next() {
		if r == '\n' || r == EOS || r == '(' {
			p.Pos = s.Pos()
			p.Errorf("missing %s", token.RightParenthesis)
			return
		}
		p.Val += string(r)
	}
	s.Back()
	next = lexRightParen
	p.Val = strings.TrimSpace(p.Val)
	return
}
