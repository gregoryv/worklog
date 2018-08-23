package timesheet

import (
	"strings"
)

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

func lexLeftParenthesis(s *Scanner, out chan Part) lexFn {
	p := Part{Tok: LeftParenthesis, Pos: s.Pos()}
	val, ok := s.Scan("(")
	if !ok {
		p.Errorf("invalid %s", LeftParenthesis)
		out <- p
		return nil
	}
	p.Val = val
	out <- p
	return lexOperator
}

func lexReported(s *Scanner, out chan Part) lexFn {
	if s.PeekIs(" \n") {
		skipToNextLine(s, out)
		return lexWeek
	}
	out <- ScanPart(s, Number)
	s.ScanAll(" ")
	if s.PeekIs("(") {
		return lexLeftParenthesis
	}
	if s.PeekIs("\n") {
		s.Scan("\n")
		return lexWeek
	}
	p := Part{}
	p.Errorf("invalid %s", Number)
	out <- p
	return nil
}

const validDays = "MonTueWenThuFriSatSun"

func lexDay(s *Scanner, out chan Part) lexFn {
	p := Part{Tok: Day, Pos: s.Pos()}
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
	out <- p
	s.ScanAll(" ")
	return lexReported
}
*/
func lexDate(s *Scanner) (p Part, next lexFn) {
	p, next = ScanPart(s, Number), nil //lexDay
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
