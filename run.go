package timesheet

import (
	"strings"
)

func lexDate(s *Scanner, out chan Part) lexFn {
	p := Part{Tok: Number, Pos: s.Pos()}
	val, ok := s.ScanAll("0123456789")
	if !ok {
		p.Errorf("invalid date")
	} else {
		p.Val = val
	}
	out <- p
	return nil
}

func lexWeek(s *Scanner, out chan Part) lexFn {
	p := Part{Tok: Number, Pos: s.Pos()}
	val, ok := s.ScanAll("0123456789")
	if !ok {
		p.Errorf("invalid week")
	} else {
		p.Val = val
	}
	out <- p
	return nil
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
	if s.PeekIs(" ") { // No week number
		s.ScanAll(" ")
		return lexDate
	}
	return lexWeek
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
	pos := s.Pos()
	val, ok := s.ScanAll("0123456789")
	p := Part{Tok: Number, Val: val, Pos: pos}
	if !ok {
		p.Errorf("invalid year")
	}
	out <- p
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
