package timesheet

import (
	"strings"
)

const digits = "0123456789"

/*
func lexOperator(s *Scanner, out chan Part) lexFn {
	p := Part{Tok: Operator, Pos: s.Pos()}
	val, ok := s.Scan("-+")
	if !ok {
		p.Errorf("invalid %s", Operator)
		out <- p
	} else {
		p.Val = val
		out <- p
	}
	return nil
}
*/
func lexLeftParenthesis(s *Scanner) (p Part, next lexFn) {
	p = Part{Tok: LeftParenthesis, Pos: s.Pos()}
	val, ok := s.Scan("(")
	if !ok {
		p.Errorf("invalid %s", LeftParenthesis)
		return
	}
	p.Val = val
	return p, nil //lexOperator
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
		p.Tok = Undefined
		next = lexLeftParenthesis
		return
	}
	// should be unreachable
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
	p, next = ScanPart(s, Number), lexDay
	s.Scan(" ")
	return
}

func lexWeek(s *Scanner) (p Part, next lexFn) {
	next = lexDate
	if s.PeekIs(" ") {
		s.ScanAll(" ")
		return
	}
	p = ScanPart(s, Number)
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
	p, next = ScanPart(s, Number), lexMonth
	s.Scan(" ")
	return
}

func ScanPart(s *Scanner, tok Token) (p Part) {
	p = Part{Tok: tok, Pos: s.Pos()}
	var valid string
	switch tok {
	case Number:
		valid = "0123456789"
	case Hours:
		valid = "0123456789"
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
