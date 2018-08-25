package timesheet

import (
	"strings"
)

const (
	digits = "0123456789"
)

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
			p.Val += string(r)
			next = lexWeek
			return
		}
		if r == EOS {
			return
		}
		p.Val += string(r)
	}
	if r == '(' {
		s.Back()
		if len(p.Val) == 0 {
			p.Tok = Undefined
		}
		next = lexLeftParen
		return
	}
	// should be unreachable
	return
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

func lexMinutes(s *Scanner) (p Part, next lexFn) {
	p = ScanPart(s, Minutes)
	s.ScanAll(" ")
	if s.inTag {
		next = lexTag
	} else {
		next = lexNote
	}
	return
}

func lexColon(s *Scanner) (p Part, next lexFn) {
	p = Part{Tok: Colon, Pos: s.Pos()}
	val := s.Next()
	if val == ':' {
		p.Val = ":"
		next = lexMinutes
		return
	}
	s.Back()
	p.Tok = Undefined
	if s.inTag {
		next = lexTag
	} else {
		next = lexNote
	}
	return
}

func lexHours(s *Scanner) (p Part, next lexFn) {
	p = ScanPart(s, Hours)
	next = lexColon
	return
}

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

func lexReported(s *Scanner) (p Part, next lexFn) {
	p, next = Part{Pos: s.Pos()}, lexWeek
	if s.PeekIs("\n") {
		s.Scan("\n")

		return
	}
	p = ScanPart(s, Hours)
	next = lexNote
	return
}

const validDays = "MonTueWenThuFriSatSun"

func lexDay(s *Scanner) (p Part, next lexFn) {
	p, next = Part{Tok: Day, Pos: s.Pos()}, lexReported
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

func lexDate(s *Scanner) (p Part, next lexFn) {
	p, next = ScanPart(s, Date), lexDay
	s.Scan(" ")
	return
}

func lexWeek(s *Scanner) (p Part, next lexFn) {
	next = lexDate
	if s.PeekIs(" ") {
		s.ScanAll(" ")
		return
	}
	p = ScanPart(s, Week)
	s.ScanAll(" ")
	return
}

func lexSep(s *Scanner) (p Part, next lexFn) {
	p, next = ScanPart(s, Separator), lexWeek
	s.Scan("\n")
	return
}

const validMonths = "JanuaryFebruaryMarchAprilMayJune" +
	"JulyAugustSeptemberOctoberNovemberDecember"

func lexMonth(s *Scanner) (p Part, next lexFn) {
	p, next = Part{Tok: Month, Pos: s.Pos()}, lexSep
	val, ok := s.Scan("JFMASOND")
	if !ok {
		p.Errorf("invalid month")
		return
	}
	rest, _ := s.ScanAll("abcdefghijklmnopqrstuvxyz")
	p.Val = val + rest
	if !strings.Contains(validMonths, p.Val) {
		p.Errorf("invalid month")
		return
	}
	if p := skipToNextLine(s); p.Defined() {
		return p, next
	}
	return
}

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

func skipToNextLine(s *Scanner) (p Part) {
	pos := s.Pos()
	s.ScanAll(" \t")
	_, ok := s.Scan("\n")
	if !ok {
		p = Part{Tok: Error, Pos: pos, Val: "expect newline"}
	}
	return
}

func (l *Lexer) run(start lexFn, s *Scanner, out chan Part) {
	// We expect to start the file with a year
	for p, next := start(s); next != nil; p, next = next(s) {
		if p.Tok != Undefined {
			out <- p
		}
	}
	close(out)
}

type lexFn func(s *Scanner) (Part, lexFn)
