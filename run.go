package timesheet

import (
	"strings"
)

func lexYear(s *Scanner) (p Part, next lexFn) {
	p, next = ScanPart(s, Year), lexMonth
	s.Scan(" ")
	return
}

func ScanPart(s *Scanner, tok Token) (p Part) {
	p = Part{Tok: tok, Pos: s.Pos()}
	var valid string
	switch tok {
	case Hours, Minutes, Year, Date, Week:
		valid = digits
	case Separator:
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
	p, next = Part{Tok: Month, Pos: s.Pos()}, lexSep
	val, ok := s.Scan("JFMASOND")
	if !ok {
		p.Errorf("invalid %s", Month)
		return
	}
	rest, _ := s.ScanAll("abcdefghijklmnopqrstuvxyz")
	p.Val = val + rest
	if !strings.Contains(validMonths, p.Val) {
		p.Errorf("invalid %s", Month)
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
		p = Part{Tok: Error, Pos: pos, Val: "expect newline"}
	}
	return
}

func lexSep(s *Scanner) (p Part, next lexFn) {
	p, next = ScanPart(s, Separator), lexWeek
	s.Scan("\n")
	return
}

func lexWeek(s *Scanner) (p Part, next lexFn) {
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
	p = ScanPart(s, Week)
	s.ScanAll(" ")
	return
}

func lexDate(s *Scanner) (p Part, next lexFn) {
	p, next = ScanPart(s, Date), lexDay
	s.Scan(" ")
	return
}

const validDays = "MonTueWenThuFriSatSun"

func lexDay(s *Scanner) (p Part, next lexFn) {
	p, next = Part{Tok: Day, Pos: s.Pos()}, lexHours
	val, ok := s.Scan("MTWFS")
	if !ok {
		p.Errorf("invalid %s", Day)
	} else {
		rest, _ := s.ScanAll("aehniortu")
		p.Val = val + rest
		if len(p.Val) != 3 || !strings.Contains(validDays, p.Val) {
			p.Errorf("invalid %s", Day)
		}
	}
	s.Scan(" ")
	return
}

func lexRightParen(s *Scanner) (p Part, next lexFn) {
	p = Part{Tok: RightParenthesis, Pos: s.Pos()}
	val, ok := s.Scan(")")
	if !ok {
		p.Errorf("invalid %s", RightParenthesis)
		return
	}
	p.Val = val
	next = lexNote
	s.inTag = false
	s.ScanAll(" ")
	return
}

func lexLeftParen(s *Scanner) (p Part, next lexFn) {
	p = Part{Tok: LeftParenthesis, Pos: s.Pos()}
	val, ok := s.Scan("(")
	if !ok {
		p.Errorf("invalid %s", LeftParenthesis)
		return
	}
	p.Val = val
	next = lexOperator
	s.inTag = true
	return
}

func lexNote(s *Scanner) (p Part, next lexFn) {
	p = Part{Tok: Note, Pos: s.Pos()}
	var r rune
	for r = s.Next(); r != '('; r = s.Next() {
		if r == '\n' {
			if p.Val == "" { // skip notes with only newline
				p.Tok = Undefined
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
		p.Tok = Undefined
	}
	next = lexLeftParen
	return
}

const digits = "0123456789"

func lexOperator(s *Scanner) (p Part, next lexFn) {
	p = Part{Tok: Operator, Pos: s.Pos()}
	next = lexHours
	if s.PeekIs("-+") {
		p.Val = string(s.Next())
		return
	}
	if s.PeekIs(digits) {
		p.Tok = Undefined
		return
	}
	p.Errorf("invalid %s", Operator)
	return
}

func lexHours(s *Scanner) (p Part, next lexFn) {
	p = ScanPart(s, Hours)
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
	p = Part{Tok: Colon, Pos: s.Pos()}
	next = lexMinutes
	r := s.Next()
	if r != ':' {
		p.Tok = Undefined
		return
	}
	p.Val = string(r)
	return
}

func lexMinutes(s *Scanner) (Part, lexFn) {
	p := ScanPart(s, Minutes)
	s.ScanAll(" ")
	if s.inTag {
		return p, lexTag
	}
	return p, lexNote
}

func lexTag(s *Scanner) (p Part, next lexFn) {
	p = Part{Tok: Tag, Pos: s.Pos()}
	var r rune
	for r = s.Next(); r != ')'; r = s.Next() {
		if r == '\n' || r == EOS {
			p.Pos = s.Pos()
			p.Errorf("missing %s", RightParenthesis)
			return
		}
		p.Val += string(r)
	}
	s.Back()
	next = lexRightParen
	p.Val = strings.TrimSpace(p.Val)
	return
}
