package timesheet

import (
	"strings"
)

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
	return nil
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
		rest, _ := s.ScanAll("abcdefghijklmnopqrstuvxyz")
		p.Val = val + rest
		if len(p.Val) != 3 || !strings.Contains(validDays, p.Val) {
			p.Errorf("invalid %s", Day)
		}
	}
	out <- p
	s.ScanAll(" ")
	return lexReported
}

func lexDate(s *Scanner, out chan Part) lexFn {
	out <- ScanPart(s, Number)
	s.Scan(" ")
	return lexDay
}

func lexWeek(s *Scanner, out chan Part) lexFn {
	// todo week num is only available sometimes
	if s.PeekIs(" ") {
		s.ScanAll(" ")
		return lexDate
	}
	out <- ScanPart(s, Number)
	s.ScanAll(" ")
	return lexDate
}

func lexSep(s *Scanner, out chan Part) lexFn {
	out <- ScanPart(s, Separator)
	s.Scan("\n")
	return lexWeek
}

const validMonths = "JanuaryFebruaryMarchAprilMayJune" +
	"JulyAugustSeptemberOctoberNovemberDecember"

func lexMonth(s *Scanner, out chan Part) lexFn {
	p := Part{Tok: Month, Pos: s.Pos()}
	val, ok := s.Scan("JFMASOND")
	if !ok {
		p.Errorf("invalid month")
	} else {
		rest, _ := s.ScanAll("abcdefghijklmnopqrstuvxyz")
		p.Val = val + rest
		if !strings.Contains(validMonths, p.Val) {
			p.Errorf("invalid month")
		}
	}
	out <- p
	skipToNextLine(s, out)
	return lexSep
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

func skipToNextLine(s *Scanner, out chan Part) {
	pos := s.Pos()
	s.ScanAll(" \t")
	_, ok := s.Scan("\n")
	if !ok {
		out <- Part{Pos: pos, Val: "expect newline"}
	}
}

func lexYear(s *Scanner, out chan Part) lexFn {
	out <- ScanPart(s, Number)
	s.Scan(" ")
	return nil
}

func (l *Lexer) run(start lexFn, s *Scanner, out chan Part) {
	// We expect to start the file with a year
	for fn := start; fn != nil; fn = fn(s, out) {
	}
	close(out)
}

type lexFn func(s *Scanner, out chan Part) lexFn
